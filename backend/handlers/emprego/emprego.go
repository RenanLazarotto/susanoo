package emprego

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"tsukuyomi/models"
	"tsukuyomi/services/emprego"
)

type EmpregoHandler interface {
	Create(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	FindByID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type empregoHandler struct {
	Service emprego.Service
}

var (
	ERROR_CREATE   = "Falha ao criar o emprego informada."
	ERROR_FIND_ALL = "Falha ao consultar empregos."
	ERROR_FIND_BY  = "Falha ao consultar emprego por ID."
	ERROR_UPDATE   = "Falha ao atualizar emprego."
	ERROR_DELETE   = "Falha ao apagar o emprego informado."

	CREATE_SUCCESS   = "Emprego criada com sucesso."
	FIND_ALL_SUCCESS = "Consulta realizada com sucesso."
	FIND_BY_SUCCESS  = "Consulta realizada com sucesso."
	UPDATE_SUCCESS   = "Emprego atualizada com sucesso."
	DELETE_SUCCESS   = "Emprego apagado com sucesso."

	FIND_BY_RESULT_EMPTY  = "Nenhum resultado encontrado para os parâmetros informados."
	FIND_ALL_RESULT_EMPTY = "Nenhum resultado encontrado para os parâmetros informados."
)

func NewHandler(service emprego.Service) EmpregoHandler {
	return &empregoHandler{
		Service: service,
	}
}

// Create godoc
// @Summary     Cadastra um novo emprego
// @Description Cadastra um novo emprego de acordo com as informações fornecidas
//
// @Tags    Emprego
// @Accept  json
// @Produce json
//
// @Param id_empresa          body int    true "ID da empresa"
// @Param ocupacao            body string true "Nome da ocupação"
// @Param remuneracao_inicial body number true "Valor da remuneração inicial"
// @Param tipo_contrato       body string true "Tipo de contratação"
// @Param data_inicio         body string true "Data de admissão"
// @Param data_fim            body string true "Data de demissão"
// @Param carga_horaria       body int    true "Carga horária em minutos"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /emprego [post]
func (h *empregoHandler) Create(c *fiber.Ctx) error {
	emprego := models.Emprego{}

	c.BodyParser(&emprego)

	emprego.Criado = time.Now()

	emprego, err := h.Service.Create(c.UserContext(), emprego)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_CREATE,
			Errors:  []string{err.Error()},
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Count:   1,
		Message: CREATE_SUCCESS,
		Data:    emprego,
	})
}

// FindAll godoc
// @Summary     Retorna todos os empregos
// @Description Retorna todos os empregos que atendam aos critérios informados
//
// @Tags    Emprego
// @Accept  json
// @Produce json
//
// @Param search              query string false "Campo aberto para pesquisa"
// @Param empresa             query string false "Nome do empresa"
// @Param ocupacao            query string false "Nome da ocupação"
// @Param remuneracao_inicial query string false "Valor da remuneração inicial"
// @Param tipo_contrato       query string false "Tipo de contratação"
// @Param data_inicio         query string false "Data de admissão"
// @Param data_fim            query string false "Data de demissão"
// @Param carga_horaria       query string false "Carga horária em minutos"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /emprego [get]
func (h *empregoHandler) FindAll(c *fiber.Ctx) error {
	search := c.Query("search", "")
	empresa := c.Query("empresa", "")
	ocupacao := c.Query("ocupacao", "")
	remuneracao_inicial := c.Query("remuneracao_inicial", "")
	tipo_contrato := c.Query("tipo_contrato", "")
	data_inicio := c.Query("data_inicio", "")
	data_fim := c.Query("data_fim", "")
	carga_horaria := c.Query("carga_horaria", "")

	result, err := h.Service.FindAll(c.UserContext(), search, empresa, ocupacao, remuneracao_inicial, tipo_contrato, data_inicio, data_fim, carga_horaria)
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
// @Summary     Consulta um emprego por ID
// @Description Retorna as informações de um emprego de acordo com seu ID
//
// @Tags    Emprego
// @Accept  json
// @Produce json
//
// @Param id path string true "O ID do emprego para retornar"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /emprego/{id} [get]
func (h *empregoHandler) FindByID(c *fiber.Ctx) error {
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
// @Summary     Atualiza um emprego
// @Description Atualiza um registro de emprego de acordo com o ID e as informações informadas
//
// @Tags    Emprego
// @Accept  json
// @Produce json
//
// @Param id                  path string true  "O ID do emprego a ser atualizada"
// @Param id_empresa          body int    false "ID da empresa"
// @Param ocupacao            body string false "Nome da ocupação"
// @Param remuneracao_inicial body number false "Valor da remuneração inicial"
// @Param tipo_contrato       body string false "Tipo de contratação"
// @Param data_inicio         body string false "Data de admissão"
// @Param data_fim            body string false "Data de demissão"
// @Param carga_horaria       body int    false "Carga horária em minutos"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /emprego/{id} [put]
func (h *empregoHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_UPDATE,
			Errors:  []string{"Nenhum ID informado."},
		})
	}

	emprego, err := h.Service.FindByID(c.UserContext(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_UPDATE,
			Errors:  []string{err.Error()},
		})
	}

	c.BodyParser(&emprego)

	now := time.Now()
	emprego.Atualizado = &now

	err = h.Service.Update(c.UserContext(), emprego)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_UPDATE,
			Errors:  []string{err.Error()},
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Count:   1,
		Message: UPDATE_SUCCESS,
		Data:    emprego,
	})
}

// Delete godoc
// @Summary     Apaga umo emprego
// @Description Realiza um soft-delete de umo emprego com base no ID informado
//
// @Tags    Emprego
// @Accept  json
// @Produce json
//
// @Param id path string true "O ID do emprego a ser apagada"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /emprego/{id} [delete]
func (h *empregoHandler) Delete(c *fiber.Ctx) error {
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
