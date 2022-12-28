package model

type Machine struct {
	ID        int64  `json:"id"`
	CountFg   int64  `json:"count_fg"`
	CountNg   int64  `json:"count_ng"`
	TimeStamp string `json:"time_stamp"`
	FactoryId int64  `json:"factory_id"`
	PlantId   int64  `json:"plant_id"`
	StationId int64  `json:"station_id"`
	SkuId     int64  `json:"sku_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
