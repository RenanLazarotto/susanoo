package database

import (
	"errors"
	"fmt"

	"github.com/charmbracelet/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"tsukuyomi/config"
)

type Service interface {
	BeginTransaction() *gorm.DB
	Commit() error
	Query() (*gorm.DB, error)
	Rollback() error
}

type service struct {
	connection *gorm.DB
	tx         *gorm.DB
}

func New(config *config.Config) Service {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Pass,
		config.Database.Host,
		config.Database.Port,
		config.Database.Schema,
		config.Database.Charset,
		config.Database.Collation,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: NewDatabaseLogger().LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	return &service{
		connection: db,
	}
}

func (s *service) BeginTransaction() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}

	s.tx = s.connection.Begin()
	return s.tx
}

func (s *service) Commit() error {
	if s.tx == nil {
		return errors.New("transaction not started")
	}

	defer func() {
		s.tx = nil
	}()

	if err := s.tx.Commit(); err.Error != nil {
		s.tx.Rollback()
		return err.Error
	}

	return nil
}

func (s *service) Query() (*gorm.DB, error) {
	if s.tx != nil {
		return nil, errors.New("transaction in progress")
	}

	return s.connection, nil

}

func (s *service) Rollback() error {
	if s.tx != nil {
		return errors.New("transaction not started")
	}

	defer func() {
		s.tx = nil
	}()

	if err := s.tx.Rollback(); err != nil {
		return err.Error
	}

	return nil
}
