package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iotassss/fushinsha-map-api/internal/usecase"
)

type Presenter struct {
	ctx *gin.Context
}

func NewPresenter(ctx *gin.Context) *Presenter {
	return &Presenter{ctx: ctx}
}

func (p *Presenter) Present(user usecase.User) error {
	p.ctx.Set("user", user)
	return nil
}

func (p *Presenter) PresentError(err error) error {
	p.ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	return err
}
