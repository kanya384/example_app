package domain

type SendStatus int8

const (
	NotSended SendStatus = iota
	Sended
	SendError
)
