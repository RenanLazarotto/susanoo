package endereco

import (
	"tsukuyomi/models"
	"tsukuyomi/repositories/endereco"
)

type Service interface {
	Create(model models.Endereco) (models.Endereco, error)
	GetAll(criteria map[string]interface{}) ([]models.Endereco, error)
	GetByID(id string) (models.Endereco, error)
	Update(model models.Endereco) (models.Endereco, error)
}

type service struct {
	repository endereco.Repository
}

func NewService(repository endereco.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (e *service) Create(model models.Endereco) (models.Endereco, error) {
	return e.repository.Create(model)
}

func (e *service) GetAll(criteria map[string]interface{}) ([]models.Endereco, error) {
	return e.repository.GetAll(criteria)
}

func (e *service) GetByID(id string) (models.Endereco, error) {
	return e.repository.GetByID(id)
}

func (e *service) Update(model models.Endereco) (models.Endereco, error) {
	return e.repository.Update(model)
}
