package blockchain

type NotificationType int

type NotificationCallback func()

type Notification struct {
	Type NotificationType
	Data interface{}
}
