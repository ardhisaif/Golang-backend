package model

import "time"

type Reservation struct {
	ReservationID uint `gorm:"primaryKey" json:"id,omitempty"`
	UserID        int		`json:"user_id"`
	VehicleID     int       `json:"vehicle_id"`
	Quantity      int       `json:"quantity"`
	StartDate     time.Time `json:"start_date"`
	ReturnDate    time.Time `json:"return_date"`
	TotalPayment  int       `json:"total_payment"`
	IsBooked      bool      `json:"is_booked"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	User    User    `gorm:"foreignKey:UserID" json:"-"`
	Vehicle Vehicle `gorm:"foreignKey:VehicleID" json:"-"`
}

type Reservations []Reservation
