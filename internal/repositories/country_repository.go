package repositories

import (
	"context"

	"codeid.hr-api/internal/models"
	"gorm.io/gorm"
)

// 1. declare interface with methods
type CountryRepository interface {
	FindAll(ctx context.Context) ([]models.Country, error)
	FindByID(ctx context.Context, id string) (*models.Country, error) // 🔥 STRING
	Create(ctx context.Context, country *models.Country) error
	Update(ctx context.Context, country *models.Country) error
	Delete(ctx context.Context, id string) error // 🔥 STRING
}

// 2. keep countryRepository private (non exported)
// declare field DB with pointer *gorm.DB
type countryRepository struct {
	DB *gorm.DB
}

// 3. pastikan create constructor return-nya interface CountryRepository, jangan struct
// tujuan : agar object CountryRepository bsia akses semua method implementation
// seperti FindAll, FindById, ...
func NewCountryRepository(db *gorm.DB) CountryRepository {
	return &countryRepository{
		DB: db,
	}
}

// Create implements [CountryRepository].
func (r *countryRepository) Create(ctx context.Context, country *models.Country) error {
	return r.DB.WithContext(ctx).Create(country).Error
}

// Delete implements [CountryRepository].
func (r *countryRepository) Delete(ctx context.Context, id string) error {
	return r.DB.WithContext(ctx).Delete(&models.Country{}, "id = ?", id).Error
}

// FindAll implements [CountryRepository].
func (r *countryRepository) FindAll(ctx context.Context) ([]models.Country, error) {
	var countries []models.Country
	err := r.DB.WithContext(ctx).Find(&countries).Error
	return countries, err
}

// FindByID implements [CountryRepository].
func (r *countryRepository) FindByID(ctx context.Context, id string) (*models.Country, error) {
	var country models.Country
	err := r.DB.WithContext(ctx).First(&country, "country_id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &country, nil
}

// Update implements [CountryRepository].
func (r *countryRepository) Update(ctx context.Context, country *models.Country) error {
	return r.DB.WithContext(ctx).Save(country).Error
}
