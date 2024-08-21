package database

import (
	"gorm.io/gorm"

	"tsukuyomi/config"
)

type Service struct {
	tx         *gorm.Tx
	connection *gorm.DB
}

func New(config config.Config) (*Service, error) {
	return &Service{}, nil
}
