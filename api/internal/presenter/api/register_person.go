package presenter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iotassss/fushinsha-map-api/internal/usecase"
)

type RegisterPersonPresenter struct {
	ctx *gin.Context
}

func NewRegisterPersonPresenter(ctx *gin.Context) *RegisterPersonPresenter {
	return &RegisterPersonPresenter{
		ctx: ctx,
	}
}

func (p *RegisterPersonPresenter) Present(outputData usecase.RegisterPersonOutputData) error {
	p.ctx.JSON(http.StatusOK, gin.H{
		"uuid": outputData.UUID,
	})
	return nil
}

func (p *RegisterPersonPresenter) PresentError(err error) error {
	writeErrorResponse(p.ctx, err)
	return nil
}
