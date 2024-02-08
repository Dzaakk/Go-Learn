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
	Wallet      Wallet    `gorm:"foreignKey:user_id;references:id"`
	Addresses   []Address `gorm:"foreignKey:user_id;references:id"`
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
	gorm.Model
	UserId      string `gorm:"column:user_id"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
}

func (t *Todo) TableName() string {
	return "todos"
}

type Wallet struct {
	ID        string    `gorm:"primary_key;column:id;<-:create;autoIncrement"` //create only
	UserId    string    `gorm:"column:user_id"`
	Balance   int64     `gorm:"column:balance"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	User      *User     `gorm:"foreignKey:user_id;references:id"`
}

func (w *Wallet) TableName() string {
	return "wallets"
}

type Address struct {
	ID        string    `gorm:"primary_key;column:id;autoIncrement"`
	UserId    string    `gorm:"column:user_id"`
	Address   string    `gorm:"column:address"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	User      User      `gorm:"foreignKey:user_id;references:id"`
}

func (a *Address) TableName() string {
	return "addresses"
}
