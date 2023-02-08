package model

import (
	"time"

	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)

type Reservation struct {
	ReservationID uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id,omitempty"`
	UserID        string    `gorm:"foreignKey:UserID;references:UserID" json:"user_id"`
	VehicleID     string    `gorm:"foreignKey:VehicleID;references:VehicleID" json:"vehicle_id"`
	Quantity      int       `json:"quantity"`
	StartDate     time.Time `json:"start_date"`
	ReturnDate    time.Time `json:"return_date"`
	TotalPayment  int       `json:"total_payment"`
	IsBooked      bool      `json:"is_booked"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	User    User    `json:"user"`
	Vehicle Vehicle `json:"vehicle"`
}

type Reservations []Reservation
