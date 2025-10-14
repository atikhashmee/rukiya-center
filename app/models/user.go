package models

import (
	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	Name        string `json:"name"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Type        string `json:"type"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	DateOfBirth string `json:"date_of_birth"`
}
