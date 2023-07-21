package service

import (
	"time"

	"github.com/gopheramol/domain-driven-arch/internal/ad/repository"
	"github.com/gopheramol/domain-driven-arch/pkg/db/orm/sqlc"
)

// AdService is the interface that defines the methods for ad-related business logic.
type AdService interface {
	CreateAd(ad sqlc.Ad) (sqlc.Ad, error)
	GetAdByID(id int64) (sqlc.Ad, error)
	GetAllAds() ([]sqlc.Ad, error)
	UpdateAd(id int64, updatedAd sqlc.Ad) error
	DeleteAd(id int64) error
	GetAdByUserID(id int64) ([]sqlc.Ad, error)
}

// AdServiceImpl is the implementation of the AdService interface.
type AdServiceImpl struct {
	adRepo repository.AdRepository
}

// NewAdService creates a new instance of AdServiceImpl with the provided ad repository.
func NewAdService(adRepo repository.AdRepository) *AdServiceImpl {
	return &AdServiceImpl{
		adRepo: adRepo,
	}
}

// CreateAd creates a new ad with the given information.
func (s *AdServiceImpl) CreateAd(ad sqlc.Ad) (sqlc.Ad, error) {
	params := sqlc.CreateAdParams{
		Title:     ad.Title,
		Content:   ad.Content,
		UserID:    ad.UserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdAd, err := s.adRepo.Create(params)
	if err != nil {
		return sqlc.Ad{}, err
	}
	return createdAd, nil
}

// GetAdByID retrieves an ad based on its ID.
func (s *AdServiceImpl) GetAdByID(id int64) (sqlc.Ad, error) {
	ad, err := s.adRepo.GetByID(id)
	if err != nil {
		return sqlc.Ad{}, err
	}
	return ad, nil
}

// GetAllAds retrieves all ads.
func (s *AdServiceImpl) GetAllAds() ([]sqlc.Ad, error) {
	ads, err := s.adRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return ads, nil
}

// UpdateAd updates an existing ad with the given information.
func (s *AdServiceImpl) UpdateAd(id int64, updatedAd sqlc.Ad) error {
	updateAd := sqlc.UpdateAdParams{
		ID:    id,
		Title: updatedAd.Title,
	}

	err := s.adRepo.Update(updateAd)
	if err != nil {
		return err
	}
	return nil
}

// DeleteAd deletes an ad based on its ID.
func (s *AdServiceImpl) DeleteAd(id int64) error {
	err := s.adRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// GetAdByUserID retrieves ads associated with a specific user ID.
func (s *AdServiceImpl) GetAdByUserID(id int64) ([]sqlc.Ad, error) {
	userAds, err := s.adRepo.GetAdByUserID(id)
	if err != nil {
		return []sqlc.Ad{}, err
	}
	return userAds, nil
}
