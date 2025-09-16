package models

import "time"

type Role struct {
	ID        int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string     `gorm:"type:varchar(100);not null;unique" json:"name"`
	Code      string     `gorm:"type:varchar(100);not null;unique" json:"code"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
