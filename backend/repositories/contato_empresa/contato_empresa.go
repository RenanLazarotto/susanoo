package contatoempresa

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"

	"tsukuyomi/models"
	"tsukuyomi/repositories"
)

type Repository interface {
	Create(ctx context.Context, contato models.ContatoEmpresa) (models.ContatoEmpresa, error)
	FindAll(ctx context.Context, search, empresa, tipo, contato string) ([]models.ContatoEmpresa, error)
	FindByID(ctx context.Context, id string) (models.ContatoEmpresa, error)
	Update(ctx context.Context, contato models.ContatoEmpresa) error
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

func (r *repository) Create(ctx context.Context, contato models.ContatoEmpresa) (models.ContatoEmpresa, error) {
	result, err := r.DB().Write(
		ctx,
		`INSERT INTO contato_empresa(id_empresa, tipo, contato, criado)
		VALUES(?, ?, ?, ?)`,
		contato.IDEmpresa,
		contato.Tipo,
		contato.Contato,
		contato.Criado,
	)

	if err != nil {
		log.Error(repositories.ERROR_INSERT, err)
		return models.ContatoEmpresa{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Error(repositories.ERROR_INSERT, err)
		return models.ContatoEmpresa{}, err
	}

	contato.ID = id

	return contato, nil
}

func (r *repository) FindAll(ctx context.Context, search, empresa, tipo, contato string) ([]models.ContatoEmpresa, error) {
	arguments := []interface{}{}
	var searchLike string

	conditions := ""

	if search != "" {
		searchLike = fmt.Sprintf("%%%s%%", search)
		conditions += " AND (nome LIKE ? OR cnpj LIKE ?)"
		arguments = append(arguments, searchLike, searchLike)
	}

	if empresa != "" {
		conditions += " AND (id_empresa = ?)"
		arguments = append(arguments, empresa)
	}

	if tipo != "" {
		conditions += " AND (tipo = ?)"
		arguments = append(arguments, tipo)
	}

	if contato != "" {
		conditions += " AND (contato = ?)"
		arguments = append(arguments, contato)
	}

	rows, err := r.DB().Select(
		ctx,
		`SELECT	
			cont.id,
			emp.id,
			emp.nome,
			emp.cnpj,
			emp.criado,
			emp.atualizado,
			emp.apagado,
			cont.tipo,
			cont.contato,
			cont.criado,
			cont.atualizado,
			cont.apagado
		FROM contato_empresa cont
		JOIN empresas emp ON emp.id = cont.id_empresa
		WHERE cont.apagado IS NULL
		AND emp.apagado IS NULL 
		`+conditions,
		arguments...,
	)

	if err != nil {
		log.Error(repositories.ERROR_SELECT, err)
		return []models.ContatoEmpresa{}, err
	}

	defer rows.Close()

	var contatos []models.ContatoEmpresa

	for rows.Next() {
		var contato = models.ContatoEmpresa{}

		err := rows.Scan(
			&contato.ID,
			&contato.Empresa.ID,
			&contato.Empresa.Nome,
			&contato.Empresa.CNPJ,
			&contato.Empresa.Criado,
			&contato.Empresa.Atualizado,
			&contato.Empresa.Apagado,
			&contato.Tipo,
			&contato.Contato,
			&contato.Criado,
			&contato.Atualizado,
			&contato.Apagado,
		)

		if err != nil {
			log.Error(repositories.ERROR_SELECT_SCAN, err)
			return []models.ContatoEmpresa{}, err
		}

		contatos = append(contatos, contato)
	}

	return contatos, nil
}

func (r *repository) FindByID(ctx context.Context, id string) (models.ContatoEmpresa, error) {
	rows, err := r.DB().Select(
		ctx,
		`SELECT	
			cont.id,
			emp.id,
			emp.nome,
			emp.cnpj,
			emp.criado,
			emp.atualizado,
			emp.apagado,
			cont.tipo,
			cont.contato,
			cont.criado,
			cont.atualizado,
			cont.apagado
		FROM contato_empresa cont
		JOIN empresas emp ON emp.id = cont.id_empresa
		WHERE cont.apagado IS NULL
		AND emp.apagado IS NULL 
		AND cont.id = ?`,
		id,
	)

	if err != nil {
		log.Error(repositories.ERROR_SELECT, err)
		return models.ContatoEmpresa{}, err
	}

	defer rows.Close()

	var contato = models.ContatoEmpresa{}

	for rows.Next() {
		err := rows.Scan(
			&contato.ID,
			&contato.Empresa.ID,
			&contato.Empresa.Nome,
			&contato.Empresa.CNPJ,
			&contato.Empresa.Criado,
			&contato.Empresa.Atualizado,
			&contato.Empresa.Apagado,
			&contato.Tipo,
			&contato.Contato,
			&contato.Criado,
			&contato.Atualizado,
			&contato.Apagado,
		)

		if err != nil {
			log.Error(repositories.ERROR_SELECT_SCAN, err)
			return models.ContatoEmpresa{}, err
		}
	}

	return contato, nil
}

func (r *repository) Update(ctx context.Context, contato models.ContatoEmpresa) error {
	_, err := r.DB().Write(
		ctx,
		`UPDATE contato_empresa SET 
		id_empresa = ?, 
		tipo = ?, 
		contato = ?,
		atualizado = ?
		WHERE id = ?`,
		contato.IDEmpresa,
		contato.Tipo,
		contato.Contato,
		contato.Atualizado,
		contato.ID,
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
		`UPDATE contato_empresa SET 
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
