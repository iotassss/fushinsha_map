package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	presenter "github.com/iotassss/fushinsha-map-api/internal/presenter/api"
	"github.com/iotassss/fushinsha-map-api/internal/usecase"
)

type UpdatePerson struct {
	usecase usecase.UpdatePersonUsecase
}

func NewUpdatePersonHandler(
	usecase usecase.UpdatePersonUsecase,
) *UpdatePerson {
	return &UpdatePerson{
		usecase: usecase,
	}
}

func (h *UpdatePerson) Handle(c *gin.Context) {
	uuid := c.Param("uuid")
	var input usecase.UpdatePersonInputData
	if err := c.ShouldBindJSON(&input); err != nil {
		slog.Error("UpdatePerson.Handle: invalid input", "error", err)
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}
	input.UUID = uuid
	presenter := presenter.NewUpdatePersonPresenter(c)
	if err := h.usecase.Execute(c.Request.Context(), input, presenter); err != nil {
		slog.Error("UpdatePerson.Handle: usecase.Execute", "error", err)
		return
	}
}
