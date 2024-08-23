package models

import (
	"time"

	"github.com/charmbracelet/log"
	"gorm.io/gorm"
)

type Endereco struct {
	ID          int            `gorm:"primaryKey" json:"id"`
	Logradouro  string         `gorm:"column:logradouro;size:255" json:"logradouro"`
	Numero      int            `gorm:"column:numero" json:"numero"`
	Complemento *string        `gorm:"column:complemento;size:100" json:"complemento"`
	Bairro      string         `gorm:"column:bairro;size:100" json:"bairro"`
	Cidade      string         `gorm:"column:cidade;size:100" json:"cidade"`
	CEP         string         `gorm:"column:cep;size:9" json:"cep"`
	Estado      string         `gorm:"column:estado;size:20" json:"estado"`
	Criado      time.Time      `gorm:"column:criado;autoCreateTime"`
	Atualizado  *time.Time     `gorm:"column:atualizado;autoUpdateTime"`
	Apagado     gorm.DeletedAt `gorm:"column:apagado"`
}

func (e *Endereco) AfterCreate(tx *gorm.DB) error {
	// TODO: insert log reg
	log.Info("Created endereco", "empresa", e)
	return nil
}
