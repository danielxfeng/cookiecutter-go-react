package dependency

import (
	"log/slog"
	"time"

	"github.com/getsentry/sentry-go"
)

type Dependency struct {
	Cfg    *Config
	Logger *slog.Logger
}

func NewDependency(cfg *Config, logger *slog.Logger) *Dependency {
	return &Dependency{
		Cfg:    cfg,
		Logger: logger,
	}
}

func InitDependency(level slog.Level) (*Dependency, error) {
	cfg, err := LoadConfigFromEnv()
	if err != nil {
		return nil, err
	}

	err = InitSentry(cfg.AppMode, cfg.SentryDSN)
	if err != nil {
		return nil, err
	}

	logger := GetLogger(level, cfg.AppMode, cfg.SentryDSN)
	return NewDependency(cfg, logger), nil
}

func CloseDependency(dep *Dependency) {
	sentry.Flush(2 * time.Second)
}
