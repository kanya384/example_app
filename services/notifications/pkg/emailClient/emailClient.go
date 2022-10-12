package emailClient

import (
	"context"
	"fmt"
	"net/smtp"
)

type Client struct {
	auth    smtp.Auth
	sender  string
	address string
}

func NewEmailClient(host, port, sender, password string) *Client {
	address := fmt.Sprintf("%s:%s", host, port)
	auth := smtp.PlainAuth("", sender, password, host)
	return &Client{
		auth:    auth,
		sender:  sender,
		address: address,
	}
}

func (c *Client) SendEmail(ctx context.Context, sendTo, subject, text string) (err error) {
	err = smtp.SendMail(c.address, c.auth, c.sender, []string{sendTo}, []byte(generateMessage(c.sender, sendTo, subject, text)))
	return
}

func generateMessage(sender, sendTo, subject, body string) string {
	return "From: " + sender + "\r\n" +
		"To: " + sendTo + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" + body + "\r\n"
}
