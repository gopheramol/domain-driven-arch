package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/gopheramol/domain-driven-arch/pkg/db/orm/sqlc"
)

// AdRepository is the interface that defines the methods for interacting with the ad data.
type AdRepository interface {
	Create(ad sqlc.CreateAdParams) (sqlc.Ad, error)
	GetByID(id int64) (sqlc.Ad, error)
	GetAll() ([]sqlc.Ad, error)
	Update(ad sqlc.UpdateAdParams) error
	Delete(id int64) error
	GetAdByUserID(userID int64) ([]sqlc.Ad, error)
}

// AdRepositoryImpl is the implementation of the AdRepository interface.
type AdRepositoryImpl struct {
	db *sql.DB
}

// NewAdRepository creates a new instance of AdRepositoryImpl with the provided database connection.
func NewAdRepository(db *sql.DB) *AdRepositoryImpl {
	return &AdRepositoryImpl{
		db: db,
	}
}

// Create creates a new ad record in the database.
func (r *AdRepositoryImpl) Create(ad sqlc.CreateAdParams) (sqlc.Ad, error) {
	ctx := context.Background()
	queries := sqlc.New(r.db)

	createdAd, err := queries.CreateAd(ctx, ad)
	if err != nil {
		log.Fatal(err)
	}

	return createdAd, nil
}

// GetByID retrieves an ad record from the database based on its ID.
func (r *AdRepositoryImpl) GetByID(id int64) (sqlc.Ad, error) {
	ctx := context.Background()
	queries := sqlc.New(r.db)

	ad, err := queries.GetAd(ctx, id)
	if err != nil {
		return sqlc.Ad{}, err
	}
	return ad, nil
}

// GetAll retrieves all ad records from the database.
func (r *AdRepositoryImpl) GetAll() ([]sqlc.Ad, error) {
	ctx := context.Background()
	queries := sqlc.New(r.db)

	ads, err := queries.ListAds(ctx)
	if err != nil {
		return []sqlc.Ad{}, err
	}
	return ads, nil
}

// Update updates an existing ad record in the database.
func (r *AdRepositoryImpl) Update(ad sqlc.UpdateAdParams) error {
	ctx := context.Background()
	queries := sqlc.New(r.db)

	err := queries.UpdateAd(ctx, ad)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes an ad record from the database based on its ID.
func (r *AdRepositoryImpl) Delete(id int64) error {
	ctx := context.Background()
	queries := sqlc.New(r.db)

	err := queries.DeleteAd(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// GetAdByUserID retrieves ad records from the database associated with a specific user ID.
func (r *AdRepositoryImpl) GetAdByUserID(userID int64) ([]sqlc.Ad, error) {
	ctx := context.Background()
	queries := sqlc.New(r.db)

	userAds, err := queries.GetAdsByUserID(ctx, userID)
	if err != nil {
		return []sqlc.Ad{}, err
	}
	return userAds, nil
}
