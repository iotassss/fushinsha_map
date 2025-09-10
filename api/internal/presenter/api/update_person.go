package presenter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iotassss/fushinsha-map-api/internal/usecase"
)

type UpdatePersonPresenter struct {
	ctx *gin.Context
}

func NewUpdatePersonPresenter(ctx *gin.Context) *UpdatePersonPresenter {
	return &UpdatePersonPresenter{
		ctx: ctx,
	}
}

func (p *UpdatePersonPresenter) Present(outputData usecase.UpdatePersonOutputData) error {
	p.ctx.JSON(http.StatusOK, gin.H{
		"success": outputData.Success,
	})
	return nil
}

func (p *UpdatePersonPresenter) PresentError(err error) error {
	writeErrorResponse(p.ctx, err)
	return nil
}
