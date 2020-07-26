package frame

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"go.elastic.co/apm/module/apmgorm"
	"sync"
	"time"
)

var db *gorm.DB
var dbOnce sync.Once

const (
	MAX_IDLE_CONNS int = 10
	MAX_OPEN_CONNS int = 20
	MAX_LIFE_TIME  int = 60
)

func Instance() (*gorm.DB, error) {
	var err error
	dbOnce.Do(func() {
		user := viper.GetString("database.username")
		password := viper.GetString("database.password")
		host := viper.GetString("database.host")
		port := viper.GetInt("database.port")
		dbName := viper.GetString("database.dbname")
		args := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbName)
		db, err = apmgorm.Open("mysql", args)
		if err == nil {
			maxIdle := viper.GetInt("database.maxIdle")
			maxOpen := viper.GetInt("database.maxOpen")
			maxLifetime := viper.GetInt("database.maxLifetime")
			if maxIdle < 1 {
				maxIdle = MAX_IDLE_CONNS
			}
			if maxOpen < 1 {
				maxOpen = MAX_OPEN_CONNS
			}
			if maxLifetime < 1 {
				maxLifetime = MAX_LIFE_TIME
			}
			db.DB().SetMaxIdleConns(maxIdle)
			db.DB().SetMaxOpenConns(maxOpen)
			db.DB().SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second)
			if viper.GetBool("database.debug") {
				db.LogMode(true)
			}
		}
	})

	return db, err
}
