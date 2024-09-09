package models

import (
	"time"
)

type Endereco struct {
	ID          int64      `json:"id"`
	Logradouro  string     `json:"logradouro"`
	Numero      string     `json:"numero"`
	Complemento *string    `json:"complemento"`
	Bairro      string     `json:"bairro"`
	Cidade      string     `json:"cidade"`
	CEP         string     `json:"cep"`
	Estado      string     `json:"estado"`
	Empresas    []*Empresa `json:"empresas"`
	Criado      time.Time  `json:"criado"`
	Atualizado  *time.Time `json:"atualizado"`
	Apagado     *time.Time `json:"apagado"`
}
