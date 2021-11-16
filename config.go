package main

import (
	"btc-node/blockchain"
	"btc-node/chaincfg"
	"btc-node/database"
	"btc-node/mempool"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"btc-node/peer"

	"github.com/btcsuite/btcutil"
	"github.com/jessevdk/go-flags"
)

const (
	defaultConfigFilename        = "btcd.conf"
	defaultDataDirname           = "data"
	defaultLogLevel              = "info"
	defaultLogDirname            = "logs"
	defaultLogFilename           = "btcd.log"
	defaultMaxPeers              = 125
	defaultBanDuration           = time.Hour * 24
	defaultBanThreshold          = 100
	defaultConnectTimeout        = time.Second * 30
	defaultMaxRPCClients         = 10
	defaultMaxRPCWebsockets      = 25
	defaultMaxRPCConcurrentReqs  = 20
	defaultDbType                = "ffldb"
	defaultFreeTxRelayLimit      = 15.0
	defaultTrickleInterval       = peer.DefaultTrickleInterval
	defaultBlockMinSize          = 0
	defaultBlockMaxSize          = 750000
	defaultBlockMinWeight        = 0
	defaultBlockMaxWeight        = 3000000
	blockMaxSizeMin              = 1000
	blockMaxSizeMax              = blockchain.MaxBlockBaseSize - 1000
	blockMaxWeightMin            = 4000
	blockMaxWeightMax            = blockchain.MaxBlockWeight - 4000
	defaultGenerate              = false
	defaultMaxOrphanTransactions = 100
	defaultMaxOrphanTxSize       = 100000
	defaultSigCacheMaxSize       = 100000
	sampleConfigFilename         = "sample-btcd.conf"
	defaultTxIndex               = false
	defaultAddrIndex             = false
)

var (
	defaultHomeDir     = btcutil.AppDataDir("btcd", false)
	defaultConfigFile  = filepath.Join(defaultHomeDir, defaultConfigFilename)
	defaultDataDir     = filepath.Join(defaultHomeDir, defaultDataDirname)
	knownDbTypes       = database.SupportedDrivers()
	defaultRPCKeyFile  = filepath.Join(defaultHomeDir, "rpc.key")
	defaultRPCCertFile = filepath.Join(defaulthomedire, "rpc.cert")
	defaultLogDir      = filepath.Join(defaultHomeDir, defaultLogDirname)
)

