package models

import "time"

type Emprego struct {
	ID                 int64      `json:"id"`
	IDEmpresa          int64      `json:"id_empresa,omitempty"`
	Empresa            Empresa    `json:"empresa,omitempty"`
	Ocupacao           string     `json:"ocupacao"`
	RemuneracaoInicial float64    `json:"remuneracao_inicial"`
	TipoContrato       string     `json:"tipo_contrato"`
	DataInicio         time.Time  `json:"data_inicio"`
	DataFim            *time.Time `json:"data_fim"`
	CargaHoraria       int64      `json:"carga_horaria"`
	Criado             time.Time  `json:"criado"`
	Atualizado         *time.Time `json:"atualizado"`
	Apagado            *time.Time `json:"apagado"`
}
