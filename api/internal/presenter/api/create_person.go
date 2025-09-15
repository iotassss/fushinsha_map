package presenter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iotassss/fushinsha-map-api/internal/usecase"
)

type CreatePersonPresenter struct {
	ctx *gin.Context
}

func NewCreatePersonPresenter(ctx *gin.Context) *CreatePersonPresenter {
	return &CreatePersonPresenter{
		ctx: ctx,
	}
}

func (p *CreatePersonPresenter) Present(outputData usecase.CreatePersonOutputData) error {
	p.ctx.JSON(http.StatusOK, gin.H{
		"uuid": outputData.UUID,
	})
	return nil
}

func (p *CreatePersonPresenter) PresentError(err error) error {
	writeErrorResponse(p.ctx, err)
	return nil
}
