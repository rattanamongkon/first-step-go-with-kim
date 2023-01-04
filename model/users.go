package model

import "github.com/go-pg/pg/v10"

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	IsActive  bool   `json:"is_active"`
	Role      string `json:"role"`
	FactoryId int64  `json:"factory_id"`
	PlantId   int64  `json:"plant_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (s *User) CreateUser(db *pg.DB) error {
	_, err := db.Model(s).Insert()
	return err
}
