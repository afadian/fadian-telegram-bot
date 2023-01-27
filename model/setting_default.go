package model

import "github.com/AH-dark/fadian-telegram-bot/internal/env"

var DefaultSettings = []Setting{
	{Key: "system_version", Type: SettingTypeSystem, Value: env.AppVersion},
}
