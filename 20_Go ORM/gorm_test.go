package learn_gorm

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
	"testing"
)

func OpenConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Unable to load env")
	}
	host := os.Getenv("db_host")
	port := os.Getenv("db_port")
	name := os.Getenv("db_name")
	username := os.Getenv("db_username")
	password := os.Getenv("db_password")
	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + name + " port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	return db
}

var db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}

func TestExecuteSQL(t *testing.T) {
	err := db.Exec("insert into public.sample (id, name) values(?, ?)", "1", "Dadan").Error
	assert.Nil(t, err)

	err = db.Exec("insert into public.sample (id, name) values(?, ?)", "2", "Sasuke").Error
	assert.Nil(t, err)
}

type Sample struct {
	Id   string
	Name string
}

func TestRawSQL(t *testing.T) {
	var sample Sample

	err := db.Raw("select id, name from sample where id = ?", "1").Scan(&sample).Error
	assert.Nil(t, err)
	assert.Equal(t, "Dadan", sample.Name)

	var samples []Sample
	err = db.Raw("select id, name from sample").Scan(&samples).Error
	assert.Nil(t, err)
	assert.Equal(t, 2, len(samples))
}

func TestSqlRow(t *testing.T) {
	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Sample

	for rows.Next() {
		var id string
		var name string

		err := rows.Scan(&id, &name)
		assert.Nil(t, err)

		samples = append(samples, Sample{
			Id:   id,
			Name: name,
		})
	}
	assert.Equal(t, 2, len(samples))
}
func TestScanRow(t *testing.T) {
	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Sample

	for rows.Next() {
		err := db.ScanRows(rows, &samples)
		assert.Nil(t, err)
	}
	assert.Equal(t, 2, len(samples))
}

func TestCreateUser(t *testing.T) {
	user := User{
		ID:       "1",
		Password: "rhs",
		Name: Name{
			FirstName:  "Bambang",
			MiddleName: "Pamungkas",
			LastName:   "20",
		},
		Information: "ignore",
	}
	response := db.Create(&user)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(1), response.RowsAffected)
}

func TestBatchInsert(t *testing.T) {
	var users []User
	for i := 2; i < 10; i++ {
		users = append(users, User{
			ID:       strconv.Itoa(i),
			Password: "rhs",
			Name: Name{
				FirstName: "User " + strconv.Itoa(i),
			},
		})
	}
	response := db.Create(&users)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(8), response.RowsAffected)
}
func TestCreateInBatch(t *testing.T) {
	var users []User
	for i := 10; i < 20; i++ {
		users = append(users, User{
			ID:       strconv.Itoa(i),
			Password: "rhs",
			Name: Name{
				FirstName: "User " + strconv.Itoa(i),
			},
		})
	}
	response := db.CreateInBatches(&users, 10)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(10), response.RowsAffected)
}

func TestTransactionSuccess(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{ID: "20", Password: "rhs", Name: Name{FirstName: "User 20"}}).Error
		if err != nil {
			return err
		}
		err = tx.Create(&User{ID: "21", Password: "rhs", Name: Name{FirstName: "User 21"}}).Error
		if err != nil {
			return err
		}
		err = tx.Create(&User{ID: "22", Password: "rhs", Name: Name{FirstName: "User 22"}}).Error
		if err != nil {
			return err
		}
		return nil
	})

	assert.Nil(t, err)
}
func TestTransactionRollback(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{ID: "23", Password: "rhs", Name: Name{FirstName: "User 23"}}).Error
		if err != nil {
			return err
		}
		err = tx.Create(&User{ID: "11", Password: "rhs", Name: Name{FirstName: "User 11"}}).Error
		if err != nil {
			return err
		}
		return nil
	})

	assert.NotNil(t, err)
}
func TestTransactionManualSuccess(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{ID: "23", Password: "rhs", Name: Name{FirstName: "User 23"}}).Error
	assert.Nil(t, err)
	err = tx.Create(&User{ID: "24", Password: "rhs", Name: Name{FirstName: "User 24"}}).Error
	assert.Nil(t, err)

	if err == nil {
		tx.Commit()
	}
}
func TestTransactionManualError(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{ID: "25", Password: "rhs", Name: Name{FirstName: "User 25"}}).Error
	assert.Nil(t, err)
	err = tx.Create(&User{ID: "16", Password: "rhs", Name: Name{FirstName: "User 16"}}).Error
	assert.Nil(t, err)

	if err == nil {
		tx.Commit()
	}
}
