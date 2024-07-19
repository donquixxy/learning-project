package entity

import (
	"time"
)

type User struct {
	ID         int        `gorm:"primaryKey;column:id;autoIncrement:true" json:"id"`
	Name       string     `gorm:"column:name" json:"name,omitempty"`
	LoginToken string     `gorm:"column:login_token" json:"login_token,omitempty"`
	Password   string     `gorm:"column:password" json:"password,omitempty"`
	Phone      string     `gorm:"column:phone" json:"phone,omitempty"`
	Email      string     `gorm:"column:email" json:"email,omitempty"`
	CreatedAt  time.Time  `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt  *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	DeletedAt  *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
}
