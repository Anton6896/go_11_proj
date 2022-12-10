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

	/* 
		docker run --rm --name ant-mysql \
		-e MYSQL_USER=anton \
		-e MYSQL_PASSWORD=132123 \
		-e MYSQL_DATABASE=simplerest \
		-e MYSQL_ROOT_PASSWORD=123123 \
		-p 3306:3306 \
		-d mysql:latest
	
	*/
	d, err := gorm.Open("mysql", "anton:132123@/simplerest?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDb() *gorm.DB {
	return db
}
