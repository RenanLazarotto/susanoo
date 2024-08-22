package repositories

import (
	"tsukuyomi/config"
	"tsukuyomi/database"
)

type Repository interface {
	DB() database.Service
}

type repository struct {
	db database.Service
}

func NewRepository(config *config.Config) Repository {
	return &repository{
		db: database.New(config),
	}
}

func (r repository) DB() database.Service {
	return r.db
}
