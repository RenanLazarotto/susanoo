package emprego

import (
	"context"

	"tsukuyomi/models"
	"tsukuyomi/repositories/emprego"
)

type Service interface {
	Create(ctx context.Context, emprego models.Emprego) (models.Emprego, error)
	FindAll(ctx context.Context, search, empresa, ocupacao, remuneracao_inicial, tipo_contrato, data_inicio, data_fim, carga_horaria string) ([]models.Emprego, error)
	FindByID(ctx context.Context, id string) (models.Emprego, error)
	Update(ctx context.Context, emprego models.Emprego) error
	Delete(ctx context.Context, id string) error
}

type service struct {
	repository emprego.Repository
}

func NewService(repository emprego.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, emprego models.Emprego) (models.Emprego, error) {
	return s.repository.Create(ctx, emprego)
}

func (s *service) FindAll(ctx context.Context, search, empresa, ocupacao, remuneracao_inicial, tipo_contrato, data_inicio, data_fim, carga_horaria string) ([]models.Emprego, error) {
	return s.repository.FindAll(ctx, search, empresa, ocupacao, remuneracao_inicial, tipo_contrato, data_inicio, data_fim, carga_horaria)
}

func (s *service) FindByID(ctx context.Context, id string) (models.Emprego, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *service) Update(ctx context.Context, emprego models.Emprego) error {
	return s.repository.Update(ctx, emprego)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
