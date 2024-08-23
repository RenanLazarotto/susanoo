package models

import (
	"time"

	"github.com/charmbracelet/log"
	"github.com/invopop/validation"
	"gorm.io/gorm"
)

type Endereco struct {
	ID          int            `gorm:"primaryKey" json:"id"`
	Logradouro  string         `gorm:"column:logradouro;size:255;unique;not null" json:"logradouro"`
	Numero      string         `gorm:"column:numero;size:10;not null" json:"numero"`
	Complemento *string        `gorm:"column:complemento;size:100" json:"complemento"`
	Bairro      string         `gorm:"column:bairro;size:100;not null" json:"bairro"`
	Cidade      string         `gorm:"column:cidade;size:100;not null" json:"cidade"`
	CEP         string         `gorm:"column:cep;size:9;not null" json:"cep"`
	Estado      string         `gorm:"column:estado;size:20;not null" json:"estado"`
	Criado      time.Time      `gorm:"column:criado;autoCreateTime"`
	Atualizado  *time.Time     `gorm:"column:atualizado;autoUpdateTime"`
	Apagado     gorm.DeletedAt `gorm:"column:apagado"`
}

func (e Endereco) BeforeCreate(tx *gorm.DB) error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Logradouro, validation.Required, validation.NilOrNotEmpty),
		validation.Field(&e.Numero, validation.Required, validation.NilOrNotEmpty),
		validation.Field(&e.Bairro, validation.Required, validation.NilOrNotEmpty),
		validation.Field(&e.Cidade, validation.Required, validation.NilOrNotEmpty),
		validation.Field(&e.CEP, validation.Required, validation.NilOrNotEmpty),
		validation.Field(&e.Estado, validation.Required, validation.NilOrNotEmpty),
	)
}

func (e *Endereco) AfterCreate(tx *gorm.DB) error {
	// TODO: insert log reg
	log.Info("Created endereco", "empresa", e)
	return nil
}
