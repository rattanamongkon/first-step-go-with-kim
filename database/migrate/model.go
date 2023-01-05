package migrate

import "example/go-rest-api/model"

func GetModel() []interface{} {
	return []interface{}{
		(*model.User)(nil),
		(*model.Station)(nil),
		(*model.Sku)(nil),
		(*model.Plant)(nil),
		(*model.Factory)(nil),
		(*model.Machine)(nil),
	}
}
