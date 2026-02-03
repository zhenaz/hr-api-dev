package repositories

import (
	"context"
	"fmt"

	"codeid.hr-api/internal/models"
	"gorm.io/gorm"
)

// 1. declare interface with methods
type RegionRepository interface {
	FindAll(ctx context.Context) ([]models.Region, error)
	FindByID(ctx context.Context, id uint) (*models.Region, error)
	Create(ctx context.Context, region *models.Region) error
	Update(ctx context.Context, region *models.Region) error
	Delete(ctx context.Context, id uint) error
	GetAllWithCountries() ([]models.Region, error)
	GetRegionWithCountries(id uint) (models.Region, error)
}

// 2. keep regionRepository private (non exported)
// declare field DB with pointer *gorm.DB
type regionRepository struct {
	DB *gorm.DB
}

// 3. pastikan create constructor return-nya interface RegionRepository, jangan struct
// tujuan : agar object RegionRepository bsia akses semua method implementation
// seperti FindAll, FindById, ...
func NewRegionRepository(db *gorm.DB) RegionRepository {
	return &regionRepository{
		DB: db,
	}
}

// Create implements [RegionRepository].
func (r *regionRepository) Create(ctx context.Context, region *models.Region) error {
	return r.DB.WithContext(ctx).Create(region).Error
}

// Delete implements [RegionRepository].
func (r *regionRepository) Delete(ctx context.Context, id uint) error {
	return r.DB.WithContext(ctx).Delete(&models.Region{}, id).Error
}

// FindAll implements [RegionRepository].
func (r *regionRepository) FindAll(ctx context.Context) ([]models.Region, error) {
	var regions []models.Region
	err := r.DB.WithContext(ctx).Find(&regions).Error
	return regions, err
}

// FindByID implements [RegionRepository].
func (r *regionRepository) FindByID(ctx context.Context, id uint) (*models.Region, error) {
	var region models.Region
	err := r.DB.WithContext(ctx).First(&region, id).Error
	if err != nil {
		return nil, err
	}

	return &region, nil
}

// Update implements [RegionRepository].
func (r *regionRepository) Update(ctx context.Context, region *models.Region) error {
	return r.DB.WithContext(ctx).Save(region).Error
}

func (r *regionRepository) GetAllWithCountries() ([]models.Region, error) {
	var regions []models.Region
	err := r.DB.Preload("Countries").Find(&regions).Error
	return regions, err
}

func (r *regionRepository) GetRegionWithCountries(id uint) (models.Region, error) {
	fmt.Println("🔥 GetRegionWithCountries CALLED") // ← LOG DI SINI
	var region models.Region
	err := r.DB.Preload("Countries").
		Where("region_id = ?", id).
		First(&region).Error
	return region, err
}
