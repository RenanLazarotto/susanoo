package endereco

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"

	"tsukuyomi/models"
	"tsukuyomi/repositories"
)

type Repository interface {
	Create(ctx context.Context, endereco models.Endereco) (models.Endereco, error)
	FindAll(ctx context.Context, search, logradouro, numero, complemento, bairro, cidade, cep, estado string) ([]models.Endereco, error)
	FindByID(ctx context.Context, id string) (models.Endereco, error)
	Update(ctx context.Context, endereco models.Endereco) error
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

func (r *repository) Create(ctx context.Context, endereco models.Endereco) (models.Endereco, error) {
	result, err := r.DB().Write(
		ctx,
		`INSERT INTO enderecos(logradouro, numero, complemento, bairro, cidade, cep, estado, criado)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?)`,
		endereco.Logradouro,
		endereco.Numero,
		endereco.Complemento,
		endereco.Bairro,
		endereco.Cidade,
		endereco.CEP,
		endereco.Estado,
		endereco.Criado,
	)

	if err != nil {
		log.Error(repositories.ERROR_INSERT, err)
		return models.Endereco{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Error(repositories.ERROR_INSERT, err)
		return models.Endereco{}, err
	}

	endereco.ID = id

	return endereco, nil
}

func (r *repository) FindAll(ctx context.Context, search, logradouro, numero, complemento, bairro, cidade, cep, estado string) ([]models.Endereco, error) {
	arguments := []interface{}{}
	var searchLike string

	conditions := ""

	if search != "" {
		searchLike = fmt.Sprintf("%%%s%%", search)
		conditions += " AND (logradouro LIKE ? OR numero LIKE ? OR complemento LIKE ? or bairro LIKE ? OR cidade LIKE ? OR cep LIKE ? OR estado LIKE ?)"
		arguments = append(arguments, searchLike, searchLike, searchLike, searchLike, searchLike, searchLike, searchLike)
	}

	if logradouro != "" {
		conditions += " AND (logradouro = ?)"
		arguments = append(arguments, logradouro)
	}

	if numero != "" {
		conditions += " AND (numero = ?)"
		arguments = append(arguments, numero)
	}

	if complemento != "" {
		conditions += " AND (complemento = ?)"
		arguments = append(arguments, complemento)
	}

	if bairro != "" {
		conditions += " AND (bairro = ?)"
		arguments = append(arguments, bairro)
	}

	if cidade != "" {
		conditions += " AND (cidade = ?)"
		arguments = append(arguments, cidade)
	}

	if cep != "" {
		conditions += " AND (cep = ?)"
		arguments = append(arguments, cep)
	}

	if estado != "" {
		conditions += " AND (estado = ?)"
		arguments = append(arguments, estado)
	}

	rows, err := r.DB().Select(
		ctx,
		`SELECT	
			* 
		FROM enderecos 
		WHERE apagado IS NULL 
		`+conditions,
		arguments...,
	)

	if err != nil {
		log.Error(repositories.ERROR_SELECT, err)
		return []models.Endereco{}, err
	}

	defer rows.Close()

	var enderecos []models.Endereco

	for rows.Next() {
		var endereco = &models.Endereco{}

		err = rows.Scan(
			&endereco.ID,
			&endereco.Logradouro,
			&endereco.Numero,
			&endereco.Complemento,
			&endereco.Bairro,
			&endereco.Cidade,
			&endereco.CEP,
			&endereco.Estado,
			&endereco.Criado,
			&endereco.Apagado,
			&endereco.Atualizado,
		)

		if err != nil {
			log.Error(repositories.ERROR_SELECT_SCAN, err)
			return []models.Endereco{}, err
		}

		rows, err := r.DB().Select(
			ctx,
			`SELECT	
				emp.* 
			FROM empresas emp
			JOIN endereco_empresa endemp ON emp.id = endemp.id_empresa
			JOIN enderecos end ON endemp.id_endereco = end.id
			WHERE end.id = ?
				AND end.apagado IS NULL
				AND emp.apagado IS NULL
				AND endemp.apagado IS NULL`,
			endereco.ID,
		)

		if err != nil {
			log.Error(repositories.ERROR_SELECT, err)
			return []models.Endereco{}, err
		}

		defer rows.Close()

		for rows.Next() {
			var empresa = &models.Empresa{}

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
				return []models.Endereco{}, err
			}

			endereco.Empresas = append(endereco.Empresas, empresa)
		}

		enderecos = append(enderecos, *endereco)
	}

	return enderecos, nil
}

func (r *repository) FindByID(ctx context.Context, id string) (models.Endereco, error) {
	rows, err := r.DB().Select(
		ctx,
		`SELECT	
			* 
		FROM enderecos 
		WHERE apagado IS NULL 
		AND id = ?`,
		id,
	)

	if err != nil {
		log.Error(repositories.ERROR_SELECT, err)
		return models.Endereco{}, err
	}

	defer rows.Close()

	endereco := models.Endereco{}

	for rows.Next() {
		err = rows.Scan(
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
			log.Error(repositories.ERROR_SELECT_SCAN, err)
			return models.Endereco{}, err
		}

		rows, err := r.DB().Select(
			ctx,
			`SELECT	
				emp.* 
			FROM empresas emp
			JOIN endereco_empresa endemp ON emp.id = endemp.id_empresa
			JOIN enderecos end ON endemp.id_endereco = end.id
			WHERE end.id = ?
				AND end.apagado IS NULL
				AND emp.apagado IS NULL
				AND endemp.apagado IS NULL`,
			endereco.ID,
		)

		if err != nil {
			log.Error(repositories.ERROR_SELECT, err)
			return models.Endereco{}, err
		}

		defer rows.Close()

		for rows.Next() {
			var empresa = &models.Empresa{}

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
				return models.Endereco{}, err
			}

			endereco.Empresas = append(endereco.Empresas, empresa)
		}
	}

	return endereco, nil
}

func (r *repository) Update(ctx context.Context, endereco models.Endereco) error {
	_, err := r.DB().Write(
		ctx,
		`UPDATE enderecos SET 
		logradouro = ?, 
		numero = ?, 
		complemento = ?, 
		bairro = ?, 
		cidade = ?, 
		cep = ?, 
		estado = ?, 
		atualizado = ?
		WHERE id = ?`,
		endereco.Logradouro,
		endereco.Numero,
		endereco.Complemento,
		endereco.Bairro,
		endereco.Cidade,
		endereco.CEP,
		endereco.Estado,
		endereco.Atualizado,
		endereco.ID,
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
		`UPDATE enderecos SET 
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
