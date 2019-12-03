package bootstrap

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"strconv"
)

func LoadDatabase() func(*Bootstrapper) {
	return func(bootstrap *Bootstrapper) {
		// 加载database
		database := viper.GetStringMapString("db")
		db, err := gorm.Open(database["drive"],database["url"] )

		if err!=nil{
			panic(err)
		}

		// 连接池
		maxIdleConnection, _ := strconv.Atoi(database["maxIdleConnection"])
		maxOpenConnection, _ := strconv.Atoi(database["maxOpenConnection"])

		db.DB().SetMaxIdleConns(maxIdleConnection)
		db.DB().SetMaxOpenConns(maxOpenConnection)

		db.LogMode(true)
		bootstrap.DB = db
	}
}