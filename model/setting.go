package model

import "gorm.io/gorm"

type Setting struct {
	gorm.Model
	Key   string      `gorm:"uniqueIndex;not null"`
	Type  SettingType `gorm:"not null;default:0"`
	Value string      `gorm:"not null;default:''"`
}

type SettingType int

const (
	SettingTypeSystem SettingType = iota
)
