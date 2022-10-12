package device

type DeviceType int8

const (
	Web DeviceType = iota
	IOS
	Android
)
