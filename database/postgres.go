package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/braejan/platzi-go-rest-websockets/entities"
	"github.com/braejan/platzi-go-rest-websockets/repository"
	"github.com/braejan/platzi-go-rest-websockets/utils"
)

type PostgresRepositoryImpl struct {
	db *sql.DB
}

func (repository *PostgresRepositoryImpl) Init(ctx *context.Context) (err error) {
	if repository.db == nil {
		databaseUrl := utils.GetEnvValue(utils.DATABASE_URL)
		repository.db, err = repository.openDatabase(ctx, databaseUrl)
	}
	return
}

func NewPostgresRepository(ctx *context.Context) (repository repository.PostgresRepository, err error) {
	repository = &PostgresRepositoryImpl{}
	repository.Init(ctx)
	return
}

// User repository interface implementation
const (
	CreateUserQuery = `INSERT INTO users (ID, email, password) VALUES($1, $2, $3)`
)

func (repository *PostgresRepositoryImpl) CreateUser(ctx context.Context, user *entities.User) (err error) {
	_, err = repository.db.ExecContext(ctx, CreateUserQuery, user.ID, user.Email, user.Password)
	return
}

const (
	GetUserByIDQuery = "SELECT email FROM user where ID = $1"
)

func (repository *PostgresRepositoryImpl) GetUserByID(ctx context.Context, userID int64) (user *entities.User, err error) {
	rows, err := repository.db.QueryContext(ctx, GetUserByIDQuery, userID)
	if err != nil {
		return
	}
	defer func() {
		err = rows.Close()
	}()
	_ = rows.Next()
	err = rows.Scan(&user.Email)
	if err != nil {
		return
	}
	user.ID = userID
	user.Password = fmt.Sprintf("%32s", "*")
	return
}
func (repository *PostgresRepositoryImpl) DeleteUserByID(ctx context.Context, userID int64) (err error) {
	return
}
func (repository *PostgresRepositoryImpl) UpdateUser(ctx context.Context, user *entities.User) (err error) {
	return
}

// Postgresql interface implementation

func (repository *PostgresRepositoryImpl) openDatabase(ctx *context.Context, url string) (db *sql.DB, err error) {
	db, err = sql.Open("postgres", url)
	return
}
func (repository *PostgresRepositoryImpl) BeginTransaction(ctx *context.Context) (dbTx *sql.Tx, err error) {
	dbTx, err = repository.db.BeginTx(*ctx, nil)
	return
}
func (repository *PostgresRepositoryImpl) CommitTransaction(dbTx *sql.Tx) (err error) {
	err = dbTx.Commit()
	return
}
func (repository *PostgresRepositoryImpl) RollbackTransaction(dbTx *sql.Tx) (err error) {
	err = dbTx.Rollback()
	return
}
