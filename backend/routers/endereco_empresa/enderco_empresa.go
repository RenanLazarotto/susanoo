package enderecoempresa

import (
	"github.com/gofiber/fiber/v2"

	enderecoEmpresaHandler "tsukuyomi/handlers/endereco_empresa"
	"tsukuyomi/repositories"
	"tsukuyomi/repositories/empresa"
	"tsukuyomi/repositories/endereco"
	enderecoEmpresaRepository "tsukuyomi/repositories/endereco_empresa"
	enderecoEmpresaService "tsukuyomi/services/endereco_empresa"
)

func RegisterRoutes(app *fiber.App, repository repositories.Repository) {
	enderecoEmpresaRepository := enderecoEmpresaRepository.NewRepository(repository)
	empresaRepository := empresa.NewRepository(repository)
	enderecoRepository := endereco.NewRepository(repository)

	enderecoService := enderecoEmpresaService.NewService(enderecoEmpresaRepository, empresaRepository, enderecoRepository)

	handler := enderecoEmpresaHandler.NewHandler(enderecoService)

	router := app.Group("/endereco-empresa")
	router.Post("/assign", handler.Assign)
	router.Get("/empresas-por-endereco/:id", handler.GetEmpresasByEndereco)
	router.Get("/enderecos-por-empresa/:id", handler.GetEnderecosByEmpresa)
}
