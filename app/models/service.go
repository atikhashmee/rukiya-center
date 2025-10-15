package models

import "github.com/goravel/framework/database/orm"

type Service struct {
	orm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	Status      string  `json:"status"`
}
