package empresa

import (
	"tsukuyomi/models"
	"tsukuyomi/repositories"
)

type Repository interface {
	Create(e models.Empresa) (models.Empresa, error)
}

type repository struct {
	repositories.Repository
}

func NewRepository(repo repositories.Repository) Repository {
	return &repository{
		Repository: repo,
	}
}

func (r *repository) Create(e models.Empresa) (models.Empresa, error) {
	tx := r.DB().BeginTransaction()
	result := tx.Create(&e)

	if result.Error != nil {
		r.DB().Rollback()
		return e, result.Error
	}

	r.DB().Commit()
	return e, nil
}
