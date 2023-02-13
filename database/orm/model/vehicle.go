package model

import (
	"time"
)

type Vehicle struct {
	VehicleID string    `gorm:"primaryKey;default:uuid_generate_v4()" json:"id,omitempty"`
	Type      string    `json:"type"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	Location  string    `json:"location"`
	Capacity  int       `json:"capacity"`
	Price     int64     `json:"price"`
	Point     int64     `json:"point"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Vehicles []Vehicle
