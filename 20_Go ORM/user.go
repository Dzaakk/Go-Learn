package learn_gorm

import "time"

type User struct {
	ID        string    `gorm:"primary_key;column:id"`
	password  string    `gorm:"column:password"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:name;autoCreateTime;autoUpdateTime"`
}

func (u User) TableName() string {
	return "users"
}
