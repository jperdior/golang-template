package main

import (
	"golang-template/cmd/api/bootstrap"
	_ "golang-template/docs"
	"golang-template/internal/platform/database/mysql"
	"log"
)

func main() {

	mysql.ConnectDB()
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
