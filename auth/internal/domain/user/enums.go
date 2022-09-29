package user

type UserRole string

const (
	Administrator UserRole = "administrator"
	DeliveryMan   UserRole = "deliveryman"
	Root          UserRole = "root"
	UserR         UserRole = "user"
)
