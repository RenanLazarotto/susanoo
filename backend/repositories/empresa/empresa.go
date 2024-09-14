package empresa

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"

	"tsukuyomi/models"
	"tsukuyomi/repositories"
)

type Repository interface {
	Create(ctx context.Context, empresa models.Empresa) (models.Empresa, error)
	FindAll(ctx context.Context, search, nome, cnpj string) ([]models.Empresa, error)
	FindByID(ctx context.Context, id string) (models.Empresa, error)
	Update(ctx context.Context, empresa models.Empresa) error
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

func (r *repository) Create(ctx context.Context, empresa models.Empresa) (models.Empresa, error) {
	r.DB().BeginTransaction(ctx)

	result, err := r.DB().Write(
		ctx,
		`INSERT INTO empresas(nome, cnpj, criado)
		VALUES(?, ?, ?)`,
		empresa.Nome,
		empresa.CNPJ,
		empresa.Criado,
	)

	if err != nil {
		r.DB().Rollback(ctx)

		log.Error(repositories.ERROR_INSERT, err)
		return models.Empresa{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		r.DB().Rollback(ctx)

		log.Error(repositories.ERROR_INSERT, err)
		return models.Empresa{}, err
	}

	r.DB().Commit(ctx)

	empresa.ID = id

	return empresa, nil
}

func (r *repository) FindAll(ctx context.Context, search, nome, cnpj string) ([]models.Empresa, error) {
	arguments := []interface{}{}
	var searchLike string

	conditions := ""

	if search != "" {
		searchLike = fmt.Sprintf("%%%s%%", search)
		conditions += " AND (nome LIKE ? OR cnpj LIKE ?)"
		arguments = append(arguments, searchLike, searchLike)
	}

	if nome != "" {
		conditions += " AND (nome = ?)"
		arguments = append(arguments, nome)
	}

	if cnpj != "" {
		conditions += " AND (cnpj = ?)"
		arguments = append(arguments, cnpj)
	}

	rows, err := r.DB().Select(
		ctx,
		`SELECT	
			emp.id,
			emp.nome,
			emp.cnpj,
			emp.criado,
			emp.atualizado,
			emp.apagado
		FROM empresas emp
		WHERE apagado IS NULL 
		`+conditions,
		arguments...,
	)

	if err != nil {
		log.Error(repositories.ERROR_SELECT, err)
		return []models.Empresa{}, err
	}

	defer rows.Close()

	var empresas []models.Empresa

	for rows.Next() {
		var empresa = &models.Empresa{}

		err := rows.Scan(
			&empresa.ID,
			&empresa.Nome,
			&empresa.CNPJ,
			&empresa.Criado,
			&empresa.Atualizado,
			&empresa.Apagado,
		)

		if err != nil {
			log.Error(repositories.ERROR_SELECT_SCAN, err)
			return []models.Empresa{}, err
		}

		rows, err := r.DB().Select(
			ctx,
			`SELECT 
				end.id,
				end.logradouro,
				end.numero,
				end.complemento,
				end.bairro,
				end.cidade,
				end.cep,
				end.estado,
				end.criado,
				end.atualizado,
				end.apagado
			FROM enderecos end
			JOIN endereco_empresa endemp ON end.id = endemp.id_endereco
			JOIN empresas emp ON endemp.id_empresa = emp.id
			WHERE emp.id = ?
				AND end.apagado IS NULL
				AND emp.apagado IS NULL
				AND endemp.apagado IS NULL`,
			empresa.ID,
		)

		if err != nil {
			log.Error(repositories.ERROR_SELECT, err)
			return []models.Empresa{}, err
		}

		for rows.Next() {
			var endereco = &models.Endereco{}

			err := rows.Scan(
				&endereco.ID,
				&endereco.Logradouro,
				&endereco.Numero,
				&endereco.Complemento,
				&endereco.Bairro,
				&endereco.Cidade,
				&endereco.CEP,
				&endereco.Estado,
				&endereco.Criado,
				&endereco.Atualizado,
				&endereco.Apagado,
			)

			if err != nil {
				log.Error(repositories.ERROR_SELECT, err)
				return []models.Empresa{}, err
			}

			empresa.Enderecos = append(empresa.Enderecos, endereco)
		}

		empresas = append(empresas, *empresa)
	}

	return empresas, nil
}

func (r *repository) FindByID(ctx context.Context, id string) (models.Empresa, error) {
	rows, err := r.DB().Select(
		ctx,
		`SELECT	
			emp.id,
			emp.nome,
			emp.cnpj,
			emp.criado,
			emp.atualizado,
			emp.apagado
		FROM empresas emp
		WHERE apagado IS NULL 
		AND id = ?`,
		id,
	)

	if err != nil {
		log.Error(repositories.ERROR_SELECT, err)
		return models.Empresa{}, err
	}

	defer rows.Close()

	empresa := models.Empresa{}

	for rows.Next() {
		err = rows.Scan(
			&empresa.ID,
			&empresa.Nome,
			&empresa.CNPJ,
			&empresa.Criado,
			&empresa.Atualizado,
			&empresa.Apagado,
		)

		if err != nil {
			log.Error(repositories.ERROR_SELECT_SCAN, err)
			return models.Empresa{}, err
		}

		rows, err := r.DB().Select(
			ctx,
			`SELECT 
				end.id,
				end.logradouro,
				end.numero,
				end.complemento,
				end.bairro,
				end.cidade,
				end.cep,
				end.estado,
				end.criado,
				end.atualizado,
				end.apagado 
			FROM enderecos end
			JOIN endereco_empresa endemp ON end.id = endemp.id_endereco
			JOIN empresas emp ON endemp.id_empresa = emp.id
			WHERE emp.id = ?
				AND end.apagado IS NULL
				AND emp.apagado IS NULL
				AND endemp.apagado IS NULL`,
			empresa.ID,
		)

		if err != nil {
			log.Error(repositories.ERROR_SELECT, err)
			return models.Empresa{}, err
		}

		for rows.Next() {
			var endereco = &models.Endereco{}

			err := rows.Scan(
				&endereco.ID,
				&endereco.Logradouro,
				&endereco.Numero,
				&endereco.Complemento,
				&endereco.Bairro,
				&endereco.Cidade,
				&endereco.CEP,
				&endereco.Estado,
				&endereco.Criado,
				&endereco.Atualizado,
				&endereco.Apagado,
			)

			if err != nil {
				log.Error(repositories.ERROR_SELECT, err)
				return models.Empresa{}, err
			}

			empresa.Enderecos = append(empresa.Enderecos, endereco)
		}
	}

	return empresa, nil
}

func (r *repository) Update(ctx context.Context, empresa models.Empresa) error {
	r.DB().BeginTransaction(ctx)

	_, err := r.DB().Write(
		ctx,
		`UPDATE empresas SET 
		nome = ?, 
		cnpj = ?, 
		atualizado = ?
		WHERE id = ?`,
		empresa.Nome,
		empresa.CNPJ,
		empresa.Atualizado,
		empresa.ID,
	)

	if err != nil {
		r.DB().Rollback(ctx)

		log.Error(repositories.ERROR_UPDATE, err)
		return err
	}

	r.DB().Commit(ctx)

	return nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	r.DB().BeginTransaction(ctx)

	_, err := r.DB().Write(
		ctx,
		`UPDATE empresas SET 
		atualizado = CURRENT_DATE(),
		apagado = CURRENT_DATE()
		WHERE id = ?`,
		id,
	)

	if err != nil {
		r.DB().Rollback(ctx)

		log.Error(repositories.ERROR_DELETE, err)
		return err
	}

	r.DB().Commit(ctx)

	return nil
}
