package models

import (
	"time"

	"github.com/charmbracelet/log"
	"gorm.io/gorm"
)

type Empresa struct {
	ID         int            `gorm:"primaryKey" json:"id"`
	Nome       string         `gorm:"column:nome;size:255" json:"nome"`
	CNPJ       string         `gorm:"column:cnpj;size:20" json:"cnpj"`
	Criado     time.Time      `gorm:"column:criado;autoCreateTime"`
	Atualizado *time.Time     `gorm:"column:atualizado;autoUpdateTime"`
	Apagado    gorm.DeletedAt `gorm:"column:apagado"`
}

func (e *Empresa) AfterCreate(tx *gorm.DB) error {
	// TODO: insert log reg
	log.Info("Created empresa", "empresa", e)
	return nil
}
