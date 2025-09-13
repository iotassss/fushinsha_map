package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iotassss/fushinsha-map-api/internal/usecase"
)

type AuthMiddleware struct {
	usecase usecase.FindOrCreateUserByIDTokenUsecase
}

func NewAuthMiddleware(uc usecase.FindOrCreateUserByIDTokenUsecase) *AuthMiddleware {
	return &AuthMiddleware{
		usecase: uc,
	}
}

func (a *AuthMiddleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			return
		}
		input := usecase.FindOrCreateUserByIDTokenInput{
			IDToken: strings.TrimPrefix(header, "Bearer "),
		}
		presenter := NewPresenter(c)

		a.usecase.Execute(c.Request.Context(), input, presenter)

		c.Next()
	}
}
