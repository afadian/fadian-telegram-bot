package env

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"

	"github.com/afadian/fadian-telegram-bot/pkg/util"
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
			Listen:    util.EnvString("SYSTEM_LISTEN", ":8080"),
			PublicURL: util.EnvString("SYSTEM_PUBLIC_URL", "http://localhost:8080"),
			CertFile:  util.EnvString("SYSTEM_CERT_FILE", ""),
			KeyFile:   util.EnvString("SYSTEM_KEY_FILE", ""),
		},
		Telegram: &telegram{
			URL:           util.EnvString("TELEGRAM_URL", "https://api.telegram.org"),
			Token:         util.EnvString("TELEGRAM_TOKEN", ""),
			Updates:       util.EnvInt("TELEGRAM_UPDATES", 100),
			PollerTimeout: util.EnvInt("TELEGRAM_POLLER_TIMEOUT", 30),
			Offline:       util.EnvBool("TELEGRAM_OFFLINE", false),
		},
		Database: &database{
			Type:        util.EnvString("DATABASE_TYPE", "sqlite3"),
			Host:        util.EnvString("DATABASE_HOST", "localhost"),
			Port:        util.EnvInt("DATABASE_PORT", 3306),
			Database:    util.EnvString("DATABASE_DATABASE", "fadian"),
			User:        util.EnvString("DATABASE_USER", "root"),
			Password:    util.EnvString("DATABASE_PASSWORD", ""),
			Charset:     util.EnvString("DATABASE_CHARSET", "utf8mb4"),
			SSLMode:     util.EnvString("DATABASE_SSLMODE", "disable"),
			DBFile:      util.EnvString("DATABASE_DBFILE", "fadian.db"),
			TablePrefix: util.EnvString("DATABASE_TABLE_PREFIX", "fadian_"),
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
