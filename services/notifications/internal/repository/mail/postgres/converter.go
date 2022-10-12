package postgres

import (
	"notifications/internal/domain"
	"notifications/internal/domain/mail"
	"notifications/internal/domain/mail/email"
	"notifications/internal/domain/mail/message"
	"notifications/internal/domain/mail/subject"
	"notifications/internal/repository/mail/postgres/dao"
)

func (r Repository) toDomainMail(dao *dao.Mail) (result *mail.Mail, err error) {
	recipient, err := email.NewEmail(dao.Recipient)
	if err != nil {
		return
	}

	subject, err := subject.NewSubject(dao.Subject)
	if err != nil {
		return
	}

	message, err := message.NewMessage(dao.Message)
	if err != nil {
		return
	}

	result, err = mail.NewWithID(
		dao.ID,
		dao.CreatedAt,
		dao.ModifiedAt,
		*recipient,
		*subject,
		*message,
		domain.SendStatus(dao.Status),
	)

	return
}
