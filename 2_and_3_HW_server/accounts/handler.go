package accounts

import (
	"Go_HSE_2024/2_and_3_HW_server/accounts/dto"
	"Go_HSE_2024/2_and_3_HW_server/accounts/models"
	"context"
	"github.com/jackc/pgx/v4"
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

// Создает аккаунт
func (h *Handler) CreateAccount(c echo.Context, conn *pgx.Conn) error {
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
	defer h.guard.Unlock()

	var exists bool
	err := conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM accountss WHERE name=$1)",
		request.Name).Scan(&exists)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "database error")
	}

	if exists {
		return c.String(http.StatusForbidden, "account already exists")
	}

	_, err = conn.Exec(context.Background(), "INSERT INTO accountss (name, balance) VALUES ($1, $2)",
		request.Name, request.Amount)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "database error")
	}

	return c.NoContent(http.StatusCreated)
}

// Показывает данные аккаунта
func (h *Handler) GetAccount(c echo.Context, conn *pgx.Conn) error {
	name := c.QueryParams().Get("name")
	if name == "" {
		return c.String(http.StatusBadRequest, "valid name is required")
	}

	h.guard.RLock()
	defer h.guard.RUnlock()

	var account models.Account
	err := conn.QueryRow(context.Background(), "SELECT name, balance FROM accountss WHERE name=$1", name).Scan(&account.Name, &account.Amount)
	if err != nil {
		if err == pgx.ErrNoRows {
			return c.String(http.StatusNotFound, "account not found")
		}
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "database error")
	}

	response := dto.GetAccountResponse{
		Name:   account.Name,
		Amount: account.Amount,
	}

	return c.JSON(http.StatusOK, response)
}

// Удаляет аккаунт
func (h *Handler) DeleteAccount(c echo.Context, conn *pgx.Conn) error {
	var request dto.DeleteAccountRequest
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if request.Name == "" {
		return c.String(http.StatusBadRequest, "invalid name")
	}

	h.guard.Lock()
	defer h.guard.Unlock()

	var existsd bool
	err := conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM accountss WHERE name=$1)", request.Name).Scan(&existsd)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "database error")
	}
	if !existsd {
		return c.String(http.StatusForbidden, "account does not exist")
	}

	_, err = conn.Exec(context.Background(), "DELETE FROM accountss WHERE name=$1", request.Name)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "database error")
	}

	return c.NoContent(http.StatusOK)
}

// Меняет баланс
func (h *Handler) PatchAccount(c echo.Context, conn *pgx.Conn) error {
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

	var exists bool
	err := conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM accountss WHERE name=$1)",
		request.Name).Scan(&exists)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "database error")
	}
	if !exists {
		return c.String(http.StatusNotFound, "account not found")
	}

	_, err = conn.Exec(context.Background(), "UPDATE accountss SET balance=$1 WHERE name=$2", request.Amount, request.Name)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "database error")
	}

	return c.NoContent(http.StatusOK)
}

// Меняет имя
func (h *Handler) ChangeAccount(c echo.Context, conn *pgx.Conn) error {
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

	var oldAccountExists, newAccountExists bool
	err := conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM accountss WHERE name=$1)", request.OldName).Scan(&oldAccountExists)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "database error")
	}

	if !oldAccountExists {
		return c.String(http.StatusNotFound, "account not found")
	}

	err = conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM accountss WHERE name=$1)", request.NewName).Scan(&newAccountExists)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "database error")
	}

	if newAccountExists {
		return c.String(http.StatusForbidden, "new account name already exists")
	}

	_, err = conn.Exec(context.Background(), "UPDATE accountss SET name=$1 WHERE name=$2", request.NewName, request.OldName)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "database error")
	}

	return c.NoContent(http.StatusOK)
}
