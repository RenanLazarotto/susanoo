package models

import (
	"time"
)

type EndereoEmpresaDTO struct {
	IDEmpresa  string `json:"id_empresa"`
	IDEndereco string `json:"id_endereco"`
}

type EnderecoEmpresa struct {
	ID         int64      `json:"id"`
	Empresa    Empresa    `json:"empresa"`
	Endereco   Endereco   `json:"endereco"`
	Criado     time.Time  `json:"criado"`
	Atualizado *time.Time `json:"atualizado"`
	Apagado    *time.Time `json:"apagado"`
}
