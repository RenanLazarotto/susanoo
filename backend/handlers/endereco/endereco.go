package endereco

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"tsukuyomi/models"
	enderecoService "tsukuyomi/services/endereco"
)

type EnderecoHandler interface {
	Create(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	FindByID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type enderecoHandler struct {
	Service enderecoService.Service
}

var (
	ERROR_CREATE   = "Falha ao criar o endereço informado."
	ERROR_FIND_ALL = "Falha ao consultar endereços."
	ERROR_FIND_BY  = "Falha ao consultar endereço."
	ERROR_UPDATE   = "Falha ao atualizar endereço."
	ERROR_DELETE   = "Falha ao apagar o endereço informado."

	CREATE_SUCCESS   = "Endereço criado com sucesso."
	FIND_ALL_SUCCESS = "Consulta realizada com sucesso."
	FIND_BY_SUCCESS  = "Consulta realizada com sucesso."
	UPDATE_SUCCESS   = "Endereço atualizado com sucesso."
	DELETE_SUCCESS   = "Endereço apagado com sucesso."
)

func NewHandler(service enderecoService.Service) EnderecoHandler {
	return &enderecoHandler{
		Service: service,
	}
}

// Create godoc
// @Summary     Cadastra um novo endereço
// @Description Cadastra um novo endereço de acordo com as informações fornecidas
//
// @Tags    Endereco
// @Accept  json
// @Produce json
//
// @Param logradouro  body string true  "Logradouro do endereço"
// @Param numero      body string true  "Número do endereço"
// @Param complemento body string false "Complemento do endereço, caso exista"
// @Param bairro      body string true  "Nome do bairro"
// @Param cidade      body string true  "Nome da cidade"
// @Param cep         body string true  "CEP"
// @Param estado      body string true  "Estado"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /endereco [post]
func (h *enderecoHandler) Create(c *fiber.Ctx) error {
	endereco := models.Endereco{}

	c.BodyParser(&endereco)

	endereco.Criado = time.Now()

	endereco, err := h.Service.Create(c.UserContext(), endereco)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_CREATE,
			Errors:  []string{err.Error()},
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Count:   1,
		Message: CREATE_SUCCESS,
		Data:    endereco,
	})
}

// FindAll godoc
// @Summary     Retorna todos os endereços
// @Description Retorna todos os endereços que atendam aos critérios informados
//
// @Tags    Endereco
// @Accept  json
// @Produce json
//
// @Param search      query string false "Campo aberto para pesquisa"
// @Param logradouro  query string false "Logradouro do endereço"
// @Param numero      query string false "Número do endereço"
// @Param complemento query string false "Complemento do endereço, caso exista"
// @Param bairro      query string false "Nome do bairro"
// @Param cidade      query string false "Nome da cidade"
// @Param cep         query string false "CEP"
// @Param estado      query string false "Estado"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /endereco [get]
func (h *enderecoHandler) FindAll(c *fiber.Ctx) error {
	search := c.Query("search", "")
	logradouro := c.Query("logradouro", "")
	numero := c.Query("numero", "")
	complemento := c.Query("complemento", "")
	bairro := c.Query("bairro", "")
	cidade := c.Query("cidade", "")
	cep := c.Query("cep", "")
	estado := c.Query("estado", "")

	result, err := h.Service.FindAll(c.UserContext(), search, logradouro, numero, complemento, bairro, cidade, cep, estado)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_FIND_ALL,
			Errors:  []string{err.Error()},
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Count:   len(result),
		Message: FIND_ALL_SUCCESS,
		Data:    result,
	})
}

// FindByID godoc
// @Summary     Consulta um endereço por ID
// @Description Retorna as informações de um endereço de acordo com seu ID
//
// @Tags    Endereco
// @Accept  json
// @Produce json
//
// @Param id path string true  "O ID do endereço para retornar"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /endereco/{id} [get]
func (h *enderecoHandler) FindByID(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_UPDATE,
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

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Count:   1,
		Message: FIND_BY_SUCCESS,
		Data:    result,
	})

}

// Update godoc
// @Summary     Atualiza um endereço
// @Description Atualiza um registro de endereço de acordo com o ID e as informações informadas
//
// @Tags    Endereco
// @Accept  json
// @Produce json
//
// @Param id          path string true  "O ID do endereço a ser atualizado"
// @Param logradouro  body string false "Logradouro do endereço"
// @Param numero      body string false "Número do endereço"
// @Param complemento body string false "Complemento do endereço, caso exista"
// @Param bairro      body string false "Nome do bairro"
// @Param cidade      body string false "Nome da cidade"
// @Param cep         body string false "CEP"
// @Param estado      body string false "Estado"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /endereco/{id} [put]
func (h *enderecoHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if id == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_UPDATE,
			Errors:  []string{"Nenhum ID informado."},
		})
	}

	endereco, err := h.Service.FindByID(c.UserContext(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_UPDATE,
			Errors:  []string{err.Error()},
		})
	}

	c.BodyParser(&endereco)

	now := time.Now()
	endereco.Atualizado = &now

	err = h.Service.Update(c.UserContext(), endereco)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: ERROR_UPDATE,
			Errors:  []string{err.Error()},
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Count:   1,
		Message: UPDATE_SUCCESS,
		Data:    endereco,
	})
}

// Delete godoc
// @Summary     Apaga um endereço
// @Description Realiza um soft-delete de um endereço com base no ID informado
//
// @Tags    Endereco
// @Accept  json
// @Produce json
//
// @Param id path string true "O ID do endereço a ser apagado"
//
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
//
// @Router /endereco/{id} [delete]
func (h *enderecoHandler) Delete(c *fiber.Ctx) error {
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
