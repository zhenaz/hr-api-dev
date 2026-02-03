package services

import (
	"context"
	"errors"

	"codeid.hr-api/internal/models"
	"codeid.hr-api/internal/repositories"
)

// 4. define implementation methods, use tombol Ctrl + . saat kursor ada di
// NewRegionRepository, supy bisa generate all method interface.
type CountryService interface {
	GetAllCountries(ctx context.Context) ([]models.Country, error)
	GetCountryByID(ctx context.Context, id string) (*models.Country, error)
	CreateCountry(ctx context.Context, country *models.Country) error
	UpdateCountry(ctx context.Context, country *models.Country) error
	DeleteCountry(ctx context.Context, id string) error
}

type countryService struct {
	countryRepo repositories.CountryRepository
}

func NewCountryService(countryRepo repositories.CountryRepository) CountryService {
	return &countryService{
		countryRepo: countryRepo,
	}
}
func (s *countryService) GetAllCountries(ctx context.Context) ([]models.Country,
	error) {
	return s.countryRepo.FindAll(ctx)
}
func (s *countryService) GetCountryByID(ctx context.Context, id string) (*models.Country, error) {
	if id == "" {
		return nil, errors.New("country ID cannot be empty")
	}
	return s.countryRepo.FindByID(ctx, id)
}

func (s *countryService) CreateCountry(ctx context.Context, country *models.Country) error {
	if country.CountryName == "" {
		return errors.New("country name cannot be empty")
	}
	if len(country.CountryName) > 25 {
		return errors.New("country name cannot exceed 25 characters")
	}
	return s.countryRepo.Create(ctx, country)
}
func (s *countryService) UpdateCountry(ctx context.Context, country *models.Country) error {
	if country.CountryID == "" {
		return errors.New("country ID cannot be empty")
	}
	if country.CountryName == "" {
		return errors.New("country name cannot be empty")
	}
	if len(country.CountryName) > 40 { // DB kamu varchar(40)
		return errors.New("country name cannot exceed 40 characters")
	}

	if _, err := s.countryRepo.FindByID(ctx, country.CountryID); err != nil {
		return errors.New("country doesn't exist")
	}

	return s.countryRepo.Update(ctx, country)
}

func (s *countryService) DeleteCountry(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("country ID cannot be empty")
	}
	if _, err := s.countryRepo.FindByID(ctx, id); err != nil {
		return errors.New("country doesn't exists")
	}
	return s.countryRepo.Delete(ctx, id)
}
