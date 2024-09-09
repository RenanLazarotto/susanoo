package models

import (
	"time"
)

type Empresa struct {
	ID         int64       `json:"id"`
	Nome       string      `json:"nome"`
	CNPJ       string      `json:"cnpj"`
	Enderecos  []*Endereco `json:"enderecos,omitempty"`
	Criado     time.Time   `json:"criado"`
	Atualizado *time.Time  `json:"atualizado"`
	Apagado    *time.Time  `json:"apagado"`
}
