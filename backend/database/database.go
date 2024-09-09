package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"tsukuyomi/config"
)

const (
	ERROR_TX_NOT_STARTED = "transaction not started"
)

type Query interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type DatabaseService interface {
	StartConnection() error
	Select(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Write(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	BeginTransaction(ctx context.Context) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type service struct {
	tx *sql.Tx
	ro *connection
	rw *connection
}

type connection struct {
	db        *sql.DB
	dsn       string
	lastError error
}

func New(config *config.Config) DatabaseService {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&parseTime=True&loc=Local&multiStatements=True",
		config.Database.User,
		config.Database.Pass,
		config.Database.Host,
		config.Database.Port,
		config.Database.Schema,
		config.Database.Charset,
		config.Database.Collation,
	)
	return &service{
		ro: &connection{
			dsn: dsn,
		},
		rw: &connection{
			dsn: dsn,
		},
	}
}

func (c *connection) Connect() {
	c.db, c.lastError = sql.Open("mysql", c.dsn)
}

func (s *service) StartConnection() error {
	if s.connectRO().lastError != nil {
		return s.ro.lastError
	}

	if s.connectRW().lastError != nil {
		return s.rw.lastError
	}

	return nil
}

func (s *service) connectRO() *connection {
	if s.ro.db == nil {
		s.ro.Connect()
	}

	return s.ro
}

func (s *service) connectRW() *connection {
	if s.rw.db == nil {
		s.rw.Connect()
	}

	return s.rw
}

func (s *service) QueryRO() Query {
	if s.ro.db == nil {
		s.connectRO()
	}

	return s.ro.db
}

func (s *service) QueryRW() Query {
	if s.tx != nil {
		return s.tx
	}

	if s.rw.db == nil {
		s.connectRW()
	}

	return s.rw.db
}

func (s *service) Select(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := s.QueryRO().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (s *service) Write(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	result, err := s.QueryRW().ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	result, err := s.QueryRW().ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) BeginTransaction(ctx context.Context) error {
	if s.rw.lastError != nil {
		return s.rw.lastError
	}

	tx, err := s.connectRW().db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	s.tx = tx
	return nil
}

func (s *service) Commit(ctx context.Context) error {
	if s.tx == nil {
		return errors.New(ERROR_TX_NOT_STARTED)
	}

	defer func() {
		s.tx = nil
	}()

	return s.tx.Commit()
}

func (s *service) Rollback(ctx context.Context) error {
	if s.tx == nil {
		return errors.New(ERROR_TX_NOT_STARTED)
	}

	defer func() {
		s.tx = nil
	}()

	return s.tx.Rollback()
}
