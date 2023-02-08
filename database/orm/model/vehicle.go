package model

import (
	"time"

	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)

type Vehicle struct {
	VehicleID    uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id,omitempty"`
	Type      string    `json:"type"`
	Name      string    `json:"name"`
	Image 		string `json:"image"`
	Location  string    `json:"location"`
	Capacity  int       `json:"capacity"`
	Price     int64     `json:"price"`
	Point     int64     `json:"point"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Vehicles []Vehicle
