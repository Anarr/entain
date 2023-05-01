package handler

import (
	"errors"
	"fmt"
	"github.com/Anarr/entain/internal/manager"
	"github.com/Anarr/entain/internal/middleware"
	"github.com/Anarr/entain/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

const (
	pathProcessRequest = "/process"
	defaultUserID      = 1
)

type ProcessRequest struct {
	State         string  `json:"state"`
	TransactionID string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
}

func (pr *ProcessRequest) validate() error {

	if pr.Amount <= 0 {
		return errors.New("amount should be greater than 0")
	}

	if pr.State != "win" && pr.State != "lost" {
		return fmt.Errorf("unsupported state: %s", pr.State)
	}

	if strings.TrimSpace(pr.TransactionID) == "" {
		return errors.New("transaction_id is required")
	}

	return nil
}

type Handler struct {
	manager manager.Manager
}

func New(manager manager.Manager) *Handler {
	return &Handler{manager: manager}
}

func (h Handler) RegisterRoutes(group *echo.Group) {
	group.POST(pathProcessRequest, h.ProcessRequestHandler, middleware.CheckSourceHeader)
}

// Entain godoc
// @Summary Process new incoming requests
//
// @Param requestData body handler.processRequest true "body"
// @Param Source-Type header string true "Source-Type enums" Enums(game, server, payment, invalid-header)
//
// @Failure 400 {object} model.APIError
// @Failure 422 {object} model.APIError
// @Failure 500 {object} model.APIError
//
// @Tags [request]
//
// @Router /api/process [post]
func (h Handler) ProcessRequestHandler(c echo.Context) error {
	pr := new(ProcessRequest)

	if err := c.Bind(pr); err != nil {
		return model.APIError{StatusCode: http.StatusBadRequest, Err: "unsupported user request"}
	}

	if err := pr.validate(); err != nil {
		return model.APIError{StatusCode: http.StatusUnprocessableEntity, Err: err.Error()}
	}

	m := model.Request{
		UserID:        defaultUserID,
		TransactionID: pr.TransactionID,
		Amount:        pr.Amount,
		State:         pr.State,
	}
	if err := h.manager.ProcessRequest(m); err != nil {
		return model.APIError{StatusCode: http.StatusInternalServerError, Err: "something went wrong"}
	}

	return c.JSON(http.StatusOK, "Processed")
}
