package code

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"user-service/internal/model/entity"
	def "user-service/internal/repository"
)

var _ def.CodeRepository = (*repository)(nil)

const (
	tableName       = "code"
	idColumnName    = "id"
	emailColumnName = "email"
)

type repository struct {
	client *pgxpool.Pool
}

func NewRepository(client *pgxpool.Pool) *repository {
	return &repository{client: client}
}

func (r *repository) Save(email string) (entity.Code, error) {
	builder := squirrel.Insert(tableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns(emailColumnName).
		Values(email).
		Suffix("RETURNING *")

	sql, args, err := builder.ToSql()
	if err != nil {
		return entity.CodeNil(), err
	}

	ctx := context.TODO()

	rows, err := r.client.Query(ctx, sql, args...)
	if err != nil {
		return entity.CodeNil(), err
	}

	result, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.Code])
	if err != nil {
		return entity.CodeNil(), err
	}

	return result, nil
}

func (r *repository) Get(id uuid.UUID) (entity.Code, error) {
	builder := squirrel.Select(idColumnName, emailColumnName).
		PlaceholderFormat(squirrel.Dollar).
		From(tableName).
		Where(squirrel.Eq{idColumnName: id.String()})

	sql, args, err := builder.ToSql()
	if err != nil {
		return entity.CodeNil(), err
	}

	ctx := context.TODO()

	rows, err := r.client.Query(ctx, sql, args...)
	if err != nil {
		return entity.CodeNil(), err
	}

	result, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.Code])
	if err != nil {
		return entity.CodeNil(), err
	}

	return result, nil
}

func (r *repository) Delete(id uuid.UUID) error {
	builder := squirrel.Delete(tableName).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{idColumnName: id.String()}).
		Suffix(fmt.Sprintf("RETURNING %s", idColumnName))

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
