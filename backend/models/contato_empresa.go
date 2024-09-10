package models

import "time"

type ContatoEmpresa struct {
	ID         int64      `json:"id"`
	IDEmpresa  int64      `json:"id_empresa,omitempty"`
	Empresa    Empresa    `json:"empresa,omitempty"`
	Tipo       string     `json:"tipo"`
	Contato    string     `json:"contato"`
	Criado     time.Time  `json:"criado"`
	Atualizado *time.Time `json:"atualizado"`
	Apagado    *time.Time `json:"apagado"`
}
