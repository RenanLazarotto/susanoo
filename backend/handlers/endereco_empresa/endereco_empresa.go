package enderecoempresa

import (
	"github.com/gofiber/fiber/v2"

	"tsukuyomi/models"
	enderecoempresa "tsukuyomi/services/endereco_empresa"
)

type EnderecoEmpresaHandler interface {
	Assign(c *fiber.Ctx) error
	GetEmpresasByEndereco(c *fiber.Ctx) error
	GetEnderecosByEmpresa(c *fiber.Ctx) error
}

type enderecoEmpresaHandler struct {
	Service enderecoempresa.Service
}

var (
	ERROR_ASSIGN               = "Erro ao atribuir endereço à empresa."
	ERROR_EMPRESAS_BY_ENDERECO = "Falha ao consultar empresas pelo endereço."
	ERROR_ENDERECOS_BY_EMPRESA = "Falha ao consultar endereços por empresa."

	ASSIGN_SUCCESS               = "Endereço atribuído à empresa com sucesso."
	EMPRESAS_BY_ENDERECO_SUCCESS = "Consulta de empresas por endereço realizada com sucesso."
	ENDERECOS_BY_EMPRESA_SUCCESS = "Consulta de empresas por endereço realizada com sucesso."

	EMPRESAS_BY_ENDERECO_EMPTY = "Nenhum resultado encontrado para os parâmetros informados."
	ENDERECOS_BY_EMPRESA_EMPTY = "Nenhum resultado encontrado para os parâmetros informados."
)

func NewHandler(service enderecoempresa.Service) EnderecoEmpresaHandler {
	return &enderecoEmpresaHandler{
		Service: service,
	}
}

// Assign godoc
// @Summary     Associa um endereço a uma empresa.
// @Description Faz a associação de um endereço com uma empresa através dos IDs de ambas as entidades.
//
// @Tags    EnderecoEmpresa
// @Accept  json
// @Produce json
//
// @Param id_empresa  body int true "ID da empresa"
// @Param id_endereco body int true "ID do endereço"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /endereco-empresa/assign [post]
func (h *enderecoEmpresaHandler) Assign(c *fiber.Ctx) error {
	dto := models.EndereoEmpresaDTO{}

	c.BodyParser(&dto)
	if dto.IDEmpresa == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_ASSIGN,
			Errors:  []string{"ID da empresa não informado informado."},
		})
	}

	empresa, err := h.Service.GetEmpresaByID(c.UserContext(), dto.IDEmpresa)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_ASSIGN,
			Errors:  []string{err.Error()},
		})
	}

	if dto.IDEndereco == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_ASSIGN,
			Errors:  []string{"ID do endereço não informado informado."},
		})
	}

	endereco, err := h.Service.GetEnderecoByID(c.UserContext(), dto.IDEndereco)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_ASSIGN,
			Errors:  []string{err.Error()},
		})
	}

	enderecoEmpresa, err := h.Service.Assign(c.UserContext(), empresa, endereco)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_ASSIGN,
			Errors:  []string{err.Error()},
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Count:   1,
		Message: ASSIGN_SUCCESS,
		Data:    enderecoEmpresa,
	})
}

// GetEmpresasByEndereco godoc
// @Summary     Retorna as empresas associadas à um endereço.
// @Description Consulta todas as empresas que estão associadas com um endereço, pelo ID do endereço.
//
// @Tags    EnderecoEmpresa
// @Accept  json
// @Produce json
//
// @Param id path int true "ID do endereço"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /endereco-empresa/empresas-por-endereco/{id} [get]
func (h *enderecoEmpresaHandler) GetEmpresasByEndereco(c *fiber.Ctx) error {
	id_endereco := c.Params("id", "")

	result, err := h.Service.GetEmpresasByEndereco(c.UserContext(), id_endereco)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_ENDERECOS_BY_EMPRESA,
			Errors:  []string{err.Error()},
		})
	}

	if len(result) == 0 {
		return c.Status(fiber.StatusOK).JSON(models.Response{
			Message: EMPRESAS_BY_ENDERECO_EMPTY,
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Count:   len(result),
		Message: EMPRESAS_BY_ENDERECO_SUCCESS,
		Data:    result,
	})
}

// GetEnderecosByEmpresa godoc
// @Summary     Retorna os endereços associados à uma empresa
// @Description Consulta todos os endereços que estão associadas com uma empresa, pelo ID da empresa.
//
// @Tags    EnderecoEmpresa
// @Accept  json
// @Produce json
//
// @Param id path int true "ID da empresa"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /endereco-empresa/enderecos-por-empresa/{id} [get]
func (h *enderecoEmpresaHandler) GetEnderecosByEmpresa(c *fiber.Ctx) error {
	id_empresa := c.Params("id", "")

	result, err := h.Service.GetEnderecosByEmpresa(c.UserContext(), id_empresa)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_EMPRESAS_BY_ENDERECO,
			Errors:  []string{err.Error()},
		})
	}

	if len(result) == 0 {
		return c.Status(fiber.StatusOK).JSON(models.Response{
			Message: ENDERECOS_BY_EMPRESA_EMPTY,
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Count:   len(result),
		Message: ENDERECOS_BY_EMPRESA_SUCCESS,
		Data:    result,
	})
}
