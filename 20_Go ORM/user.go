package learn_gorm

import "time"

type User struct {
	ID          string    `gorm:"primary_key;column:id;<-:create"` //create only
	password    string    `gorm:"column:password"`
	Name        string    `gorm:"column:name;<-"` //create and update
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt   time.Time `gorm:"column:name;autoCreateTime;autoUpdateTime"`
	Information string    `gorm:"-"` //no field on db so it will ignore this
}

func (u User) TableName() string {
	return "users"
}