type config struct {
	AddCheckpoints       []string      `long:"addcheckpoint" description:"Add a custom checkpoint.  Format: '<height>:<hash>'"`
	AddPeers             []string      `short:"a" long:"addpeer" description:"Add a peer to connect with at startup"`
	AddrIndex            bool          `long:"addrindex" description:"Maintain a full address-based transaction index which makes the searchrawtransactions RPC available"`
	AgentBlacklist       []string      `long:"agentblacklist" description:"A comma separated list of user-agent substrings which will cause btcd to reject any peers whose user-agent contains any of the blacklisted substrings."`
	AgentWhitelist       []string      `long:"agentwhitelist" description:"A comma separated list of user-agent substrings which will cause btcd to require all peers' user-agents to contain one of the whitelisted substrings. The blacklist is applied before the blacklist, and an empty whitelist will allow all agents that do not fail the blacklist."`
	BanDuration          time.Duration `long:"banduration" description:"How long to ban misbehaving peers.  Valid time units are {s, m, h}.  Minimum 1 second"`
	BanThreshold         uint32        `long:"banthreshold" description:"Maximum allowed ban score before disconnecting and banning misbehaving peers."`
	BlockMaxSize         uint32        `long:"blockmaxsize" description:"Maximum block size in bytes to be used when creating a block"`
	BlockMinSize         uint32        `long:"blockminsize" description:"Mininum block size in bytes to be used when creating a block"`
	BlockMaxWeight       uint32        `long:"blockmaxweight" description:"Maximum block weight to be used when creating a block"`
	BlockMinWeight       uint32        `long:"blockminweight" description:"Mininum block weight to be used when creating a block"`
	BlockPrioritySize    uint32        `long:"blockprioritysize" description:"Size in bytes for high-priority/low-fee transactions when creating a block"`
	BlocksOnly           bool          `long:"blocksonly" description:"Do not accept transactions from remote peers."`
	ConfigFile           string        `short:"C" long:"configfile" description:"Path to configuration file"`
	ConnectPeers         []string      `long:"connect" description:"Connect only to the specified peers at startup"`
	CPUProfile           string        `long:"cpuprofile" description:"Write CPU profile to the specified file"`
	DataDir              string        `short:"b" long:"datadir" description:"Directory to store data"`
	DbType               string        `long:"dbtype" description:"Database backend to use for the Block Chain"`
	DebugLevel           string        `short:"d" long:"debuglevel" description:"Logging level for all subsystems {trace, debug, info, warn, error, critical} -- You may also specify <subsystem>=<level>,<subsystem2>=<level>,... to set the log level for individual subsystems -- Use show to list available subsystems"`
	DropAddrIndex        bool          `long:"dropaddrindex" description:"Deletes the address-based transaction index from the database on start up and then exits."`
	DropCfIndex          bool          `long:"dropcfindex" description:"Deletes the index used for committed filtering (CF) support from the database on start up and then exits."`
	DropTxIndex          bool          `long:"droptxindex" description:"Deletes the hash-based transaction index from the database on start up and then exits."`
	ExternalIPs          []string      `long:"externalip" description:"Add an ip to the list of local addresses we claim to listen on to peers"`
	Generate             bool          `long:"generate" description:"Generate (mine) bitcoins using the CPU"`
	FreeTxRelayLimit     float64       `long:"limitfreerelay" description:"Limit relay of transactions with no transaction fee to the given amount in thousands of bytes per minute"`
	Listeners            []string      `long:"listen" description:"Add an interface/port to listen for connections (default all interfaces port: 8333, testnet: 18333)"`
	LogDir               string        `long:"logdir" description:"Directory to log output."`
	MaxOrphanTxs         int           `long:"maxorphantx" description:"Max number of orphan transactions to keep in memory"`
	MaxPeers             int           `long:"maxpeers" description:"Max number of inbound and outbound peers"`
	MiningAddrs          []string      `long:"miningaddr" description:"Add the specified payment address to the list of addresses to use for generated blocks -- At least one address is required if the generate option is set"`
	MinRelayTxFee        float64       `long:"minrelaytxfee" description:"The minimum transaction fee in BTC/kB to be considered a non-zero fee."`
	DisableBanning       bool          `long:"nobanning" description:"Disable banning of misbehaving peers"`
	NoCFilters           bool          `long:"nocfilters" description:"Disable committed filtering (CF) support"`
	DisableCheckpoints   bool          `long:"nocheckpoints" description:"Disable built-in checkpoints.  Don't do this unless you know what you're doing."`
	DisableDNSSeed       bool          `long:"nodnsseed" description:"Disable DNS seeding for peers"`
	DisableListen        bool          `long:"nolisten" description:"Disable listening for incoming connections -- NOTE: Listening is automatically disabled if the --connect or --proxy options are used without also specifying listen interfaces via --listen"`
	NoOnion              bool          `long:"noonion" description:"Disable connecting to tor hidden services"`
	NoPeerBloomFilters   bool          `long:"nopeerbloomfilters" description:"Disable bloom filtering support"`
	NoRelayPriority      bool          `long:"norelaypriority" description:"Do not require free or low-fee transactions to have high priority for relaying"`
	NoWinService         bool          `long:"nowinservice" description:"Do not start as a background service on Windows -- NOTE: This flag only works on the command line, not in the config file"`
	DisableRPC           bool          `long:"norpc" description:"Disable built-in RPC server -- NOTE: The RPC server is disabled by default if no rpcuser/rpcpass or rpclimituser/rpclimitpass is specified"`
	DisableTLS           bool          `long:"notls" description:"Disable TLS for the RPC server -- NOTE: This is only allowed if the RPC server is bound to localhost"`
	OnionProxy           string        `long:"onion" description:"Connect to tor hidden services via SOCKS5 proxy (eg. 127.0.0.1:9050)"`
	OnionProxyPass       string        `long:"onionpass" default-mask:"-" description:"Password for onion proxy server"`
	OnionProxyUser       string        `long:"onionuser" description:"Username for onion proxy server"`
	Profile              string        `long:"profile" description:"Enable HTTP profiling on given port -- NOTE port must be between 1024 and 65536"`
	Proxy                string        `long:"proxy" description:"Connect via SOCKS5 proxy (eg. 127.0.0.1:9050)"`
	ProxyPass            string        `long:"proxypass" default-mask:"-" description:"Password for proxy server"`
	ProxyUser            string        `long:"proxyuser" description:"Username for proxy server"`
	RegressionTest       bool          `long:"regtest" description:"Use the regression test network"`
	RejectNonStd         bool          `long:"rejectnonstd" description:"Reject non-standard transactions regardless of the default settings for the active network."`
	RejectReplacement    bool          `long:"rejectreplacement" description:"Reject transactions that attempt to replace existing transactions within the mempool through the Replace-By-Fee (RBF) signaling policy."`
	RelayNonStd          bool          `long:"relaynonstd" description:"Relay non-standard transactions regardless of the default settings for the active network."`
	RPCCert              string        `long:"rpccert" description:"File containing the certificate file"`
	RPCKey               string        `long:"rpckey" description:"File containing the certificate key"`
	RPCLimitPass         string        `long:"rpclimitpass" default-mask:"-" description:"Password for limited RPC connections"`
	RPCLimitUser         string        `long:"rpclimituser" description:"Username for limited RPC connections"`
	RPCListeners         []string      `long:"rpclisten" description:"Add an interface/port to listen for RPC connections (default port: 8334, testnet: 18334)"`
	RPCMaxClients        int           `long:"rpcmaxclients" description:"Max number of RPC clients for standard connections"`
	RPCMaxConcurrentReqs int           `long:"rpcmaxconcurrentreqs" description:"Max number of concurrent RPC requests that may be processed concurrently"`
	RPCMaxWebsockets     int           `long:"rpcmaxwebsockets" description:"Max number of RPC websocket connections"`
	RPCQuirks            bool          `long:"rpcquirks" description:"Mirror some JSON-RPC quirks of Bitcoin Core -- NOTE: Discouraged unless interoperability issues need to be worked around"`
	RPCPass              string        `short:"P" long:"rpcpass" default-mask:"-" description:"Password for RPC connections"`
	RPCUser              string        `short:"u" long:"rpcuser" description:"Username for RPC connections"`
	SigCacheMaxSize      uint          `long:"sigcachemaxsize" description:"The maximum number of entries in the signature verification cache"`
	SimNet               bool          `long:"simnet" description:"Use the simulation test network"`
	SigNet               bool          `long:"signet" description:"Use the signet test network"`
	SigNetChallenge      string        `long:"signetchallenge" description:"Connect to a custom signet network defined by this challenge instead of using the global default signet test network -- Can be specified multiple times"`
	SigNetSeedNode       []string      `long:"signetseednode" description:"Specify a seed node for the signet network instead of using the global default signet network seed nodes"`
	TestNet3             bool          `long:"testnet" description:"Use the test network"`
	TorIsolation         bool          `long:"torisolation" description:"Enable Tor stream isolation by randomizing user credentials for each connection."`
	TrickleInterval      time.Duration `long:"trickleinterval" description:"Minimum time between attempts to send new inventory to a connected peer"`
	TxIndex              bool          `long:"txindex" description:"Maintain a full hash-based transaction index which makes all transactions available via the getrawtransaction RPC"`
	UserAgentComments    []string      `long:"uacomment" description:"Comment to add to the user agent -- See BIP 14 for more information."`
	Upnp                 bool          `long:"upnp" description:"Use UPnP to map our listening port outside of NAT"`
	ShowVersion          bool          `short:"V" long:"version" description:"Display version information and exit"`
	Whitelists           []string      `long:"whitelist" description:"Add an IP network or IP that will not be banned. (eg. 192.168.1.0/24 or ::1)"`
	lookup               func(string) ([]net.IP, error)
	oniondial            func(string, string, time.Duration) (net.Conn, error)
	dial                 func(string, string, time.Duration) (net.Conn, error)
	addCheckpoints       []chaincfg.Checkpoint
	miningAddrs          []btcutil.Address
	minRelayTxFee        btcutil.Amount
	whitelists           []*net.IPNet
}

