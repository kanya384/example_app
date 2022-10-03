package postgres

import (
	"context"
	"errors"
	"fmt"
	"notifications/internal/domain/mail"
	"notifications/internal/repository/mail/postgres/dao"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const (
	tableName = `public."mail"`
)

var (
	ErrDuplicateKey = errors.New("ERROR: duplicate key value violates unique constraint \"mail_pkey\" (SQLSTATE 23505)")
	ErrNotFound     = errors.New("not found")
	ErrUpdate       = errors.New("error updating or no changes")
	ErrEmptyResult  = errors.New("no rows in result set")
)

func (r *Repository) CreateMail(ctx context.Context, mail *mail.Mail) (err error) {
	rawQuery := r.Builder.Insert(tableName).Columns(dao.ColumnsMail...).Values(mail.ID(), mail.Recipient(), mail.Subject(), mail.Message(), mail.Status(), mail.CreatedAt(), mail.ModifiedAt())
	query, args, _ := rawQuery.ToSql()

	_, err = r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return
	}
	return
}

func (r *Repository) UpdateMail(ctx context.Context, ID uuid.UUID, updateFn func(mail *mail.Mail) (*mail.Mail, error)) (mail *mail.Mail, err error) {
	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) error {
		if err != nil {
			if errRollback := t.Rollback(ctx); errRollback != nil {
				return fmt.Errorf("transaction rollback error: %s", err)
			}
			return err
		} else {
			if commitErr := tx.Commit(ctx); commitErr != nil {
				return fmt.Errorf("transaction commit error: %s", err)
			}
			return nil
		}
	}(ctx, tx)

	upMail, err := r.oneMailTx(ctx, tx, ID)
	if err != nil {
		return
	}

	mail, err = updateFn(upMail)
	if err != nil {
		return
	}

	rawQuery := r.Builder.Update(tableName).Set("recipient", mail.Recipient()).Set("subject", mail.Subject()).Set("message", mail.Message()).Set("status", mail.Status()).Set("modified_at", mail.ModifiedAt()).Where("id = ?", mail.ID())
	query, args, _ := rawQuery.ToSql()

	res, err := tx.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	if res.RowsAffected() == 0 {
		return nil, ErrUpdate
	}

	return
}

func (r *Repository) DeleteMail(ctx context.Context, ID uuid.UUID) (err error) {
	rawQuery := r.Builder.Delete(tableName).Where("id = ?", ID)
	query, args, _ := rawQuery.ToSql()

	res, err := r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	if res.RowsAffected() == 0 {
		return ErrNotFound
	}
	return
}

func (r *Repository) ReadMailByID(ctx context.Context, ID uuid.UUID) (mail *mail.Mail, err error) {
	rawQuery := r.Builder.Select(dao.ColumnsMail...).From(tableName).Where("id = ?", ID)
	query, args, _ := rawQuery.ToSql()

	row, err := r.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	daoMail, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[dao.Mail])
	if err != nil {
		return nil, err
	}

	return r.toDomainMail(&daoMail)
}

func (r *Repository) ReadMailFiltredList(ctx context.Context, filter map[string]interface{}) (mailsList []*mail.Mail, err error) {
	rawQuery := r.Builder.Select(dao.ColumnsMail...).From(tableName)
	for filtKey, filtVal := range filter {
		switch filtVal.(type) {
		case string:
			rawQuery = rawQuery.Where(sq.Like{filtKey: filtVal})
		default:
			rawQuery = rawQuery.Where(sq.Eq{filtKey: filtVal})
		}
	}

	query, args, err := rawQuery.ToSql()
	if err != nil {
		return
	}

	rows, err := r.Pool.Query(ctx, query, args...)
	if err != nil {
		return
	}

	for rows.Next() {
		mail := dao.Mail{}
		rows.Scan(&mail.ID, &mail.Recipient, &mail.Message, &mail.Subject, &mail.Status, &mail.CreatedAt, &mail.ModifiedAt)

		dMail, err := r.toDomainMail(&mail)
		if err != nil {
			break
		}
		mailsList = append(mailsList, dMail)
	}

	return
}

func (r *Repository) oneMailTx(ctx context.Context, tx pgx.Tx, ID uuid.UUID) (response *mail.Mail, err error) {
	rawQuery := r.Builder.Select(dao.ColumnsMail...).From(tableName).Where("id = ?", ID)
	query, args, _ := rawQuery.ToSql()

	row, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	daoMail, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[dao.Mail])
	if err != nil {
		return nil, err
	}

	return r.toDomainMail(&daoMail)
}
