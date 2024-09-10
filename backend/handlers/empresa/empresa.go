package empresa

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"tsukuyomi/models"
	empresaService "tsukuyomi/services/empresa"
)

type EmpresaHandler interface {
	Create(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	FindByID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type empresaHandler struct {
	Service empresaService.Service
}

var (
	ERROR_CREATE   = "Falha ao criar a empresa informada."
	ERROR_FIND_ALL = "Falha ao consultar empresas."
	ERROR_FIND_BY  = "Falha ao consultar empresa por ID."
	ERROR_UPDATE   = "Falha ao atualizar empresa."
	ERROR_DELETE   = "Falha ao apagar a empresa informado."

	CREATE_SUCCESS   = "Empresa criada com sucesso."
	FIND_ALL_SUCCESS = "Consulta realizada com sucesso."
	FIND_BY_SUCCESS  = "Consulta realizada com sucesso."
	UPDATE_SUCCESS   = "Empresa atualizada com sucesso."
	DELETE_SUCCESS   = "Empresa apagada com sucesso."

	FIND_BY_RESULT_EMPTY  = "Nenhum resultado encontrado para os parâmetros informados."
	FIND_ALL_RESULT_EMPTY = "Nenhum resultado encontrado para os parâmetros informados."
)

func NewHandler(service empresaService.Service) EmpresaHandler {
	return &empresaHandler{
		Service: service,
	}
}

// Create godoc
// @Summary     Cadastra um nova empresa
// @Description Cadastra um nova empresa de acordo com as informações fornecidas
//
// @Tags    Empresa
// @Accept  json
// @Produce json
//
// @Param nome body string true "Nome da empresa"
// @Param cnpj body string true "CNPJ da empresa"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /empresa [post]
func (h *empresaHandler) Create(c *fiber.Ctx) error {
	empresa := models.Empresa{}

	c.BodyParser(&empresa)

	empresa.Criado = time.Now()

	empresa, err := h.Service.Create(c.UserContext(), empresa)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_CREATE,
			Errors:  []string{err.Error()},
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Count:   1,
		Message: CREATE_SUCCESS,
		Data:    empresa,
	})
}

// FindAll godoc
// @Summary     Retorna todos as empresas
// @Description Retorna todos as empresas que atendam aos critérios informados
//
// @Tags    Empresa
// @Accept  json
// @Produce json
//
// @Param search query string false "Campo aberto para pesquisa"
// @Param nome   query string false "Nome da empresa"
// @Param cnpj   query string false "CNPJ da empresa"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /empresa [get]
func (h *empresaHandler) FindAll(c *fiber.Ctx) error {
	search := c.Query("search", "")
	nome := c.Query("nome", "")
	cnpj := c.Query("cnpj", "")

	result, err := h.Service.FindAll(c.UserContext(), search, nome, cnpj)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_FIND_ALL,
			Errors:  []string{err.Error()},
		})
	}

	if len(result) == 0 {
		return c.Status(fiber.StatusOK).JSON(models.Response{
			Message: FIND_ALL_RESULT_EMPTY,
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Count:   len(result),
		Message: FIND_ALL_SUCCESS,
		Data:    result,
	})
}

// FindByID godoc
// @Summary     Consulta uma empresa por ID
// @Description Retorna as informações de uma empresa de acordo com seu ID
//
// @Tags    Empresa
// @Accept  json
// @Produce json
//
// @Param id path string true "O ID da empresa para retornar"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /empresa/{id} [get]
func (h *empresaHandler) FindByID(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_FIND_BY,
			Errors:  []string{"Nenhum ID informado."},
		})
	}

	result, err := h.Service.FindByID(c.UserContext(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_FIND_BY,
			Errors:  []string{err.Error()},
		})
	}

	if result.ID == 0 {
		return c.Status(fiber.StatusOK).JSON(models.Response{
			Message: FIND_BY_RESULT_EMPTY,
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Count:   1,
		Message: FIND_BY_SUCCESS,
		Data:    result,
	})
}

// Update godoc
// @Summary     Atualiza uma empresa
// @Description Atualiza um registro de empresa de acordo com o ID e as informações informadas
//
// @Tags    Empresa
// @Accept  json
// @Produce json
//
// @Param id   path string true  "O ID da empresa a ser atualizada"
// @Param nome body string false "Nome da empresa"
// @Param cnpj body string false "CNPJ da empresa"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /empresa/{id} [put]
func (h *empresaHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_UPDATE,
			Errors:  []string{"Nenhum ID informado."},
		})
	}

	empresa, err := h.Service.FindByID(c.UserContext(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_UPDATE,
			Errors:  []string{err.Error()},
		})
	}

	c.BodyParser(&empresa)

	now := time.Now()
	empresa.Atualizado = &now

	err = h.Service.Update(c.UserContext(), empresa)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_UPDATE,
			Errors:  []string{err.Error()},
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Count:   1,
		Message: UPDATE_SUCCESS,
		Data:    empresa,
	})
}

// Delete godoc
// @Summary     Apaga uma empresa
// @Description Realiza um soft-delete de uma empresa com base no ID informado
//
// @Tags    Empresa
// @Accept  json
// @Produce json
//
// @Param id path string true "O ID da empresa a ser apagada"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /empresa/{id} [delete]
func (h *empresaHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_DELETE,
			Errors:  []string{"Nenhum ID informado."},
		})
	}

	err := h.Service.Delete(c.UserContext(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_DELETE,
			Errors:  []string{err.Error()},
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Message: DELETE_SUCCESS,
	})
}
