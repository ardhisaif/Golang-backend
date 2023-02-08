package model

import (
	"time"

	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)

type User struct {
	UserID    uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`
	Name      string    `gorm:"unique;not null" json:"name" valid:"type(string)"`
	Email     string    `gorm:"unique;not null" json:"email" valid:"email,optional"`
	Password  string    `json:"password,omitempty" valid:"type(string)"`
	Role      string    `json:"role" valid:"type(string)"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

type Users []User
