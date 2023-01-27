package model

import (
	"context"
	"gorm.io/gorm"
)

type SettingRepository interface {
	Find(ctx context.Context, key string) (*Setting, error)
	ListByType(ctx context.Context, settingType SettingType) ([]*Setting, error)
	ListAll(ctx context.Context) ([]*Setting, error)
	Create(ctx context.Context, setting *Setting) error
	Update(ctx context.Context, setting *Setting) error
	Delete(ctx context.Context, setting *Setting) error
}

// SettingService implements SettingRepository
type SettingService struct {
	db *gorm.DB
}

func NewSettingService(db *gorm.DB) SettingRepository {
	return &SettingService{db: db}
}

// Find implements SettingRepository.Find, find a setting by key
func (s *SettingService) Find(ctx context.Context, key string) (*Setting, error) {
	ctx, span := tracer.Start(ctx, "setting-service-find")
	defer span.End()

	var setting Setting
	if err := s.db.WithContext(ctx).Where("key = ?", key).First(&setting).Error; err != nil {
		return nil, err
	}

	return &setting, nil
}

// ListByType implements SettingRepository.ListByType, list settings by type
func (s *SettingService) ListByType(ctx context.Context, t SettingType) ([]*Setting, error) {
	ctx, span := tracer.Start(ctx, "setting-service-list-by-type")
	defer span.End()

	var settings []*Setting
	if err := s.db.WithContext(ctx).Where("type = ?", t).Find(&settings).Error; err != nil {
		return nil, err
	}

	return settings, nil
}

// ListAll implements SettingRepository.ListAll, list all settings
func (s *SettingService) ListAll(ctx context.Context) ([]*Setting, error) {
	ctx, span := tracer.Start(ctx, "setting-service-list-all")
	defer span.End()

	var settings []*Setting
	if err := s.db.WithContext(ctx).Find(&settings).Error; err != nil {
		return nil, err
	}

	return settings, nil
}

// Create implements SettingRepository.Create, create a setting
func (s *SettingService) Create(ctx context.Context, setting *Setting) error {
	ctx, span := tracer.Start(ctx, "setting-service-create")
	defer span.End()

	return s.db.WithContext(ctx).Create(setting).Error
}

// Update implements SettingRepository.Update, update a setting
func (s *SettingService) Update(ctx context.Context, setting *Setting) error {
	ctx, span := tracer.Start(ctx, "setting-service-update")
	defer span.End()

	return s.db.WithContext(ctx).Save(setting).Error
}

// Delete implements SettingRepository.Delete, delete a setting
func (s *SettingService) Delete(ctx context.Context, setting *Setting) error {
	ctx, span := tracer.Start(ctx, "setting-service-delete")
	defer span.End()

	return s.db.WithContext(ctx).Delete(setting).Error
}
