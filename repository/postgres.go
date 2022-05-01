package repository

import (
	"context"
	"database/sql"
)

type PostgresRepository interface {
	UserRepository
	Init(ctx *context.Context) (err error)
	BeginTransaction(ctx *context.Context) (dbTx *sql.Tx, err error)
	CommitTransaction(dbTx *sql.Tx) (err error)
	RollbackTransaction(dbTx *sql.Tx) (err error)
}
