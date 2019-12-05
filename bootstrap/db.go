package bootstrap

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/muxiaopie/go-mall/pkg/logger"
	"github.com/spf13/viper"
	"strconv"
	"time"
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


		db.SetLogger(&ormLogger{})
		db.LogMode(true)

		db.DB().SetMaxIdleConns(maxIdleConnection)
		db.DB().SetMaxOpenConns(maxOpenConnection)

		db.LogMode(true)
		bootstrap.DB = db
	}
}

type ormLogger struct {}

func (*ormLogger) Print(v ...interface{}) {
	log := logger.Logger
	if v[0] == "sql" {
		if t,ok := v[2].(time.Duration);ok {
			if t.Seconds() > 1.00 {
				log.Warn(v[3])
			}else{
				log.Info(v[3])
			}
		}else{
			log.Error(v[3])
			fmt.Println(fmt.Sprintf("%T", v[2]))
		}
	}
	if v[0] == "log" {
		log.Info(v[2])
	}
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}