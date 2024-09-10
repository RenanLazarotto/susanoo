package emprego

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"

	"tsukuyomi/models"
	"tsukuyomi/repositories"
)

type Repository interface {
	Create(ctx context.Context, emprego models.Emprego) (models.Emprego, error)
	FindAll(ctx context.Context, search, empresa, ocupacao, remuneracao_inicial, tipo_contrato, data_inicio, data_fim, carga_horaria string) ([]models.Emprego, error)
	FindByID(ctx context.Context, id string) (models.Emprego, error)
	Update(ctx context.Context, emprego models.Emprego) error
	Delete(ctx context.Context, id string) error
}

type repository struct {
	repositories.Repository
}

func NewRepository(repo repositories.Repository) Repository {
	return &repository{
		Repository: repo,
	}
}

func (r *repository) Create(ctx context.Context, emprego models.Emprego) (models.Emprego, error) {
	result, err := r.DB().Write(
		ctx,
		`INSERT INTO empregos(id_empresa, ocupacao, remuneracao_inicial, tipo_contrato, data_inicio, data_fim, carga_horaria, criado)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?)`,
		emprego.IDEmpresa,
		emprego.Ocupacao,
		emprego.RemuneracaoInicial,
		emprego.TipoContrato,
		emprego.DataInicio,
		emprego.DataFim,
		emprego.CargaHoraria,
		emprego.Criado,
	)

	if err != nil {
		log.Error(repositories.ERROR_INSERT, err)
		return models.Emprego{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Error(repositories.ERROR_INSERT, err)
		return models.Emprego{}, err
	}

	emprego.ID = id

	return emprego, nil
}

func (r *repository) FindAll(ctx context.Context, search, empresa, ocupacao, remuneracao_inicial, tipo_contrato, data_inicio, data_fim, carga_horaria string) ([]models.Emprego, error) {
	arguments := []interface{}{}
	var searchLike string

	conditions := ""

	if search != "" {
		searchLike = fmt.Sprintf("%%%s%%", search)
		conditions += " AND (emp.nome LIKE ? OR emp.cnpj LIKE ? OR job.ocupacao LIKE ? OR job.remuneracao_inicial LIKE ? OR job.tipo_contrato LIKE ? OR job.carga_horaria LIKE ?)"
		arguments = append(arguments, searchLike, searchLike, searchLike, searchLike, searchLike, searchLike)
	}

	if empresa != "" {
		conditions += " AND (cont.id_empresa = ? OR emp.nome = ?)"
		arguments = append(arguments, empresa)
	}

	if ocupacao != "" {
		conditions += " AND (job.ocupacao = ?)"
		arguments = append(arguments, ocupacao)
	}

	if remuneracao_inicial != "" {
		conditions += " AND (job.remuneracao_inicial = ?)"
		arguments = append(arguments, remuneracao_inicial)
	}

	if tipo_contrato != "" {
		conditions += " AND (job.tipo_contrato = ?)"
		arguments = append(arguments, tipo_contrato)
	}

	if data_inicio != "" {
		conditions += " AND (job.data_inicio = STR_TO_DATE(?, '%Y-%m-%d'))"
		arguments = append(arguments, data_inicio)
	}

	if data_fim != "" {
		conditions += " AND (job.data_fim = STR_TO_DATE(?, '%Y-%m-%d')))"
		arguments = append(arguments, data_fim)
	}

	if carga_horaria != "" {
		conditions += " AND (job.carga_horaria = ?)"
		arguments = append(arguments, carga_horaria)
	}

	rows, err := r.DB().Select(
		ctx,
		`SELECT	
			job.id,
			emp.id,
			emp.nome,
			emp.cnpj,
			emp.criado,
			emp.atualizado,
			emp.apagado,
			job.ocupacao,
			job.remuneracao_inicial,
			job.tipo_contrato,
			job.data_inicio,
			job.data_fim,
			job.carga_horaria,
			job.criado,
			job.atualizado,
			job.apagado
		FROM empregos job
		JOIN empresas emp ON emp.id = job.id_empresa
		WHERE job.apagado IS NULL
		AND emp.apagado IS NULL 
		`+conditions,
		arguments...,
	)

	if err != nil {
		log.Error(repositories.ERROR_SELECT, err)
		return []models.Emprego{}, err
	}

	defer rows.Close()

	var empregos []models.Emprego

	for rows.Next() {
		var emprego = models.Emprego{}

		err := rows.Scan(
			&emprego.ID,
			&emprego.Empresa.ID,
			&emprego.Empresa.Nome,
			&emprego.Empresa.Nome,
			&emprego.Empresa.CNPJ,
			&emprego.Empresa.Criado,
			&emprego.Empresa.Atualizado,
			&emprego.Empresa.Apagado,
			&emprego.Ocupacao,
			&emprego.RemuneracaoInicial,
			&emprego.TipoContrato,
			&emprego.DataInicio,
			&emprego.DataFim,
			&emprego.CargaHoraria,
			&emprego.Criado,
			&emprego.Atualizado,
			&emprego.Apagado,
		)

		if err != nil {
			log.Error(repositories.ERROR_SELECT_SCAN, err)
			return []models.Emprego{}, err
		}

		empregos = append(empregos, emprego)
	}

	return empregos, nil
}

func (r *repository) FindByID(ctx context.Context, id string) (models.Emprego, error) {
	rows, err := r.DB().Select(
		ctx,
		`SELECT	
			job.id,
			emp.id,
			emp.nome,
			emp.cnpj,
			emp.criado,
			emp.atualizado,
			emp.apagado,
			job.ocupacao,
			job.remuneracao_inicial,
			job.tipo_contrato,
			job.data_inicio,
			job.data_fim,
			job.carga_horaria,
			job.criado,
			job.atualizado,
			job.apagado
		FROM empregos job
		JOIN empresas emp ON emp.id = job.id_empresa
		WHERE job.apagado IS NULL
		AND emp.apagado IS NULL 
		AND job.id = ?`,
		id,
	)

	if err != nil {
		log.Error(repositories.ERROR_SELECT, err)
		return models.Emprego{}, err
	}

	defer rows.Close()

	var emprego = models.Emprego{}

	for rows.Next() {
		err := rows.Scan(
			&emprego.ID,
			&emprego.Empresa.ID,
			&emprego.Empresa.Nome,
			&emprego.Empresa.Nome,
			&emprego.Empresa.CNPJ,
			&emprego.Empresa.Criado,
			&emprego.Empresa.Atualizado,
			&emprego.Empresa.Apagado,
			&emprego.Ocupacao,
			&emprego.RemuneracaoInicial,
			&emprego.TipoContrato,
			&emprego.DataInicio,
			&emprego.DataFim,
			&emprego.CargaHoraria,
			&emprego.Criado,
			&emprego.Atualizado,
			&emprego.Apagado,
		)

		if err != nil {
			log.Error(repositories.ERROR_SELECT_SCAN, err)
			return models.Emprego{}, err
		}
	}

	return emprego, nil
}

func (r *repository) Update(ctx context.Context, emprego models.Emprego) error {
	_, err := r.DB().Write(
		ctx,
		`UPDATE empregos SET 
		id_empresa = ?, 
		ocupacao = ?, 
		remuneracao_inicial = ?,
		tipo_contrato = ?,
		data_inicio = ?,
		data_fim = ?,
		carga_horaria = ?,
		atualizado = ?,
		WHERE id = ?`,
		emprego.IDEmpresa,
		emprego.Ocupacao,
		emprego.RemuneracaoInicial,
		emprego.TipoContrato,
		emprego.DataInicio,
		emprego.DataFim,
		emprego.CargaHoraria,
		emprego.Atualizado,
		emprego.ID,
	)

	if err != nil {
		log.Error(repositories.ERROR_UPDATE, err)
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	_, err := r.DB().Write(
		ctx,
		`UPDATE empregos SET 
		atualizado = CURRENT_DATE(),
		apagado = CURRENT_DATE()
		WHERE id = ?`,
		id,
	)

	if err != nil {
		log.Error(repositories.ERROR_DELETE, err)
		return err
	}

	return nil
}
