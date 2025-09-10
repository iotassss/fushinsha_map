package presenter

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iotassss/fushinsha-map-api/internal/usecase"
)

// usecase層のエラーをHTTPレスポンスに変換して返す共通関数
func writeErrorResponse(ctx *gin.Context, err error) {
	if errors.Is(err, usecase.ErrValidation) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else if errors.Is(err, usecase.ErrUnauthorized) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	} else if errors.Is(err, usecase.ErrForbidden) {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	} else if errors.Is(err, usecase.ErrNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else if errors.Is(err, usecase.ErrBusinessRule) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	} else if errors.Is(err, usecase.ErrExternal) {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "service temporarily unavailable"})
	} else {
		slog.Error("internal server error", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}
