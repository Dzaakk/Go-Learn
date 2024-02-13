package learn_gorm

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
func TestQuerySingleObject(t *testing.T) {
	user := User{}
	err := db.First(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "1", user.ID)

	user = User{}
	err = db.Last(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "9", user.ID)
}
func TestQuerySingleObjectInlineCondition(t *testing.T) {
	user := User{}
	err := db.First(&user, "id = ?", "5").Error
	assert.Nil(t, err)
	assert.Equal(t, "5", user.ID)
	assert.Equal(t, "User 5", user.Name.FirstName)
}
func TestQueryAllObject(t *testing.T) {
	var users []User
	err := db.Find(&users, "id in ?", []string{"1,", "3", "4", "5"}).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))
}

func TestQueryCondition(t *testing.T) {
	var users []User
	err := db.
		Where("first_name like ?", "%User%").
		Where("password = ?", "rhs").
		Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 22, len(users))
}
func TestOperator(t *testing.T) {
	var users []User
	err := db.
		Where("first_name like ?", "%User%").
		Or("password = ?", "rhs").
		Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 23, len(users))
}
func TestNotOperator(t *testing.T) {
	var users []User
	err := db.
		Not("first_name like ?", "%User%").
		Where("password = ?", "rhs").
		Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}
func TestSelectFields(t *testing.T) {
	var users []User
	err := db.
		Select("id", "first_name").
		Find(&users).Error
	assert.Nil(t, err)
	for _, user := range users {
		assert.NotNil(t, user.ID)
		assert.NotEqual(t, "", user.Name.FirstName)
	}
	assert.Equal(t, 23, len(users))
}
func TestStructCondition(t *testing.T) {
	userCondition := User{
		Name: Name{
			FirstName: "User 5",
		},
		Password: "rhs",
	}

	var users []User
	err := db.Where(userCondition).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}
func TestMapCondition(t *testing.T) {
	mapCondition := map[string]interface{}{
		"middle_name": "",
		"last_name":   "",
	}
	var users []User
	err := db.Where(mapCondition).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 22, len(users))
}
func TestOrderLimitOffset(t *testing.T) {
	var users []User
	err := db.Order("id asc, first_name desc").
		Limit(5).
		Offset(5).
		Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 5, len(users))
}

type UserResponse struct {
	ID        string
	FirstName string
	LastName  string
}

func TestQueryNonModel(t *testing.T) {
	var users []UserResponse
	err := db.Model(&User{}).Select("id", "first_name", "last_name").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 23, len(users))
}
func TestUpdate(t *testing.T) {
	user := User{}
	err := db.Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)

	user.Name.FirstName = "User 1"
	user.Name.MiddleName = ""
	user.Name.LastName = ""

	err = db.Save(&user).Error
	assert.Nil(t, err)
}
func TestUpdateSelectedColumns(t *testing.T) {
	err := db.Model(&User{}).Where("id = ?", "1").Updates(map[string]interface{}{
		"middle_name": "",
		"last_name":   "buba",
	}).Error
	assert.Nil(t, err)

	err = db.Model(&User{}).Where("id = ?", "3").Update("password", "bukan rhs").Error
	assert.Nil(t, err)

	err = db.Where("id = ?", "4").Updates(User{
		Information: "HAHAHA",
	}).Error
	assert.Nil(t, err)
}
func TestAutoIncrement(t *testing.T) {
	for i := 0; i < 10; i++ {
		userLog := UserLog{
			UserId: "1",
			Action: "Test",
		}
		err := db.Create(&userLog).Error
		assert.Nil(t, err)

		assert.NotEqual(t, 0, userLog.ID)
	}
}
func TestSaveOrUpdate(t *testing.T) {
	userLog := UserLog{
		UserId: "1",
		Action: "Test",
	}

	err := db.Save(&userLog).Error
	assert.Nil(t, err)

	userLog.UserId = "2"
	err = db.Save(&userLog).Error
	assert.Nil(t, err)
}
func TestSaveOrUpdateNonAutoIncrement(t *testing.T) {
	user := User{
		ID: "99",
		Name: Name{
			FirstName: "User 99",
		},
	}

	err := db.Save(&user).Error
	assert.Nil(t, err)

	user.Name.FirstName = "User Updated 99"
	err = db.Save(&user).Error
	assert.Nil(t, err)
}
func TestConflict(t *testing.T) {
	user := User{
		ID: "88",
		Name: Name{
			FirstName: "User 88",
		},
	}
	err := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&user).Error
	assert.Nil(t, err)
}
func TestDelete(t *testing.T) {
	var user User
	err := db.Take(&user, "id = ?", "88").Error
	assert.Nil(t, err)

	err = db.Delete(&user).Error
	assert.Nil(t, err)

	err = db.Delete(&User{}, "id = ?", "99").Error
	assert.Nil(t, err)

	err = db.Where("id = ?", "10").Delete(&User{}).Error
	assert.Nil(t, err)
}
func TestSoftDelete(t *testing.T) {
	todo := Todo{
		UserId:      "3",
		Title:       "Todo 3",
		Description: "Desc 3",
	}

	err := db.Create(&todo).Error
	assert.Nil(t, err)

	err = db.Delete(&todo).Error
	assert.Nil(t, err)
	assert.NotNil(t, todo.DeletedAt)

	var todos []Todo
	err = db.Find(&todos).Error
	assert.Nil(t, err)
	assert.Equal(t, 0, len(todos))
}
func TestUnscoped(t *testing.T) {
	var todo Todo
	err := db.Unscoped().First(&todo, "id = ?", 2).Error
	assert.Nil(t, err)

	err = db.Unscoped().Delete(&todo).Error
	assert.Nil(t, err)

	var todos []Todo
	err = db.Find(&todos).Error
	assert.Nil(t, err)
}

