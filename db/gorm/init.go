package gorm


import (
"fmt"
	"github.com/sangnguyen09/go_template/config"
	"log"
"sync"
"time"

"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	once sync.Once
	db_gorm  *gorm.DB
)

func init() {
	// GetDB()
}

func new() *gorm.DB {
	gorm.NowFunc = func() time.Time {
		return time.Now().UTC().Truncate(1000 * time.Nanosecond)
	}
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		config.Config.PostgresDB.Host, config.Config.PostgresDB.User, config.Config.PostgresDB.DatabaseName, config.Config.PostgresDB.Password))
	if err != nil {

		log.Printf("Connect not success to postgres database at host:%s with user:%s and db:%s",
			config.Config.PostgresDB.Host, config.Config.PostgresDB.User, config.Config.PostgresDB.DatabaseName)
	}
	// db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	// db.DB().SetConnMaxLifetime(time.Nanosecond)
	// db.DB().SetConnMaxLifetime(3 * time.Second)
	db.DB().Ping()
	// db.DB().Exec("SET timezone TO 'Asia/Ho_Chi_Minh';")
	if config.Config.Db.Debug {
		db.LogMode(true)
	}
	return db
}

func GetDB() *gorm.DB {
	once.Do(func() {// gọi 1 lần
		db_gorm = new()
	})
	return db_gorm
}
