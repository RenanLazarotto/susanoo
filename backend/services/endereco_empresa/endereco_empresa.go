package enderecoempresa

import (
	"context"

	"tsukuyomi/models"
	enderecoempresa "tsukuyomi/repositories/endereco_empresa"
)

type Service interface {
	Assign(ctx context.Context, empresa models.Empresa, endereco models.Endereco) (models.EnderecoEmpresa, error)
	GetEmpresasByEndereco(ctx context.Context, id_endereco string) ([]models.Empresa, error)
	GetEnderecosByEmpresa(ctx context.Context, id_empresa string) ([]models.Endereco, error)
}

type service struct {
	repository enderecoempresa.Repository
}

func NewService(repository enderecoempresa.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Assign(ctx context.Context, empresa models.Empresa, endereco models.Endereco) (models.EnderecoEmpresa, error) {
	return s.repository.Assign(ctx, empresa, endereco)
}

func (s *service) GetEmpresasByEndereco(ctx context.Context, id_endereco string) ([]models.Empresa, error) {
	return s.repository.GetEmpresasByEndereco(ctx, id_endereco)
}

func (s *service) GetEnderecosByEmpresa(ctx context.Context, id_empresa string) ([]models.Endereco, error) {
	return s.repository.GetEnderecosByEmpresa(ctx, id_empresa)
}
