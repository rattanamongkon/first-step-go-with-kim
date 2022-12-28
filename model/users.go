package model

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	IsActive  bool   `json:"is_active"`
	Role      string `json:"role"`
	FactoryId string `json:"factory_id"`
	PlantId   string `json:"plant_id"`
}
