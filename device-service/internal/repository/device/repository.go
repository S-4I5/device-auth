package device

import (
	"context"
	"device-service/internal/model/entity"
	"device-service/internal/repository/device/model"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	tableName             = "device"
	idColumnName          = "id"
	phoneNumberColumnName = "phone_number"
	pinCodeColumnName     = "pin_code"
	userIdColumnName      = "user_id"
	isVerifiedColumnName  = "is_verified"
)

type repository struct {
	client *pgxpool.Pool
}

func NewRepository(client *pgxpool.Pool) *repository {
	return &repository{client: client}
}

func (r *repository) Create(device entity.Device) (entity.Device, error) {

	builder := squirrel.Insert(tableName).PlaceholderFormat(squirrel.Dollar).
		Columns(phoneNumberColumnName).
		Values(device.PhoneNumber).
		Suffix("RETURNING *")

	sql, args, err := builder.ToSql()
	if err != nil {
		return entity.DeviceNil(), err
	}

	ctx := context.TODO()

	rows, err := r.client.Query(ctx, sql, args...)
	if err != nil {
		return entity.DeviceNil(), err
	}

	result, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Device])
	if err != nil {
		fmt.Println("XD")
		return entity.DeviceNil(), err
	}

	return model.DbDeviceToDevice(result), nil
}

func (r *repository) Get(id uuid.UUID) (entity.Device, error) {
	builder := squirrel.Select("*").PlaceholderFormat(squirrel.Dollar).
		From(tableName).
		Where(squirrel.Eq{idColumnName: id.String()})

	sql, args, err := builder.ToSql()
	if err != nil {
		return entity.DeviceNil(), err
	}

	ctx := context.TODO()

	rows, err := r.client.Query(ctx, sql, args...)
	if err != nil {
		return entity.DeviceNil(), err
	}

	result, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Device])
	if err != nil {
		return entity.DeviceNil(), err
	}

	return model.DbDeviceToDevice(result), nil
}

func (r *repository) SetVerified(id uuid.UUID) error {
	builder := squirrel.Update(tableName).PlaceholderFormat(squirrel.Dollar).
		Set(isVerifiedColumnName, true).
		Where(squirrel.Eq{idColumnName: id})

	sql, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	ctx := context.TODO()

	rows, err := r.client.Query(ctx, sql, args...)
	if err != nil {
		return err
	}

	return rows.Err()
}

func (r *repository) SetPin(id uuid.UUID, pin string) error {
	builder := squirrel.Update(tableName).PlaceholderFormat(squirrel.Dollar).
		Set(pinCodeColumnName, pin).
		Where(squirrel.Eq{idColumnName: id})

	sql, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	ctx := context.TODO()

	rows, err := r.client.Query(ctx, sql, args...)
	if err != nil {
		return err
	}

	rows.Next()

	return rows.Err()
}

func (r *repository) SetUser(id uuid.UUID, userId uuid.UUID) error {
	builder := squirrel.Update(tableName).PlaceholderFormat(squirrel.Dollar).
		Set(userIdColumnName, userId).
		Where(squirrel.Eq{idColumnName: id})

	sql, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	ctx := context.TODO()

	rows, err := r.client.Query(ctx, sql, args...)
	if err != nil {
		return err
	}

	rows.Next()

	return rows.Err()
}
