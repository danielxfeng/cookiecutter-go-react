package testutil

import (
	"io"
	"log/slog"

	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/apps/backend-go/internal/dependency"
	"github.com/gin-gonic/gin"
)

func NewTestLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{
		Level: slog.LevelError,
	}))
}

func NewTestConfig() *dependency.Config {
	return &dependency.Config{
		AppMode: "test",
		Port:    8080,
		Cors:    "http://localhost:5173",
	}
}

func NewTestDependency(cfg *dependency.Config) *dependency.Dependency {
	if cfg == nil {
		cfg = NewTestConfig()
	}

	return dependency.NewDependency(cfg, NewTestLogger())
}

func NewMiddlewareTestRouter(middleware1 gin.HandlerFunc, middleware2 gin.HandlerFunc) *gin.Engine {
	r := gin.New()

	if middleware1 != nil {
		r.Use(middleware1)
	}

	if middleware2 != nil {
		r.Use(middleware2)
	}

	r.POST("/middleware-test", func(c *gin.Context) {
		userID := c.GetUint("userID")
		token := c.GetString("token")

		c.JSON(200, gin.H{
			"userID": userID,
			"token":  token,
		})
	})

	return r
}
