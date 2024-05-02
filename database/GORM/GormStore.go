package GORM

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormStore struct {
	Db *gorm.DB
}

func NewGormStore(dsn string) (*GormStore, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("error connecting the database :%w", err)
	}
	return &GormStore{
		Db: db,
	}, nil
}
