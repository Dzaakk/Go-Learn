package learn_gorm

import "time"

type User struct {
	ID          string    `gorm:"primary_key;column:id;<-:create"` //create only
	Password    string    `gorm:"column:password"`
	Name        Name      `gorm:"embedded"` //create and update
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Information string    `gorm:"-"` //no field on db so it will ignore this
}

func (u User) TableName() string {
	return "users"
}

type Name struct {
	FirstName  string `gorm:"column:first_name"`
	MiddleName string `gorm:"column:middle_name"`
	LastName   string `gorm:"column:last_name"`
}
