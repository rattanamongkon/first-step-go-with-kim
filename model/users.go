package model

import (
	"github.com/go-pg/pg/v10"
)

//	type User struct {
//		ID        int64  `json:"id" gorm:"primarykey"`
//		Name      string `json:"name" binding:"required"`
//		Email     string `json:"email" gorm:"unique, uniqueIndex" binding:"required"`
//		Username  string `json:"username" gorm:"unique, uniqueIndex" binding:"required" `
//		Password  string `json:"password" binding:"required"`
//		IsActive  bool   `json:"is_active"`
//		Role      string `json:"role" binding:"required"`
//		FactoryId int64  `json:"factory_id"`
//		PlantId   int64  `json:"plant_id"`
//		CreatedAt int64  `json:"created_at"`
//		UpdatedAt int64  `json:"updated_at"`
//	}
type User struct {
	ID        int64  `pg:",pk"`
	Name      string `pg:",notnull"`
	Username  string `pg:",unique"`
	Password  string `pg:",notnull"`
	Email     string `pg:""`
	Role      string `pg:",notnull"`
	IsActive  *bool  `pg:",notnull,default:false"`
	FactoryId int64  `pg:",notnull"`
	PlantId   int64  `pg:",notnull"`
	CreatedAt int64
	UpdatedAt int64
}

type UserSingUpInput struct {
	Name      string `json:"name" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required,min=8"`
	Email     string `json:"email"`
	Role      string `json:"role" binding:"required"`
	IsActive  bool   `json:"is_active"`
	FactoryId int64  `json:"factory_id" binding:"required"`
	PlantId   int64  `json:"plant_id" binding:"required"`
}
type UserSignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Role      string `json:"role,omitempty"`
	FactoryId int64  `json:"factory_id,omitempty"`
	PlantId   int64  `json:"plant_id,omitempty"`
	IsActive  bool   `json:"is_active"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (s *User) CreateUser(db *pg.DB) error {
	_, err := db.Model(s).Insert()
	return err
}
