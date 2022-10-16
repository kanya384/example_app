package device

import (
	"auth/internal/domain/device"
	"auth/internal/domain/device/deviceID"
	"auth/internal/domain/device/refreshToken"
	"auth/internal/repository/postgres/device/dao"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const (
	tableName = `public."device"`
)

var (
	ErrDuplicateKey = errors.New("ERROR: duplicate key value violates unique constraint \"device_pkey\" (SQLSTATE 23505)")
	ErrNotFound     = errors.New("not found")
	ErrUpdate       = errors.New("error updating or no changes")
	ErrEmptyResult  = errors.New("no rows in result set")
)

func (r *Repository) CreateDevice(ctx context.Context, device *device.Device) (err error) {
	rawQuery := r.Builder.Insert(tableName).Columns(dao.ColumnsDevice...).Values(device.ID(), device.UserID(), device.DeviceID(), device.Ip(), device.Agent(), device.Type(), device.RefreshToken(), device.RefreshExpiration(), device.LastSeen(), device.CreatedAt(), device.ModifiedAt()) //
	query, args, err := rawQuery.ToSql()
	if err != nil {
		return err
	}

	_, err = r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return
	}
	return
}

func (r *Repository) UpdateDevice(ctx context.Context, ID uuid.UUID, updateFn func(*device.Device) (*device.Device, error)) (device *device.Device, err error) {
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

	upDevice, err := r.oneDeviceTx(ctx, tx, ID)
	if err != nil {
		return
	}

	device, err = updateFn(upDevice)
	if err != nil {
		return
	}

	rawQuery := r.Builder.Update(tableName).Set("ip", device.Ip()).Set("agent", device.Agent()).Set("dtype", device.Type()).Set("refresh_token", device.RefreshToken()).Set("refresh_exp", device.RefreshExpiration()).Set("last_seen", device.LastSeen()).Set("modified_at", device.ModifiedAt())
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

func (r *Repository) DeleteDevice(ctx context.Context, ID uuid.UUID) (err error) {
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

func (r *Repository) ReadDeviceByID(ctx context.Context, ID uuid.UUID) (device *device.Device, err error) {
	rawQuery := r.Builder.Select(dao.ColumnsDevice...).From(tableName).Where("id = ?", ID)
	query, args, _ := rawQuery.ToSql()

	row, err := r.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	daoDevice, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[dao.Device])
	if err != nil {
		return nil, err
	}

	return r.toDomainDevice(&daoDevice)
}

func (r *Repository) ReadDeviceByUserIDAndRefresh(ctx context.Context, userID uuid.UUID, refreshToken refreshToken.RefreshToken) (device *device.Device, err error) {
	rawQuery := r.Builder.Select(dao.ColumnsDevice...).From(tableName).Where("user_id = ? and refresh_token = ?", userID, refreshToken.String())
	query, args, _ := rawQuery.ToSql()

	row, err := r.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	daoDevice, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[dao.Device])
	if err != nil {
		return nil, err
	}

	return r.toDomainDevice(&daoDevice)
}

func (r *Repository) ReadDevicesByDeviceID(ctx context.Context, deviceID deviceID.DeviceID) (device *device.Device, err error) {
	rawQuery := r.Builder.Select(dao.ColumnsDevice...).From(tableName).Where("device_id = ?", deviceID)
	query, args, _ := rawQuery.ToSql()

	row, err := r.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	daoDevice, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[dao.Device])
	if err != nil {
		return nil, err
	}

	return r.toDomainDevice(&daoDevice)
}

func (r *Repository) oneDeviceTx(ctx context.Context, tx pgx.Tx, ID uuid.UUID) (response *device.Device, err error) {
	rawQuery := r.Builder.Select(dao.ColumnsDevice...).From(tableName).Where("id = ?", ID)
	query, args, _ := rawQuery.ToSql()

	row, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	daoDevice, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[dao.Device])
	if err != nil {
		return nil, err
	}

	return r.toDomainDevice(&daoDevice)
}
