package models

import "time"

// BaseModel is a base model for other models
type BaseModel struct {
	ID        uint       `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// Supply ...
// Supply model
type Supply struct {
	BaseModel
	Location string
	Quantity float32
	Value    float32
	Km       float32
	Date     time.Time
}

// User ...
// User model
type User struct {
	BaseModel
	Name     string
	Email    string
	Password []byte
}
