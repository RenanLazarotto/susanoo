package repositories

import (
	"tsukuyomi/config"
	"tsukuyomi/database"
)

const (
	ERROR_DELETE      = "erro ao apagar registro"
	ERROR_INSERT      = "erro ao inserir registro"
	ERROR_SELECT      = "erro ao realizer consulta"
	ERROR_SELECT_SCAN = "erro ao associar valores da consulta à struct"
	ERROR_UPDATE      = "erro ao atualizar registro"
	ERROR_VALIDATE    = "erro ao validar struct"
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
