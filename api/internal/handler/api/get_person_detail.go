package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	presenter "github.com/iotassss/fushinsha-map-api/internal/presenter/api"
	"github.com/iotassss/fushinsha-map-api/internal/usecase"
)

type GetPersonDetail struct {
	usecase usecase.GetPersonDetailUsecase
}

func NewGetPersonDetailHandler(
	usecase usecase.GetPersonDetailUsecase,
) *GetPersonDetail {
	return &GetPersonDetail{
		usecase: usecase,
	}
}

func (h *GetPersonDetail) Handle(c *gin.Context) {
	uuid := c.Param("uuid")
	inputData := usecase.GetPersonDetailInputData{
		PersonUUID: uuid,
	}
	presenter := presenter.NewGetPersonDetailPresenter(c)
	if err := h.usecase.Execute(c.Request.Context(), inputData, presenter); err != nil {
		slog.Error("GetPersonDetail.Handle: usecase.Execute", "error", err)
		return
	}
}
