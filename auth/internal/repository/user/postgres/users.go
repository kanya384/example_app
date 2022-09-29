package postgres

import (
	"context"
	"errors"
	"fmt"

	"auth/internal/domain/user"
	"auth/internal/repository/user/postgres/dao"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const (
	tableName = `public."user"`
)

var (
	ErrDuplicateKey = errors.New("ERROR: duplicate key value violates unique constraint \"user_pkey\" (SQLSTATE 23505)")
	ErrNotFound     = errors.New("not found")
	ErrUpdate       = errors.New("error updating or no changes")
	ErrEmptyResult  = errors.New("no rows in result set")
)

func (r *Repository) CreateUser(ctx context.Context, user *user.User) (err error) {
	rawQuery := r.Builder.Insert(tableName).Columns(dao.ColumnsUser...).Values(user.ID(), user.Name(), user.Surname(), user.Phone(), user.Pass(), user.Email(), user.Role(), user.CreatedAt(), user.ModifiedAt())
	query, args, _ := rawQuery.ToSql()

	_, err = r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return
	}
	return
}

func (r *Repository) UpdateUser(ctx context.Context, ID uuid.UUID, updateFn func(user *user.User) (*user.User, error)) (user *user.User, err error) {
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

	upUser, err := r.oneUserTx(ctx, tx, ID)
	if err != nil {
		return
	}

	user, err = updateFn(upUser)
	if err != nil {
		return
	}

	rawQuery := r.Builder.Update(tableName).Set("name", user.Name()).Set("surname", user.Surname()).Set("phone", user.Phone()).Set("pass", user.Pass()).Set("email", user.Email()).Set("role", user.Role()).Set("modified_at", user.ModifiedAt()).Where("id = ?", user.ID())
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

func (r *Repository) DeleteUser(ctx context.Context, ID uuid.UUID) (err error) {
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

func (r *Repository) ReadUserByID(ctx context.Context, ID uuid.UUID) (*user.User, error) {
	rawQuery := r.Builder.Select(dao.ColumnsUser...).From(tableName).Where("id = ?", ID)
	query, args, _ := rawQuery.ToSql()

	row, err := r.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	daoUser, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[dao.User])
	if err != nil {
		return nil, err
	}

	return r.toDomainUser(&daoUser)
}

func (r *Repository) oneUserTx(ctx context.Context, tx pgx.Tx, ID uuid.UUID) (response *user.User, err error) {
	rawQuery := r.Builder.Select(dao.ColumnsUser...).From(tableName).Where("id = ?", ID)
	query, args, _ := rawQuery.ToSql()

	row, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	daoUser, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[dao.User])
	if err != nil {
		return nil, err
	}

	return r.toDomainUser(&daoUser)

}
