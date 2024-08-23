package models

import (
	"time"

	"github.com/charmbracelet/log"
	"github.com/invopop/validation"
	"gorm.io/gorm"
)

type Empresa struct {
	ID         int            `gorm:"primaryKey" json:"id"`
	Nome       string         `gorm:"column:nome;size:255;unique;not null" json:"nome"`
	CNPJ       string         `gorm:"column:cnpj;size:20;unique;not null" json:"cnpj"`
	Criado     time.Time      `gorm:"column:criado;autoCreateTime"`
	Atualizado *time.Time     `gorm:"column:atualizado;autoUpdateTime"`
	Apagado    gorm.DeletedAt `gorm:"column:apagado"`
}

func (e Empresa) BeforeCreate(tx *gorm.DB) error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Nome, validation.Required, validation.NilOrNotEmpty),
		validation.Field(&e.CNPJ, validation.Required, validation.NilOrNotEmpty),
	)
}

func (e *Empresa) AfterCreate(tx *gorm.DB) error {
	// TODO: insert log reg
	log.Info("Created empresa", "empresa", e)
	return nil
}
