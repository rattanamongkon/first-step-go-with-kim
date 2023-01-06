package model

import "github.com/go-pg/pg/v10"

type Factory struct {
	ID        int64  `pg:",pk"`
	Name      string `pg:",notnull"`
	Code      string `pg:",notnull,unique"`
	IsActive  *bool  `pg:",notnull,default:false"`
	CreatedAt int64
	UpdatedAt int64
}

func (s *Factory) CreateFactory(db *pg.DB) error {

	_, err := db.Model(s).Insert()
	return err
}
