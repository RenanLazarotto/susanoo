package endereco

import (
	"tsukuyomi/models"
	"tsukuyomi/repositories"
)

type Repository interface {
	Create(e models.Endereco) (models.Endereco, error)
	GetAll(criteria map[string]interface{}) ([]models.Endereco, error)
	GetByID(id string) (models.Endereco, error)
	Update(e models.Endereco) (models.Endereco, error)
}

type repository struct {
	repositories.Repository
}

func NewRepository(repo repositories.Repository) Repository {
	return &repository{
		Repository: repo,
	}
}

func (r *repository) Create(e models.Endereco) (models.Endereco, error) {
	tx := r.DB().BeginTransaction()
	result := tx.Create(&e)

	if result.Error != nil {
		r.DB().Rollback()
		return e, result.Error
	}

	r.DB().Commit()
	return e, nil
}

func (r *repository) GetAll(criteria map[string]interface{}) ([]models.Endereco, error) {
	empresas := []models.Endereco{}
	q, err := r.DB().Query()
	if err != nil {
		return empresas, err
	}
	result := q.Where(criteria).Find(&empresas)

	if result.Error != nil {
		return empresas, result.Error
	}

	return empresas, nil
}

func (r *repository) GetByID(id string) (models.Endereco, error) {
	empresa := models.Endereco{}
	q, err := r.DB().Query()
	if err != nil {
		return empresa, err
	}
	result := q.Where("id = ?", id).First(&empresa)

	if result.Error != nil {
		return empresa, result.Error
	}

	return empresa, nil
}

func (r *repository) Update(e models.Endereco) (models.Endereco, error) {
	tx := r.DB().BeginTransaction()

	result := tx.Model(&e).Updates(&e)

	if result.Error != nil {
		r.DB().Rollback()
		return e, result.Error
	}

	r.DB().Commit()
	return e, nil
}
