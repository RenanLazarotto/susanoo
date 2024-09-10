package contatoempresa

import (
	"context"

	"tsukuyomi/models"
	contatoempresa "tsukuyomi/repositories/contato_empresa"
)

type Service interface {
	Create(ctx context.Context, contato models.ContatoEmpresa) (models.ContatoEmpresa, error)
	FindAll(ctx context.Context, search, empresa, tipo, contato string) ([]models.ContatoEmpresa, error)
	FindByID(ctx context.Context, id string) (models.ContatoEmpresa, error)
	Update(ctx context.Context, contato models.ContatoEmpresa) error
	Delete(ctx context.Context, id string) error
}

type service struct {
	repository contatoempresa.Repository
}

func NewService(repository contatoempresa.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, contato models.ContatoEmpresa) (models.ContatoEmpresa, error) {
	return s.repository.Create(ctx, contato)
}

func (s *service) FindAll(ctx context.Context, search, empresa, tipo, contato string) ([]models.ContatoEmpresa, error) {
	return s.repository.FindAll(ctx, search, empresa, tipo, contato)
}

func (s *service) FindByID(ctx context.Context, id string) (models.ContatoEmpresa, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *service) Update(ctx context.Context, contato models.ContatoEmpresa) error {
	return s.repository.Update(ctx, contato)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
