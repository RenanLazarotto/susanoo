package enderecoempresa

import (
	"context"
	"time"

	"github.com/charmbracelet/log"

	"tsukuyomi/models"
	"tsukuyomi/repositories"
)

type Repository interface {
	Assign(ctx context.Context, empresa models.Empresa, endereco models.Endereco) (models.EnderecoEmpresa, error)
	GetEmpresasByEndereco(ctx context.Context, id_endereco string) ([]models.Empresa, error)
	GetEnderecosByEmpresa(ctx context.Context, id_empresa string) ([]models.Endereco, error)
}

type repository struct {
	repositories.Repository
}

func NewRepository(repo repositories.Repository) Repository {
	return &repository{
		Repository: repo,
	}
}

func (r *repository) Assign(ctx context.Context, empresa models.Empresa, endereco models.Endereco) (models.EnderecoEmpresa, error) {
	enderecoEmpresa := models.EnderecoEmpresa{
		Empresa:  empresa,
		Endereco: endereco,
		Criado:   time.Now(),
	}

	r.DB().BeginTransaction(ctx)

	result, err := r.DB().Write(
		ctx,
		`INSERT INTO endereco_empresa(id_empresa, id_endereco, criado)
		VALUES(?, ?, ?)`,
		enderecoEmpresa.Empresa.ID,
		enderecoEmpresa.Endereco.ID,
		enderecoEmpresa.Criado,
	)

	if err != nil {
		r.DB().Rollback(ctx)

		log.Error(repositories.ERROR_INSERT, err)
		return models.EnderecoEmpresa{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		r.DB().Rollback(ctx)

		log.Error(repositories.ERROR_INSERT, err)
		return models.EnderecoEmpresa{}, err
	}

	r.DB().Commit(ctx)

	enderecoEmpresa.ID = id

	return enderecoEmpresa, nil
}

func (r *repository) GetEmpresasByEndereco(ctx context.Context, id_endereco string) ([]models.Empresa, error) {
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
		JOIN endereco_empresa endEmp ON endEmp.id_empresa = emp.id
		WHERE endEmp.id_endereco = ?`,
		id_endereco,
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

		empresas = append(empresas, *empresa)
	}

	return empresas, nil
}

func (r *repository) GetEnderecosByEmpresa(ctx context.Context, id_empresa string) ([]models.Endereco, error) {
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
		JOIN endereco_empresa endEmp ON endEmp.id_empresa = end.id
		WHERE endEmp.id_empresa = ?`,
		id_empresa,
	)

	if err != nil {
		log.Error(repositories.ERROR_SELECT, err)
		return []models.Endereco{}, err
	}

	defer rows.Close()

	var enderecos []models.Endereco

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
			log.Error(repositories.ERROR_SELECT_SCAN, err)
			return []models.Endereco{}, err
		}

		enderecos = append(enderecos, *endereco)
	}

	return enderecos, nil
}
