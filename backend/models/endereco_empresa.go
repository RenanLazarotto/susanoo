package models

import (
	"time"
)

type EnderecoEmpresa struct {
	ID         int64      `json:"id"`
	Empresa    Empresa    `json:"empresa"`
	Endereco   Endereco   `json:"endereco"`
	Criado     time.Time  `json:"criado"`
	Atualizado *time.Time `json:"atualizado"`
	Apagado    *time.Time `json:"apagado"`
}
