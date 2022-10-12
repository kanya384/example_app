package emailClient

import "context"

type EmailClient interface {
	SendEmail(ctx context.Context, sendTo, subject, text string) (err error)
}
