package accounts

import (
	"Go_HSE_2024/2_and_3_HW_server/accounts/dto"
	"Go_HSE_2024/2_and_3_HW_server/accounts/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

func New() *Handler {
	return &Handler{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

type Handler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func (h *Handler) CreateAccount(c echo.Context) error {
	var request dto.CreateAccountRequest // {"name": "alice", "amount": 50}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if request.Name == "" {
		return c.String(http.StatusBadRequest, "empty name")
	}
	if request.Amount < 0 {
		return c.String(http.StatusBadRequest, "invalid amount")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; ok {
		h.guard.Unlock()
		return c.String(http.StatusForbidden, "account already exists")
	}

	h.accounts[request.Name] = &models.Account{
		Name:   request.Name,
		Amount: request.Amount,
	}

	h.guard.Unlock()

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) GetAccount(c echo.Context) error {
	name := c.QueryParams().Get("name")

	h.guard.RLock()

	account, ok := h.accounts[name]

	h.guard.RUnlock()

	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	response := dto.GetAccountResponse{
		Name:   account.Name,
		Amount: account.Amount,
	}

	return c.JSON(http.StatusOK, response)
}

// Удаляет аккаунт
func (h *Handler) DeleteAccount(c echo.Context) error {
	var request dto.DeleteAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if request.Name == "" {
		return c.String(http.StatusBadRequest, "invalid name")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; !ok {
		h.guard.Unlock()
		return c.String(http.StatusForbidden, "account does not exist")
	}

	delete(h.accounts, request.Name)

	h.guard.Unlock()

	return c.NoContent(http.StatusOK)
}

// Меняет баланс
func (h *Handler) PatchAccount(c echo.Context) error {
	var request dto.PatchAccountRequest // {"name": "alice", "amount": 50}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	if request.Name == "" {
		return c.String(http.StatusBadRequest, "empty name")
	}
	if request.Amount < 0 {
		return c.String(http.StatusBadRequest, "invalid amount")
	}

	h.guard.Lock()
	defer h.guard.Unlock()

	account, ok := h.accounts[request.Name]
	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	account.Amount = request.Amount

	return c.NoContent(http.StatusOK)
}

// Меняет имя
func (h *Handler) ChangeAccount(c echo.Context) error {
	var request dto.ChangeAccountRequest // {"old_name": "alice", "new_name": "bob"}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid request")
	}

	if request.OldName == "" || request.NewName == "" {
		return c.String(http.StatusBadRequest, "invalid name")
	}

	h.guard.Lock()
	defer h.guard.Unlock()

	account, ok := h.accounts[request.OldName]
	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	if _, ok := h.accounts[request.NewName]; ok {
		return c.String(http.StatusForbidden, "new account name already exists")
	}

	delete(h.accounts, request.OldName)
	account.Name = request.NewName
	h.accounts[request.NewName] = account

	return c.NoContent(http.StatusOK)
}
