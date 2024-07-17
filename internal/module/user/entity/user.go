package entity

import "time"

type User struct {
	ID         int        `gorm:"primaryKey;column:id;autoIncrement:true"`
	Name       string     `gorm:"column:name"`
	LoginToken string     `gorm:"column:login_token"`
	Password   string     `gorm:"column:password"`
	Phone      string     `gorm:"column:phone"`
	Email      string     `gorm:"column:email"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at"`
	DeletedAt  *time.Time `gorm:"column:deleted_at"`
}
