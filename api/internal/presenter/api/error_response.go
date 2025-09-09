package presenter

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iotassss/fushinsha-map-api/internal/usecase"
)

// usecase層のエラーをHTTPレスポンスに変換して返す共通関数
func writeErrorResponse(ctx *gin.Context, err error) {
	switch err {
	case usecase.ErrValidation:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case usecase.ErrUnauthorized:
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	case usecase.ErrForbidden:
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	case usecase.ErrNotFound:
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case usecase.ErrBusinessRule:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case usecase.ErrExternal:
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	case usecase.ErrInternal:
		fallthrough
	default:
		slog.Error("internal server error", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}
