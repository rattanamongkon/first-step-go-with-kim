package seeds

import (
	"example/go-rest-api/connection"
	"example/go-rest-api/model"
	"example/go-rest-api/utils"
	"log"
	"time"
)

func Seeds() (string, error) {
	db := connection.GetConnectionDB()
	_ = db
	var status = "fail"

	seeduser := model.UserSingUpInput{
		Name:      "admin",
		Username:  "admin",
		Password:  "admin",
		Email:     "admin@localhost.test",
		Role:      "admin",
		IsActive:  false,
		FactoryId: 1,
		PlantId:   1,
	}
	hashedPassword, _ := utils.HashPassword(seeduser.Password)

	user := model.User{
		Name:      seeduser.Name,
		Username:  seeduser.Username,
		Password:  hashedPassword,
		Email:     seeduser.Email,
		Role:      seeduser.Role,
		IsActive:  &seeduser.IsActive,
		FactoryId: seeduser.FactoryId,
		PlantId:   seeduser.PlantId,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	// log.Println("in Seeds!!!!")
	// log.Println(&user)

	if _, err := db.Model(&user).Insert(); err != nil {
		log.Println(err.Error())
	}
	status = "success"
	return status, nil
}
