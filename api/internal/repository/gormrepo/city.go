package gormrepo

import (
	"context"

	"gorm.io/gorm"
)

type CityModel struct {
	gorm.Model
	ID   string `gorm:"column:id;primaryKey;not null"`
	Name string `gorm:"column:name;not null"`
}

func (CityModel) TableName() string {
	return "cities"
}

type CityRepository struct {
	db  *gorm.DB
	ctx context.Context
}

func NewCityRepository(
	db *gorm.DB,
	ctx context.Context,
) *CityRepository {
	return &CityRepository{
		db:  db,
		ctx: ctx,
	}
}

func (r *CityRepository) FindAll() ([]*CityModel, error) {
	var models []*CityModel
	if err := r.db.WithContext(r.ctx).Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func (r *CityRepository) ResetTable() error {
	if err := r.db.WithContext(r.ctx).Exec("TRUNCATE TABLE cities").Error; err != nil {
		return err
	}
	return nil
}

func (r *CityRepository) SeedDummyCity() error {
	dummyCities := []CityModel{
		{ID: "1", Name: "Tokyo"},
		{ID: "2", Name: "Osaka"},
		{ID: "3", Name: "Kyoto"},
		{ID: "4", Name: "Nagoya"},
	}
	for _, city := range dummyCities {
		if err := r.db.WithContext(r.ctx).Create(&city).Error; err != nil {
			return err
		}
	}
	return nil
}
