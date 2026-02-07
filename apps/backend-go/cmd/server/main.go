package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/apps/backend-go/internal/api/routers"
	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.project_slug }}/apps/backend-go/internal/dependency"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func startApiServer(ctx context.Context, dep *dependency.Dependency) {
	// init router
	r := routers.InitRouter(dep)

	r.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Swagger
	r.GET("/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", dep.Cfg.Port),
		Handler: r,
	}

	// Start server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			dependency.LogFatalErr(dep.Logger, err, "failed to start api service")
		}
	}()

	// Graceful shutdown
	<-ctx.Done() // consume the signal, blocking here
	dep.Logger.Info("shutting down service...")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctxShutdown); err != nil {
		dependency.LogFatalErr(dep.Logger, err, "server forced to shutdown")
	}

	dep.Logger.Info("api service exiting")
}

func main() {
	_ = godotenv.Load()

	// init dependency
	dep, err := dependency.InitDependency(slog.LevelDebug)
	if err != nil {
		log.Fatalf("failed to init the dependency, err: %v", err)
	}
	defer dependency.CloseDependency(dep)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	var wg sync.WaitGroup
	wg.Go(func() {
		startApiServer(ctx, dep)
	})

	wg.Wait()
	dep.Logger.Info("server exiting")
}
