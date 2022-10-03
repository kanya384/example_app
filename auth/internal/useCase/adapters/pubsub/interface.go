package pubsub

import "context"

type Notification interface {
	SendEmail(ctx context.Context, email string, subject string, message string)
}
