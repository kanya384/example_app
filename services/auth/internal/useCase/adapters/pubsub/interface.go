package pubsub

import "context"

type Notification interface {
	SendMessage(ctx context.Context, key string, msg []byte) (err error)
}
