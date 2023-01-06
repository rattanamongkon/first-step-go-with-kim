package model

import (
	"github.com/go-pg/pg/v10"
)

type Plant struct {
	ID        int64  `pg:",pk"`
	Name      string `pg:",notnull"`
	Code      string `pg:",notnull,unique"`
	IsActive  *bool  `pg:",notnull,default:false"`
	FactoryId int64  `pg:",notnull"`
	CreatedAt int64
	UpdatedAt int64
}

type PlantResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	IsActive  bool   `json:"is_active"`
	FactoryId int64  `json:"factory_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (s *Plant) CreatePlant(db *pg.DB) error {
	_, err := db.Model(s).Insert()
	return err
}
