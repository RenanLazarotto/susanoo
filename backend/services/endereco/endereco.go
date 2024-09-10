package endereco

import (
	"context"

	"tsukuyomi/models"
	"tsukuyomi/repositories/endereco"
)

type Service interface {
	Create(ctx context.Context, endereco models.Endereco) (models.Endereco, error)
	FindAll(ctx context.Context, search, logradouro, numero, complemento, bairro, cidade, cep, estado string) ([]models.Endereco, error)
	FindByID(ctx context.Context, id string) (models.Endereco, error)
	Update(ctx context.Context, endereco models.Endereco) error
	Delete(ctx context.Context, id string) error
}

type service struct {
	repository endereco.Repository
}

func NewService(repository endereco.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, endereco models.Endereco) (models.Endereco, error) {
	return s.repository.Create(ctx, endereco)
}

func (s *service) FindAll(ctx context.Context, search, logradouro, numero, complemento, bairro, cidade, cep, estado string) ([]models.Endereco, error) {
	return s.repository.FindAll(ctx, search, logradouro, numero, complemento, bairro, cidade, cep, estado)
}

func (s *service) FindByID(ctx context.Context, id string) (models.Endereco, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *service) Update(ctx context.Context, endereco models.Endereco) error {
	return s.repository.Update(ctx, endereco)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
