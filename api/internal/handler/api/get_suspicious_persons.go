package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	presenter "github.com/iotassss/fushinsha-map-api/internal/presenter/api"
	"github.com/iotassss/fushinsha-map-api/internal/usecase"
)

type GetSuspiciousPersons struct {
	usecase usecase.GetSuspiciousPersonsUsecase
}

func NewGetSuspiciousPersonsHandler(
	usecase usecase.GetSuspiciousPersonsUsecase,
) *GetSuspiciousPersons {
	return &GetSuspiciousPersons{
		usecase: usecase,
	}
}

func (h *GetSuspiciousPersons) Handle(c *gin.Context) {
	lx := c.Query("lx")
	rx := c.Query("rx")
	ty := c.Query("ty")
	by := c.Query("by")

	inputData := usecase.GetSuspiciousPersonsInputData{
		LX: lx,
		RX: rx,
		TY: ty,
		BY: by,
	}

	presenter := presenter.NewGetSuspiciousPersonsPresenter(c)

	if err := h.usecase.Execute(c.Request.Context(), inputData, presenter); err != nil {
		slog.Error("GetSuspiciousPersons.Handle: usecase.Execute", "error", err)
		return
	}
}
