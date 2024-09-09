package repositories

import (
	"tsukuyomi/config"
	"tsukuyomi/database"
)

const (
	ERROR_DELETE      = "error while soft-deleting record"
	ERROR_INSERT      = "error while inserting record"
	ERROR_SELECT      = "error while querying records"
	ERROR_SELECT_SCAN = "error while scanning result into model"
	ERROR_UPDATE      = "error while updating record"
)

type Repository interface {
	DB() database.DatabaseService
}

type Condition struct {
	Value    interface{}
	Operator string
	Format   string
}

type repository struct {
	db database.DatabaseService
}

func NewRepository(config *config.Config) Repository {
	db := database.New(config)

	db.StartConnection()

	return &repository{
		db: db,
	}
}

func (r repository) DB() database.DatabaseService {
	return r.db
}
