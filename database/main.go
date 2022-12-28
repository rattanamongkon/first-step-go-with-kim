package database

import (
	"example/go-rest-api/database/migrate"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func RefreshSchema(db *pg.DB) {
	if err := DropSchema(db); err != nil {
		panic(err.Error())
	}

	if err := CreateSchema(db); err != nil {
		panic(err)
	}
}

func CreateSchema(db *pg.DB) error {
	for _, model := range migrate.GetModel() {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Varchar:       255,
			IfNotExists:   true,
			FKConstraints: true,
			Temp:          false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func DropSchema(db *pg.DB) error {
	for _, model := range migrate.GetModel() {
		err := db.Model(model).DropTable(&orm.DropTableOptions{
			IfExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
