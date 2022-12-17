package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	log "github.com/sirupsen/logrus"
)

var (
	db *gorm.DB
)

func Connect() {
	log.Info("connecting to DB [gorm]")

	d, err := gorm.Open("mysql", "anton:132123@/simplerest?charset=utf8&parseTime=True&loc=Local")
	// d, err := gorm.Open("mysql", "anton:132123@host.docker.internal:3306/simplerest?charset=utf8&parseTime=True&loc=local")
	
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDb() *gorm.DB {
	return db
}