func TestLock(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var user User
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&user, "id = ?", "1").Error
		if err != nil {
			return err
		}

		user.Name.FirstName = "sukuna"
		user.Name.LastName = "bro"
		err = tx.Save(&user).Error
		return err
	})
	assert.Nil(t, err)
}
func TestCreateWallet(t *testing.T) {
	wallet := Wallet{
		ID:      "1",
		UserId:  "1",
		Balance: 100000,
	}
	err := db.Create(&wallet).Error
	assert.Nil(t, err)
}
func TestRetrieveRelation(t *testing.T) {
	var user User
	err := db.Model(&User{}).Preload("Wallet").Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)

	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "1", user.Wallet.ID)
}
func TestRetrieveRelationJoin(t *testing.T) {
	var user User
	err := db.Model(&User{}).Joins("Wallet").Take(&user, "users.id = ?", "1").Error
	assert.Nil(t, err)

	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "1", user.Wallet.ID)
}
func TestAutoCreateUpdate(t *testing.T) {
	user := User{
		ID:       "34",
		Password: "rhs",
		Name: Name{
			FirstName: "User 34",
		},
		Wallet: Wallet{
			ID:      "20",
			UserId:  "34",
			Balance: 200000,
		},
	}

	err := db.Create(&user).Error
	assert.Nil(t, err)
}
func TestSkipAutoCreateUpdate(t *testing.T) {
	user := User{
		ID:       "35",
		Password: "rhs",
		Name: Name{
			FirstName: "User 35",
		},
		Wallet: Wallet{
			ID:      "21",
			UserId:  "35",
			Balance: 200000,
		},
	}

	err := db.Omit(clause.Associations).Create(&user).Error
	assert.Nil(t, err)
}
func TestUserAndAddresses(t *testing.T) {
	user := User{
		ID:       "91",
		Password: "rhs",
		Name: Name{
			FirstName: "user 91",
		},
		Wallet: Wallet{
			ID:      "91",
			UserId:  "91",
			Balance: 7000000,
		},
		Addresses: []Address{
			{
				UserId:  "91",
				Address: "jalan A",
			},
			{
				UserId:  "91",
				Address: "jalan B",
			},
		},
	}
	err := db.Create(&user).Error
	assert.Nil(t, err)
}
func TestPreloadJoinOneToMany(t *testing.T) {
	var users []User
	err := db.Model(&User{}).Preload("Addresses").Joins("Wallet").
		Find(&users).Error
	assert.Nil(t, err)
}
func TestTakePreloadJoinOneToMany(t *testing.T) {
	var user User
	err := db.Model(&User{}).Preload("Addresses").Joins("Wallet").
		Take(&user, "users.id = ?", "90").Error
	assert.Nil(t, err)
}
func TestBelongsToAddress(t *testing.T) {
	fmt.Println("Preload")
	var addresses []Address
	err := db.Model(&Address{}).Preload("User").Find(&addresses).Error
	assert.Nil(t, err)

	fmt.Println("Joins")
	addresses = []Address{}
	err = db.Model(&Address{}).Joins("User").Find(&addresses).Error
	assert.Nil(t, err)
}
func TestBelongsToWaller(t *testing.T) {
	fmt.Println("Preload")
	var wallets []Wallet
	err := db.Model(&Wallet{}).Preload("User").Find(&wallets).Error
	assert.Nil(t, err)

	fmt.Println("Joins")
	wallets = []Wallet{}
	err = db.Model(&Wallet{}).Joins("User").Find(&wallets).Error
	assert.Nil(t, err)
}

