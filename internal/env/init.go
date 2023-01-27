package env

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"

	"github.com/AH-dark/fadian-telegram-bot/pkg/util"
)

var tracer = otel.Tracer("env")

func NewEnvConfig(ctx context.Context) *Config {
	ctx, span := tracer.Start(ctx, "new-env-config")
	defer span.End()

	if err := godotenv.Load(".env", ".env.local"); err != nil {
		logrus.WithContext(ctx).WithError(err).Warning("Failed to load .env files")
	}

	config := &Config{
		System: &system{
			Debug:     util.EnvBool("SYSTEM_DEBUG", false),
			TracerDSN: util.EnvString("SYSTEM_TRACER_DSN", "http://localhost:14268/api/traces"),
		},
		Telegram: &telegram{
			Token: util.EnvString("TELEGRAM_TOKEN", ""),
		},
		Database: &database{
			Type:     util.EnvString("DATABASE_TYPE", "sqlite3"),
			Host:     util.EnvString("DATABASE_HOST", "localhost"),
			Port:     util.EnvInt("DATABASE_PORT", 3306),
			Database: util.EnvString("DATABASE_DATABASE", "fadian"),
			User:     util.EnvString("DATABASE_USER", "root"),
			Password: util.EnvString("DATABASE_PASSWORD", ""),
			Charset:  util.EnvString("DATABASE_CHARSET", "utf8mb4"),
			DBFile:   util.EnvString("DATABASE_DBFILE", "fadian.db"),
		},
		Redis: &redis{
			Network:  util.EnvString("REDIS_NETWORK", "tcp"),
			Host:     util.EnvString("REDIS_HOST", "localhost"),
			Port:     util.EnvInt("REDIS_PORT", 6379),
			Password: util.EnvString("REDIS_PASSWORD", ""),
		},
	}

	return config
}
