package learn_gorm

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
