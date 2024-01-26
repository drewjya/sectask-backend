package orm

import (
	"time"

	"github.com/jinzhu/gorm"
)

type (
	BaseModel struct {
		ID        uint64     `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key,column:id"`
		CreatedAt time.Time  `json:"created_at" gorm:"column:created_at" sql:"DEFAULT:current_timestamp"`
		UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at" sql:"DEFAULT:current_timestamp"`
	}
	PaginationResponse struct {
		Total       int         `json:"total"`
		PerPage     int         `json:"per_page"`
		CurrentPage int         `json:"current_page"`
		LastPage    int         `json:"last_page"`
		From        int         `json:"from"`
		To          int         `json:"to"`
		Data        interface{} `json:"data"`
	}

	DBFunc func(tx *gorm.DB) error // func type which accept *gorm.DB and return error
)

var (
	_page  = 1
	_limit = 25
	_order = "id"
	// _roleAdmin = "admin"
	RoleAdmin  = "admin"
	_algorithm = "SHA256"
)