func TestCreateManyToMany(t *testing.T) {
	product := Product{
		ID:    "P002",
		Name:  "product1",
		Price: 10000000,
	}
	err := db.Create(&product).Error
	assert.Nil(t, err)

	err = db.Table("user_like_product").Create(map[string]interface{}{
		"user_id":    "1",
		"product_id": "P002",
	}).Error
	assert.Nil(t, err)

	err = db.Table("user_like_product").Create(map[string]interface{}{
		"user_id":    "3",
		"product_id": "P002",
	}).Error
	assert.Nil(t, err)
}
func TestPreloadManyToMany(t *testing.T) {
	var product Product
	err := db.Preload("LikedByUsers").Take(&product, "id = ? ", "P002").Error
	assert.Nil(t, err)
	assert.Equal(t, 2, len(product.LikedByUsers))
}
func TestPreloadManyToManyUser(t *testing.T) {
	var user User
	err := db.Preload("LikeProducts").Take(&user, "id = ? ", "1").Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(user.LikeProducts))
}

func TestAssociationFind(t *testing.T) {
	var product Product
	err := db.First(&product, "id = ?", "P002").Error
	assert.Nil(t, err)

	var users []User
	err = db.Model(&product).Where("users.first_name like ?", "User%").Association("LikedByUsers").Find(&users)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}
func TestAssociationAppend(t *testing.T) {
	var user User
	err := db.Take(&user, "id = ?", "90").Error
	assert.Nil(t, err)

	var product Product
	err = db.Take(&product, "id = ?", "P002").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Append(&user)
	assert.Nil(t, err)
}
func TestAssociationReplace(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var user User
		err := tx.Take(&user, "id = ?", "1").Error
		assert.Nil(t, err)

		wallet := Wallet{
			ID:      "01",
			UserId:  user.ID,
			Balance: 10000200000,
		}
		err = tx.Model(&user).Association("Wallet").Replace(&wallet)
		return err
	})
	assert.Nil(t, err)
}
func TestAssociationDelete(t *testing.T) {
	var user User
	err := db.Take(&user, "id = ?", "90").Error
	assert.Nil(t, err)

	var product Product
	err = db.Take(&product, "id = ?", "P002").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Delete(&user)
	assert.Nil(t, err)
}
func TestAssociationClear(t *testing.T) {
	var product Product
	err := db.Take(&product, "id = ?", "P002").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Clear()
	assert.Nil(t, err)
}
func TestPreloadingWithCondition(t *testing.T) {
	var user User
	err := db.Preload("Wallet", "balance > ?", 1000).
		Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)
	fmt.Println(user)
}
func TestNestedPreloading(t *testing.T) {
	var wallet Wallet
	err := db.Preload("User.Addresses").
		Take(&wallet, "id = ?", "1").Error
	assert.Nil(t, err)

	fmt.Println(wallet)
	fmt.Println(wallet.User)
	fmt.Println(wallet.User.Addresses)
}
func TestPreloadingAll(t *testing.T) {
	var user User
	err := db.Preload(clause.Associations).Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)
}
func TestJoinQuery(t *testing.T) {
	var users []User
	err := db.Joins("join wallets on wallets.user_id = users.id").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))

	users = []User{}
	err = db.Joins("Wallet").Find(&users).Error //left join
	assert.Nil(t, err)
	assert.Equal(t, 27, len(users))
}
func TestJoinWithCondition(t *testing.T) {
	var users []User
	err := db.Joins("join wallets on wallets.user_id = users.id AND wallets.balance > ?", 0).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))

	users = []User{}
	err = db.Joins("Wallet").Where("Wallet.balance > ?", 1000).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))
}

func TestCount(t *testing.T) {
	var count int64
	err := db.Model(&User{}).Joins("Wallet").Where("balance > ?", 20000).Count(&count).Error
	assert.Nil(t, err)
	assert.Equal(t, int64(4), count)
}

type AggregationResult struct {
	TotalBalance int64
	MinBalance   int64
	MaxBalance   int64
	AvgBalance   float64
}

func TestAggregation(t *testing.T) {
	var result AggregationResult
	err := db.Model(&Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance", "max(balance) as max_balance", "avg(balance) as avg_balance").Take(&result).Error
	assert.Nil(t, err)
	assert.Equal(t, int64(2000000), result.TotalBalance)
	assert.Equal(t, int64(200000), result.MinBalance)
	assert.Equal(t, int64(700000), result.MaxBalance)
	assert.Equal(t, float64(500000), result.AvgBalance)
}
func TestGroupByHaving(t *testing.T) {
	var results []AggregationResult
	err := db.Model(&Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance", "max(balance) as max_balance", "avg(balance) as avg_balance").
		Joins("User").Group("User.id").Having("sum(balance) > ?", 10000000).
		Find(&results).Error
	assert.Nil(t, err)
	assert.Equal(t, 0, len(results))
}
