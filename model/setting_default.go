package model

import "github.com/afadian/fadian-telegram-bot/internal/env"

var DefaultSettings = []Setting{
	{Key: "system_version", Type: SettingTypeSystem, Value: env.AppVersion},
}
