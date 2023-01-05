package model

import "github.com/go-pg/pg/v10"

type Station struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	IsActive  bool   `json:"is_active"`
	PlantId   int64  `json:"plant_id"`
	FactoryId int64  `json:"factory_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (s *Station) CreateStation(db *pg.DB) error {
	_, err := db.Model(s).Insert()
	return err
}
