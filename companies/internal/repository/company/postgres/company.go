package postgres

import (
	"companies/internal/domain/company"
	"companies/internal/domain/company/inn"
	"companies/internal/domain/company/name"
	"companies/internal/repository/company/postgres/dao"
	"companies/pkg/types/address"
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const (
	tableName = "company"
)

var (
	ErrDuplicateKey = errors.New("ERROR: duplicate key value violates unique constraint \"company_pkey\" (SQLSTATE 23505)")
	ErrNotFound     = errors.New("not found")
	ErrEmptyResult  = errors.New("no rows in result set")
)

func (r *Repository) CreateCompany(ctx context.Context, company *company.Company) (err error) {
	rawQuery := r.Builder.Insert(tableName).Columns(dao.ColumnsCompany...).Values(company.ID(), company.Name(), company.Inn(), company.Address(), company.CreatedAt(), company.ModifiedAt())
	query, args, _ := rawQuery.ToSql()

	_, err = r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return
	}
	return
}

func (r *Repository) UpdateCompany(ctx context.Context, company *company.Company) (err error) {
	rawQuery := r.Builder.Update(tableName).Set("name", company.Name()).Set("inn", company.Inn()).Set("address", company.Address()).Set("modified_at", company.ModifiedAt()).Where("id = ?", company.ID())
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

func (r *Repository) DeleteCompany(ctx context.Context, ID uuid.UUID) (err error) {
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

func (r *Repository) ReadCompanyByID(ctx context.Context, ID uuid.UUID) (cmp *company.Company, err error) {
	rawQuery := r.Builder.Select(dao.ColumnsCompany...).From(tableName).Where("id = ?", ID)
	query, args, _ := rawQuery.ToSql()

	row, err := r.Pool.Query(ctx, query, args...)
	if err != nil {
		return
	}
	daoCompany, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[dao.Company])
	if err != nil {
		return
	}

	//daoCompany.CreatedAt, daoCompany.ModifiedAt,
	return company.NewWithID(daoCompany.ID, name.Name(daoCompany.Name), inn.Inn(daoCompany.Inn), address.Address(daoCompany.Address), daoCompany.CreatedAt, daoCompany.ModifiedAt)
}
