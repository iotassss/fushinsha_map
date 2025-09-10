package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	presenter "github.com/iotassss/fushinsha-map-api/internal/presenter/api"
	"github.com/iotassss/fushinsha-map-api/internal/usecase"
)

type RegisterPerson struct {
	usecase usecase.RegisterPersonUsecase
}

func NewRegisterPersonHandler(
	usecase usecase.RegisterPersonUsecase,
) *RegisterPerson {
	return &RegisterPerson{
		usecase: usecase,
	}
}

func (h *RegisterPerson) Handle(c *gin.Context) {
	// TODO: RegisterPersonInputDataのRegisterUUIDは認証情報から取得するようにする
	var input usecase.RegisterPersonInputData
	if err := c.ShouldBindJSON(&input); err != nil {
		slog.Error("RegisterPerson.Handle: invalid input", "error", err)
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}
	presenter := presenter.NewRegisterPersonPresenter(c)
	if err := h.usecase.Execute(c.Request.Context(), input, presenter); err != nil {
		slog.Error("RegisterPerson.Handle: usecase.Execute", "error", err)
		return
	}
}
