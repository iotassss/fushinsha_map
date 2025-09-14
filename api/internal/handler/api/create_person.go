package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	presenter "github.com/iotassss/fushinsha-map-api/internal/presenter/api"
	"github.com/iotassss/fushinsha-map-api/internal/usecase"
)

type CreatePerson struct {
	usecase usecase.CreatePersonUsecase
}

func NewCreatePersonHandler(
	usecase usecase.CreatePersonUsecase,
) *CreatePerson {
	return &CreatePerson{
		usecase: usecase,
	}
}

func (h *CreatePerson) Handle(c *gin.Context) {
	// TODO: CreatePersonInputDataのCreateUUIDは認証情報から取得するようにする
	var input usecase.CreatePersonInputData
	if err := c.ShouldBindJSON(&input); err != nil {
		slog.Error("CreatePerson.Handle: invalid input", "error", err)
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}
	presenter := presenter.NewCreatePersonPresenter(c)
	if err := h.usecase.Execute(c.Request.Context(), input, presenter); err != nil {
		slog.Error("CreatePerson.Handle: usecase.Execute", "error", err)
		return
	}
}
