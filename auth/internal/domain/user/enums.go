package user

type UserRole int8

const (
	Administrator UserRole = iota
	DeliveryMan
	Root
)
