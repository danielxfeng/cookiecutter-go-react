package routers

import (
	"log/slog"
	"time"

	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/apps/backend-go/internal/api/middleware"
	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/apps/backend-go/internal/dependency"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func InitRouter(dep *dependency.Dependency) *gin.Engine {
	r := gin.New()

	r.Use(middleware.PanicHandler(dep.Logger))

	logConfig := sloggin.Config{
		DefaultLevel:     slog.LevelInfo,
		ClientErrorLevel: slog.LevelWarn,
		ServerErrorLevel: slog.LevelError,
	}

	if dep.Cfg.AppMode == "debug" || dep.Cfg.AppMode == "test" || dep.Cfg.AppMode == "ci" {
		r.Use(cors.Default())
	} else {
		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{dep.Cfg.Cors},
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	}

	r.Use(sloggin.NewWithConfig(dep.Logger, logConfig))
	r.Use(middleware.ErrorHandler(dep.Logger))

	return r
}
