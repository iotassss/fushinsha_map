package presenter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iotassss/fushinsha-map-api/internal/usecase"
)

type GetPersonDetailPresenter struct {
	ctx *gin.Context
}

func NewGetPersonDetailPresenter(ctx *gin.Context) *GetPersonDetailPresenter {
	return &GetPersonDetailPresenter{
		ctx: ctx,
	}
}

func (p *GetPersonDetailPresenter) Present(outputData usecase.GetPersonDetailOutputData) error {
	p.ctx.JSON(http.StatusOK, gin.H{
		"person": outputData.Person,
	})
	return nil
}

func (p *GetPersonDetailPresenter) PresentError(err error) error {
	writeErrorResponse(p.ctx, err)
	return nil
}
