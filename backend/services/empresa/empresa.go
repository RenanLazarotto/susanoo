package empresa

import (
	"tsukuyomi/models"
	"tsukuyomi/repositories/empresa"
)

type Service interface {
	Create(model models.Empresa) (models.Empresa, error)
	GetAll(criteria map[string]interface{}) ([]models.Empresa, error)
	GetByID(id string) (models.Empresa, error)
	Update(model models.Empresa) (models.Empresa, error)
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

func (e *service) GetAll(criteria map[string]interface{}) ([]models.Empresa, error) {
	return e.repository.GetAll(criteria)
}

func (e *service) GetByID(id string) (models.Empresa, error) {
	return e.repository.GetByID(id)
}

func (e *service) Update(model models.Empresa) (models.Empresa, error) {
	return e.repository.Update(model)
}
