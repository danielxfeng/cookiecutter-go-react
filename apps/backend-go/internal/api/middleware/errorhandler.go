package middleware

import (
	"errors"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/apps/backend-go/internal/api/apierror"
)

func ErrorHandler(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		errs := c.Errors

		if len(errs) == 0 {
			return
		}

		err := c.Errors.Last().Err

		var apiErr *apierror.ApiError

		// Handle AuthError specifically
		if errors.As(err, &apiErr) {
			c.AbortWithStatusJSON(apiErr.Status, gin.H{
				"error": apiErr.Message,
			})
			return
		}

		var validationErr validator.ValidationErrors

		// Handle validation errors
		if errors.As(err, &validationErr) {
			messages := make([]string, 0, len(validationErr))
			for _, fe := range validationErr {
				messages = append(messages, fe.Error())
			}
			c.AbortWithStatusJSON(400, gin.H{
				"error": messages,
			})
			return
		}

		// Handle other error types or default to 500
		if logger != nil {
			logger.Error("request error", "err", err)
		}
		c.AbortWithStatusJSON(500, gin.H{
			"error": "Internal Server Error",
		})
	}
}

func PanicHandler(logger *slog.Logger) gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered any) {
		if logger != nil {
			logger.Error(
				"panic occurred",
				"panic", recovered,
			)
		}

		c.AbortWithStatusJSON(500, gin.H{
			"error": "Internal Server Error",
		})
	})
}
