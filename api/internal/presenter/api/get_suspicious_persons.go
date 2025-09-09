package presenter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iotassss/fushinsha-map-api/internal/usecase"
)

type GetSuspiciousPersonsPresenter struct {
	ctx *gin.Context
}

func NewGetSuspiciousPersonsPresenter(ctx *gin.Context) *GetSuspiciousPersonsPresenter {
	return &GetSuspiciousPersonsPresenter{
		ctx: ctx,
	}
}

func (p *GetSuspiciousPersonsPresenter) Present(outputData usecase.GetSuspiciousPersonsOutputData) error {
	p.ctx.JSON(http.StatusOK, gin.H{
		"persons": outputData.Persons,
	})
	return nil
}

func (p *GetSuspiciousPersonsPresenter) PresentError(err error) error {
	writeErrorResponse(p.ctx, err)
	return nil
}
