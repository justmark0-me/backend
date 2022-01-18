package main

import (
	"fmt"
	"log"
	_ "main/src/models"
	_ "main/src/routers"
	"os"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	postgresConfiguration := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DBNAME"))

	log.Println(fmt.Sprintf("Connecting to postgres. with user %s on host %s",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_HOST")))

	err := orm.RegisterDataBase("default", "postgres", postgresConfiguration)
	if err != nil {
		log.Fatal("Cannot connect to database.\nQUITTING.\nREASON:", err)
	}
	web.Run()
}
