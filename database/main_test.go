package database

import (
	"example/go-rest-api/conf"
	"example/go-rest-api/connection"
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
			name: `sdf`,
			args: args{
				db: connection.GetConnectionDB(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RefreshSchema(tt.args.db)
		})
	}
}
