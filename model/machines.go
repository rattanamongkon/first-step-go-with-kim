package model

import "github.com/go-pg/pg/v10"

// type Machine struct {
// 	ID        int64  `json:"id"`
// 	CountFg   int64  `json:"count_fg"`
// 	CountNg   int64  `json:"count_ng"`
// 	TimeStamp string `json:"time_stamp"`
// 	FactoryId int64  `json:"factory_id"`
// 	PlantId   int64  `json:"plant_id"`
// 	StationId int64  `json:"station_id"`
// 	SkuId     int64  `json:"sku_id"`
// 	CreatedAt int64  `json:"created_at"`
// 	UpdatedAt int64  `json:"updated_at"`
// }
type Machine struct {
	ID        int64  `pg:",pk"`
	CountFg   int64  `pg:",notnull,default:0"`
	CountNg   int64  `pg:",notnull,default:0"`
	TimeStamp string `pg:""`
	FactoryId int64  `pg:",notnull"`
	PlantId   int64  `pg:",notnull"`
	StationId int64  `pg:""`
	SkuId     int64  `pg:""`
	CreatedAt int64
	UpdatedAt int64
}

type MachineInput struct {
	CountFg   int64  `json:"count_fg"`
	CountNg   int64  `json:"count_ng"`
	TimeStamp string `json:"time_stamp"`
	FactoryId int64  `json:"factory_id"`
	PlantId   int64  `json:"plant_id"`
	StationId int64  `json:"station_id"`
	SkuId     int64  `json:"sku_id"`
}

type MachineResponse struct {
	ID        int64  `json:"id,omitempty"`
	CountFg   int64  `json:"count_fg,omitempty"`
	CountNg   int64  `json:"count_ng,omitempty"`
	TimeStamp string `json:"time_stamp,omitempty"`
	FactoryId int64  `json:"factory_id,omitempty"`
	PlantId   int64  `json:"plant_id,omitempty"`
	StationId int64  `json:"station_id,omitempty"`
	SkuId     int64  `json:"sku_id,omitempty"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (s *Machine) CreateMachine(db *pg.DB) error {
	_, err := db.Model(s).Insert()
	return err
}
