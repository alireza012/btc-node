package main

import "sync"

type wsNotificationManager struct {
	server *rpcServer

	queueNotification chan interface{}

	notificationMsgs chan interface{}

	numClients chan int

	wg   sync.WaitGroup
	quit chan struct{}
}
