package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string     `gorm:"type:varchar(100);not null" json:"name"`
	Username    string     `gorm:"type:varchar(100);not null;unique" json:"username"`
	UUID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();not null;unique" json:"uuid"`
	Pass        string     `gorm:"type:varchar(255);not null" json:"-"`
	Email       string     `gorm:"type:varchar(100);not null;unique" json:"email"`
	RoleID      int        `gorm:"not null" json:"role_id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Role        Role       `gorm:"foreignKey:RoleID" json:"role"`
	PhoneNumber string     `gorm:"type:varchar(15);not null;unique" json:"phone_number"`
}
