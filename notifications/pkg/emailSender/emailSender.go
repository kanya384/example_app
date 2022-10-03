package email

import (
	"context"
	"fmt"
	"net/smtp"
)

type Client interface {
	SendEmail(ctx context.Context, sendTo, subject, text string) (err error)
}

type client struct {
	auth    smtp.Auth
	sender  string
	address string
}

func NewEmailClient(host, port, sender, password string) Client {
	address := fmt.Sprintf("%s:%s", host, port)
	auth := smtp.PlainAuth("", sender, password, address)
	return &client{
		auth:    auth,
		sender:  sender,
		address: address,
	}
}

func (c *client) SendEmail(ctx context.Context, sendTo, subject, text string) (err error) {
	err = smtp.SendMail(c.address, c.auth, c.sender, []string{sendTo}, []byte(text))
	return
}
