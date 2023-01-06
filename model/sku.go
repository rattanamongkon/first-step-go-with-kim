package model

import "github.com/go-pg/pg/v10"

type Sku struct {
	ID        int64  `pg:",pk"`
	Name      string `pg:",notnull"`
	Code      string `pg:",notnull,unique"`
	CreatedAt int64
	UpdatedAt int64
}
type SkuResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (s *Sku) CreateSku(db *pg.DB) error {
	_, err := db.Model(s).Insert()
	return err
}
