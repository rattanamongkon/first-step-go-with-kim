package database

import (
	"example/go-rest-api/conf"
	"example/go-rest-api/connection"
	"example/go-rest-api/database/seeds"
	"log"
	"testing"

	"github.com/go-pg/pg/v10"
)

func TestRefreshSchema(t *testing.T) {
	conf.Init()
	type args struct {
		db *pg.DB
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: `CheckDatabase`,
			args: args{
				db: connection.GetConnectionDB(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RefreshSchema(tt.args.db)
			log.Println("Seed User Data 'admin' user")
			seed, _ := seeds.Seeds()
			log.Println(seed)
		})
	}
}
