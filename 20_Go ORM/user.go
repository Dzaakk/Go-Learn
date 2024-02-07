package learn_gorm

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          string    `gorm:"primary_key;column:id;<-:create"` //create only
	Password    string    `gorm:"column:password"`
	Name        Name      `gorm:"embedded"` //create and update
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Information string    `gorm:"-"` //no field on db so it will ignore this
}

func (u *User) TableName() string {
	return "users"
}

type Name struct {
	FirstName  string `gorm:"column:first_name"`
	MiddleName string `gorm:"column:middle_name"`
	LastName   string `gorm:"column:last_name"`
}

type UserLog struct {
	ID        int    `gorm:"primary_key;column:id;<-:create;autoIncrement"` //create only
	UserId    string `gorm:"column:user_id"`
	Action    string `gorm:"column:action"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (l *UserLog) TableName() string {
	return "user_logs"
}

type Todo struct {
	ID          int64          `gorm:"primary_key;column:id;autoIncrement"` //create only
	UserId      string         `gorm:"column:user_id"`
	Title       string         `gorm:"column:title"`
	Description string         `gorm:"column:description"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (t *Todo) TableName() string {
	return "todos"
}
