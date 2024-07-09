package code

import (
	"context"
	"device-service/internal/model/entity"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	client *pgxpool.Pool
}

const (
	tableName          = "activation_code"
	idColumnName       = "id"
	deviceIdColumnName = "device_id"
	codeColumnName     = "code"
)

func NewRepository(client *pgxpool.Pool) *repository {
	return &repository{client: client}
}

func (r *repository) Create(code entity.ActivationCode) (entity.ActivationCode, error) {
	builder := squirrel.Insert(tableName).PlaceholderFormat(squirrel.Dollar).
		Columns(deviceIdColumnName, codeColumnName).
		Values(code.DeviceId, code.Code).
		Suffix("RETURNING " + idColumnName)

	sql, args, err := builder.ToSql()
	if err != nil {
		return entity.ActivationCodeNil(), err
	}

	ctx := context.TODO()

	rows, err := r.client.Query(ctx, sql, args...)
	if err != nil {
		return entity.ActivationCodeNil(), err
	}

	rows.Next()

	var id uuid.UUID
	err = rows.Scan(&id)
	if err != nil {
		return entity.ActivationCodeNil(), err
	}

	code.Id = id

	return code, nil
}

func (r *repository) GetByDeviceId(deviceId uuid.UUID) (entity.ActivationCode, error) {
	builder := squirrel.Select("*").PlaceholderFormat(squirrel.Dollar).
		From(tableName).
		Where(squirrel.Eq{deviceIdColumnName: deviceId.String()})

	sql, args, err := builder.ToSql()
	if err != nil {
		return entity.ActivationCodeNil(), err
	}

	ctx := context.TODO()

	rows, err := r.client.Query(ctx, sql, args...)
	if err != nil {
		return entity.ActivationCodeNil(), err
	}

	result, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.ActivationCode])
	if err != nil {
		return entity.ActivationCodeNil(), err
	}

	return result, err
}

func (r *repository) Get(id uuid.UUID) (entity.ActivationCode, error) {
	builder := squirrel.Select("*").PlaceholderFormat(squirrel.Dollar).
		From(tableName).
		Where(squirrel.Eq{idColumnName: id.String()})

	sql, args, err := builder.ToSql()
	if err != nil {
		return entity.ActivationCodeNil(), err
	}

	ctx := context.TODO()

	rows, err := r.client.Query(ctx, sql, args...)
	if err != nil {
		return entity.ActivationCodeNil(), err
	}

	result, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.ActivationCode])
	if err != nil {
		return entity.ActivationCodeNil(), err
	}

	return result, err
}

func (r *repository) Delete(id uuid.UUID) error {
	builder := squirrel.Delete(tableName).PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{idColumnName: id.String()})

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
