package connection

import "github.com/go-pg/pg/v10"

var db *pg.DB

func GetConnectionDB() *pg.DB {
	if db == nil {
		db = pg.Connect(&pg.Options{
			Addr:     ":5432",
			User:     "postgres",
			Password: "",
			Database: "mydatabase",
		})
	}
	return db
}
