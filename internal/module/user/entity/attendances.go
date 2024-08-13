package entity

import "time"

type Attendances struct {
	ID           int        `gorm:"primary_key;column:id;auto_increment" json:"id"`
	UserID       int        `gorm:"column:user_id" json:"user_id"`
	ClockInTime  time.Time  `gorm:"column:clock_in_time" json:"clock_in_time"`
	ClockOutTime *time.Time `gorm:"column:clock_out_time" json:"clock_out_time"`
	CreatedAt    *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at" json:"updated_at"`
	User         *User      `gorm:"foreignKey:UserID" json:"user"`
}