type serviceOptions struct {
	ServiceCommand string `short:"s" long:"service" description:"Service command {install, remove, start, stop}"`
}

func newConfigParser(cfg *config, so *serviceOptions, options flags.Options) *flags.Parser {
	parser := flags.NewParser(cfg, options)
	if runtime.GOOS == "windows" {
		parser.AddGroup("Service Options", "Service Options", so)
	}

	return parser
}

func loadConfig() (*config, []string, error) {
	cfg := config{
		ConfigFile:           defaultConfigFile,
		DebugLevel:           defaultLogLevel,
		MaxPeers:             defaultMaxPeers,
		BanDuration:          defaultBanDuration,
		BanThreshold:         defaultBanThreshold,
		RPCMaxClients:        defaultMaxRPCClients,
		RPCMaxWebsockets:     defaultMaxRPCClients,
		RPCMaxConcurrentReqs: defaultMaxRPCConcurrentReqs,
		DataDir:              defaultDataDir,
		LogDir:               defaultLogDir,
		DbType:               defaultDbType,
		RPCKey:               defaultRPCKeyFile,
		RPCCert:              defaultRPCCertFile,
		MinRelayTxFee:        mempool.DefaultMinRelayTxFee.ToBTC(),
		FreeTxRelayLimit:     defaultFreeTxRelayLimit,
		TrickleInterval:      defaultTrickleInterval,
		BlockMinSize:         defaultBlockMinSize,
		BlockMaxSize:         defaultBlockMaxSize,
		BlockMinWeight:       defaultBlockMinWeight,
		BlockMaxWeight:       defaultBlockMaxWeight,
		BlockPrioritySize:    mempool.DefaultBlockPrioritySize,
		MaxOrphanTxs:         defaultMaxOrphanTransactions,
		SigCacheMaxSize:      defaultSigCacheMaxSize,
		Generate:             defaultGenerate,
		TxIndex:              defaultTxIndex,
		AddrIndex:            defaultAddrIndex,
	}

	serviceOpts := serviceOptions{}

	preCfg := cfg
	preParser := newConfigParser(&preCfg, &serviceOpts, flags.HelpFlag)
	_, err := preParser.Parse()
	if err != nil {
		if e, ok := err.(*flags.Error); ok && e.Type == flags.ErrHelp {
			fmt.Fprintln(os.Stderr, err)
			return nil, nil, err
		}
	}

	appName := filepath.Base(os.Args[0])
	appName = strings.TrimSuffix(appName, filepath.Ext(appName))
	usageMessage := fmt.Sprintf("Use %s -h to show usage", appName)
	if preCfg.ShowVersion {
		fmt.Println(appName, "version", version())
		os.Exit(0)
	}
}
