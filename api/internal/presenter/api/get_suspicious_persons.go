package presenter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iotassss/fushinsha-map-api/internal/usecase"
)

type GetPersonsPresenter struct {
	ctx *gin.Context
}

func NewGetPersonsPresenter(ctx *gin.Context) *GetPersonsPresenter {
	return &GetPersonsPresenter{
		ctx: ctx,
	}
}

func (p *GetPersonsPresenter) Present(outputData usecase.GetPersonsOutputData) error {
	p.ctx.JSON(http.StatusOK, gin.H{
		"persons": outputData.Persons,
	})
	return nil
}

func (p *GetPersonsPresenter) PresentError(err error) error {
	writeErrorResponse(p.ctx, err)
	return nil
}
