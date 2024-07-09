package user

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"user-service/internal/model/entity"
	def "user-service/internal/repository"
)

var _ def.UserRepository = (*repository)(nil)

const (
	userTableName           = "_user"
	idColumnName            = "id"
	emailColumnName         = "email"
	passwordColumnName      = "password"
	emailVerifiedColumnName = "is_email_verified"
	phoneNumberColumnName   = "phone_number"
)

type repository struct {
	client *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *repository {
	return &repository{client: pool}
}

func (r *repository) GetByEmail(email string) (entity.User, error) {
	const op = "user/repository/GetByEmail"

	builder := squirrel.Select("*").
		PlaceholderFormat(squirrel.Dollar).
		From(userTableName).Where(squirrel.Eq{emailColumnName: email})

	sql, args, err := builder.ToSql()
	if err != nil {
		return entity.UserNil(), err
	}

	ctx := context.TODO()

	rows, err := r.client.Query(ctx, sql, args...)
	if err != nil {
		return entity.UserNil(), err
	}

	result, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.User])
	if err != nil {
		return entity.UserNil(), err
	}

	return result, nil
}

func (r *repository) Create(email, password string, phoneNumber string) (entity.User, error) {
	const op = "user/repository/Create"

	builder := squirrel.Insert(userTableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns(emailColumnName, passwordColumnName, emailVerifiedColumnName, phoneNumberColumnName).
		Values(email, password, false, phoneNumber).
		Suffix("RETURNING *")

	sql, args, err := builder.ToSql()
	if err != nil {
		return entity.UserNil(), err
	}

	ctx := context.TODO()

	rows, err := r.client.Query(ctx, sql, args...)

	result, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.User])
	if err != nil {
		return entity.UserNil(), err
	}

	return result, nil
}

func (r *repository) Get(id uuid.UUID) (entity.User, error) {
	const op = "user/repository/Get"

	builder := squirrel.Select("*").
		PlaceholderFormat(squirrel.Dollar).
		From(userTableName).Where(squirrel.Eq{idColumnName: id.String()})

	sql, args, err := builder.ToSql()
	if err != nil {
		return entity.UserNil(), err
	}

	ctx := context.TODO()

	rows, err := r.client.Query(ctx, sql, args...)

	result, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.User])
	if err != nil {
		return entity.UserNil(), err
	}

	return result, nil
}

func (r *repository) VerifyEmail(email string) error {
	const op = "user/repository/VerifyEmail"

	builder := squirrel.Update(userTableName).
		PlaceholderFormat(squirrel.Dollar).
		Set(emailVerifiedColumnName, true).
		Where(squirrel.Eq{emailColumnName: email})

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

func (r *repository) GetByPhoneNumber(phoneNumber string) (entity.User, error) {
	const op = "user/repository/GetByPhoneNumber"

	builder := squirrel.Select("*").
		PlaceholderFormat(squirrel.Dollar).
		From(userTableName).Where(squirrel.Eq{phoneNumberColumnName: phoneNumber})

	sql, args, err := builder.ToSql()
	if err != nil {
		return entity.UserNil(), err
	}

	ctx := context.TODO()

	rows, err := r.client.Query(ctx, sql, args...)

	result, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.User])
	if err != nil {
		return entity.UserNil(), err
	}

	return result, nil
}
