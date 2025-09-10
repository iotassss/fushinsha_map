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
	topErr := err
	for {
		unwrapped := errors.Unwrap(topErr)
		if unwrapped == nil {
			break
		}
		topErr = unwrapped
	}

	if errors.Is(topErr, usecase.ErrValidation) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
	} else if errors.Is(topErr, usecase.ErrUnauthorized) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
	} else if errors.Is(topErr, usecase.ErrForbidden) {
		ctx.JSON(http.StatusForbidden, gin.H{"error": http.StatusText(http.StatusForbidden)})
	} else if errors.Is(topErr, usecase.ErrNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": http.StatusText(http.StatusNotFound)})
	} else if errors.Is(topErr, usecase.ErrBusinessRule) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": http.StatusText(http.StatusUnprocessableEntity)})
	} else if errors.Is(topErr, usecase.ErrExternal) {
		slog.Error("bad gateway", "error", topErr)
		ctx.JSON(http.StatusBadGateway, gin.H{"error": http.StatusText(http.StatusBadGateway)})
	} else {
		slog.Error("internal server error", "error", topErr)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
	}
}
