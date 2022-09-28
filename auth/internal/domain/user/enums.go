package user

type UserRole int8

const (
	Administrator UserRole = iota
	DeliveryMan
	Root
)

type DeviceType int8

const (
	Web DeviceType = iota
	IOS
	Android
)
