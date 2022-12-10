package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("using sqlite")
}

/*
	docker run --rm --name ant-mysql \
	-e MYSQL_USER=anton \
	-e MYSQL_PASSWORD=132123 \
	-e MYSQL_DATABASE=simplerest \
	-e MYSQL_ROOT_PASSWORD=123123 \
	-p 3306:3306 \
	-d mysql:latest

*/
