package empresa

import (
	"tsukuyomi/models"
	"tsukuyomi/repositories/empresa"
)

type Service interface {
	Create(model models.Empresa) (models.Empresa, error)
}

type service struct {
	repository empresa.Repository
}

func NewService(repository empresa.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (e *service) Create(model models.Empresa) (models.Empresa, error) {
	return e.repository.Create(model)
}
