package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	presenter "github.com/iotassss/fushinsha-map-api/internal/presenter/api"
	"github.com/iotassss/fushinsha-map-api/internal/usecase"
)

type GetPersons struct {
	usecase usecase.GetPersonsUsecase
}

func NewGetPersonsHandler(
	usecase usecase.GetPersonsUsecase,
) *GetPersons {
	return &GetPersons{
		usecase: usecase,
	}
}

func (h *GetPersons) Handle(c *gin.Context) {
	lx := c.Query("lx")
	rx := c.Query("rx")
	ty := c.Query("ty")
	by := c.Query("by")

	inputData := usecase.GetPersonsInputData{
		LX: lx,
		RX: rx,
		TY: ty,
		BY: by,
	}

	presenter := presenter.NewGetPersonsPresenter(c)

	if err := h.usecase.Execute(c.Request.Context(), inputData, presenter); err != nil {
		slog.Error("GetPersons.Handle: usecase.Execute", "error", err)
		return
	}
}
