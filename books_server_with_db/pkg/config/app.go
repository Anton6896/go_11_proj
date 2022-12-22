package config

import (
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	log "github.com/sirupsen/logrus"
)

var (
	db *gorm.DB
)

func Connect() {
	log.Info("connecting to DB [gorm]")

	loc, _ := time.LoadLocation("Asia/Jerusalem")

	c := mysql.Config{
		User:      "anton",
		Passwd:    "132123",
		Net:       "tcp",
		Addr:      "host.docker.internal:3306",
		DBName:    "simplerest",
		ParseTime: true,
		Loc:       loc,
	}

	log.Info(c.FormatDSN())  // for testing leaning propose

	// d, err := gorm.Open("mysql", "anton:132123@/simplerest?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open("mysql", c.FormatDSN())
	
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDb() *gorm.DB {
	return db
}
