package contatoempresa

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"tsukuyomi/models"
	contatoempresa "tsukuyomi/services/contato_empresa"
)

type ContatoEmpresaHandler interface {
	Create(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	FindByID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type contatoEmpresaHandler struct {
	Service contatoempresa.Service
}

var (
	ERROR_CREATE   = "Falha ao criar o contato informado."
	ERROR_FIND_ALL = "Falha ao consultar contatos."
	ERROR_FIND_BY  = "Falha ao consultar contato por ID."
	ERROR_UPDATE   = "Falha ao atualizar contato."
	ERROR_DELETE   = "Falha ao apagar o contato informado."

	CREATE_SUCCESS   = "Contato criado com sucesso."
	FIND_ALL_SUCCESS = "Consulta realizada com sucesso."
	FIND_BY_SUCCESS  = "Consulta realizada com sucesso."
	UPDATE_SUCCESS   = "Contato atualizado com sucesso."
	DELETE_SUCCESS   = "Contato apagado com sucesso."

	FIND_BY_RESULT_EMPTY  = "Nenhum resultado encontrado para os parâmetros informados."
	FIND_ALL_RESULT_EMPTY = "Nenhum resultado encontrado para os parâmetros informados."
)

func NewHandler(service contatoempresa.Service) ContatoEmpresaHandler {
	return &contatoEmpresaHandler{
		Service: service,
	}
}

// Create godoc
// @Summary     Cadastra um novo contato de empresa
// @Description Cadastra um novo contato de empresa de acordo com as informações fornecidas. O contato deve ser único de acordo com seu tipo.
// @Description Ao cadastrar um contato, os dados da empresa não são retornados.
//
// @Tags    ContatoEmpresa
// @Accept  json
// @Produce json
//
// @Param id_empresa body int    true "ID da empresa"
// @Param tipo       body string true "O tipo de contato. Aceita apenas os valores 'telefone', 'whatsapp' e 'email'" Enums(telefone, whatsapp, email)
// @Param contato    body string true "O contato em si"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /contato-empresa [post]
func (h *contatoEmpresaHandler) Create(c *fiber.Ctx) error {
	contatoEmpresa := models.ContatoEmpresa{}

	c.BodyParser(&contatoEmpresa)

	contatoEmpresa.Criado = time.Now()

	contatoEmpresa, err := h.Service.Create(c.UserContext(), contatoEmpresa)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_CREATE,
			Errors:  []string{err.Error()},
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Count:   1,
		Message: CREATE_SUCCESS,
		Data:    contatoEmpresa,
	})
}

// FindAll godoc
// @Summary     Retorna todos os contatos
// @Description Retorna todos os contatos que atendam aos critérios informados
//
// @Tags    ContatoEmpresa
// @Accept  json
// @Produce json
//
// @Param search  query string false "Campo aberto para pesquisa"
// @Param empresa query string false "Nome da empresa"
// @Param tipo    query string false "Tipo do contato. Aceita apenas os valores 'telefone', 'whatsapp' e 'email'" Enums(telefone, whatsapp, email)
// @Param contato query string false "O contato em si"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /contato-empresa [get]
func (h *contatoEmpresaHandler) FindAll(c *fiber.Ctx) error {
	search := c.Query("search", "")
	empresa := c.Query("empresa", "")
	tipo := c.Query("tipo", "")
	contato := c.Query("contato", "")

	result, err := h.Service.FindAll(c.UserContext(), search, empresa, tipo, contato)
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
// @Summary     Consulta um contato de empresa por ID
// @Description Retorna as informações de um contato de uma empresa de acordo com seu ID
//
// @Tags    ContatoEmpresa
// @Accept  json
// @Produce json
//
// @Param id path string true "O ID do contato da empresa para retornar"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /contato-empresa/{id} [get]
func (h *contatoEmpresaHandler) FindByID(c *fiber.Ctx) error {
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
// @Summary     Atualiza um contato de empresa
// @Description Atualiza um registro de contato de uma empresa de acordo com o ID e as informações informadas
//
// @Tags    ContatoEmpresa
// @Accept  json
// @Produce json
//
// @Param id         path  string true  "O ID do contato de empresa a ser atualizado"
// @Param id_empresa query int    false "ID da empresa para atualizar"
// @Param tipo       query string false "Tipo do contato. Aceita apenas os valores 'telefone', 'whatsapp' e 'email'" Enums(telefone, whatsapp, email)
// @Param contato    query string false "O contato em si"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /contato-empresa/{id} [put]
func (h *contatoEmpresaHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_UPDATE,
			Errors:  []string{"Nenhum ID informado."},
		})
	}

	contato, err := h.Service.FindByID(c.UserContext(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_UPDATE,
			Errors:  []string{err.Error()},
		})
	}

	c.BodyParser(&contato)

	now := time.Now()
	contato.Atualizado = &now

	err = h.Service.Update(c.UserContext(), contato)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_UPDATE,
			Errors:  []string{err.Error()},
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Count:   1,
		Message: UPDATE_SUCCESS,
		Data:    contato,
	})
}

// Delete godoc
// @Summary     Apaga um contato de empresa
// @Description Realiza um soft-delete de um contato de empresa com base no ID informado
//
// @Tags    ContatoEmpresa
// @Accept  json
// @Produce json
//
// @Param id path string true "O ID do contato a ser apagada"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /contato-empresa/{id} [delete]
func (h *contatoEmpresaHandler) Delete(c *fiber.Ctx) error {
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
