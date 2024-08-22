package empresa

import (
	"github.com/charmbracelet/log"

	"tsukuyomi/repositories"
)

type Repository interface {
	TestConnection()
}

type repository struct {
	repositories.Repository
}

func NewRepository(repo repositories.Repository) Repository {
	return &repository{
		Repository: repo,
	}
}

func (r *repository) TestConnection() {
	//r.DB().BeginTransaction()
	db, err := r.DB().Query()

	if err != nil {
		log.Fatal(err)
	}

	result := struct {
		One string
		Two int
	}{}

	db.Raw("SELECT 'string' as One, 1234 as Two").Scan(&result)

	log.Info(result)
}
