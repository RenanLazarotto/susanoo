package empresa

import (
	"context"

	"tsukuyomi/models"
	"tsukuyomi/repositories/empresa"
)

type Service interface {
	Create(ctx context.Context, empresa models.Empresa) (models.Empresa, error)
	FindAll(ctx context.Context, search, nome, cnpj string) ([]models.Empresa, error)
	FindByID(ctx context.Context, id string) (models.Empresa, error)
	Update(ctx context.Context, empresa models.Empresa) error
	Delete(ctx context.Context, id string) error
}

type service struct {
	repository empresa.Repository
}

func NewService(repository empresa.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, empresa models.Empresa) (models.Empresa, error) {
	return s.repository.Create(ctx, empresa)
}

func (s *service) FindAll(ctx context.Context, search, nome, cnpj string) ([]models.Empresa, error) {
	return s.repository.FindAll(ctx, search, nome, cnpj)
}

func (s *service) FindByID(ctx context.Context, id string) (models.Empresa, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *service) Update(ctx context.Context, empresa models.Empresa) error {
	return s.repository.Update(ctx, empresa)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
