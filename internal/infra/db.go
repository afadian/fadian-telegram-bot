package infra

import (
	"context"
	"fmt"
	"github.com/AH-dark/fadian-telegram-bot/internal/env"
	"github.com/AH-dark/fadian-telegram-bot/model"
	gorm_logrus "github.com/onrik/gorm-logrus"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func InitDB(ctx context.Context, config *env.Config) (*gorm.DB, error) {
	ctx, span := tracer.Start(ctx, "init-db")
	defer span.End()

	logger := logrus.WithContext(ctx)

	var dialector gorm.Dialector
	switch config.Database.Type {
	case "sqlite", "sqlite3":
		dialector = sqlite.Open(config.Database.DBFile)
	case "mysql", "mariadb":
		dialector = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local&sslmode=%s",
			config.Database.User,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Database,
			config.Database.Charset,
			config.Database.SSLMode,
		))
	case "postgres", "postgresql":
		dialector = postgres.Open(fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s TimeZone=%s",
			config.Database.Host,
			config.Database.Port,
			config.Database.User,
			config.Database.Database,
			config.Database.Password,
			config.Database.SSLMode,
			time.Local.String(),
		))
	case "mssql", "sqlserver":
		dialector = sqlserver.Open(fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&sslmode=%s",
			config.Database.User,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Database,
			config.Database.SSLMode,
		))
	default:
		logger.Panicf("Unknown database type: %s", config.Database.Type)
		return nil, fmt.Errorf("unknown database type: %s", config.Database.Type)
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: config.Database.TablePrefix,
		},
		Logger: gorm_logrus.New(),
	})
	if err != nil {
		logger.WithError(err).Panic("Failed to connect to database")
		return nil, err
	}

	db = db.WithContext(ctx) // set context for gorm

	if config.Database.Type == "mysql" || config.Database.Type == "mariadb" {
		db.Set("gorm:table_options", "ENGINE=InnoDB")
	}

	// register open-telemetry plugin
	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		logger.WithError(err).Fatal("Failed to register otelgorm plugin")
		return nil, err
	}

	return db, nil
}

func Migrate(ctx context.Context, db *gorm.DB, forceMigrate bool) error {
	ctx, span := tracer.Start(ctx, "migrate-db")
	defer span.End()

	logger := logrus.WithContext(ctx)
	db = db.WithContext(ctx)

	if !forceMigrate {
		if err := db.Model(&model.Setting{}).Where(&model.Setting{
			Key:   "system_version",
			Type:  model.SettingTypeSystem,
			Value: env.AppVersion,
		}).First(&model.Setting{}).Error; err == nil {
			logger.Infof("Database already migrated to version %s", env.AppVersion)
			return nil
		}
	}

	if err := db.AutoMigrate(
		&model.Setting{},
	); err != nil {
		logger.WithError(err).Fatal("Failed to migrate database")
		return err
	}

	if err := db.Model(&model.Setting{}).Save(&model.Setting{
		Key:   "system_version",
		Type:  model.SettingTypeSystem,
		Value: env.AppVersion,
	}).Error; err != nil {
		logger.WithError(err).Error("Failed to save system version")
		return err
	}

	// migrate settings
	for _, setting := range model.DefaultSettings {
		if err := db.Model(&model.Setting{}).
			Where("key = ?", setting.Key).
			First(&model.Setting{}).
			Error; err == nil {
			continue
		}

		if err := db.Create(&setting).Error; err != nil {
			logger.WithError(err).Errorf("Failed to create setting: %s", setting.Key)
			span.RecordError(err, trace.WithAttributes(
				attribute.String("setting_key", setting.Key),
				attribute.Int("setting_type", int(setting.Type)),
				attribute.String("setting_value", setting.Value),
			))
		}
	}

	return nil
}
