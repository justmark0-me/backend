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
	postgresConnectionSettings := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DBNAME"))

	// Since this logs will be available on server and I were debugging deployment. I decided to log first symbol of password
	log.Println(fmt.Sprintf("Connecting to postgres. with user %s on host %s with port %s on database %s, password %s",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DBNAME"), os.Getenv("POSTGRES_PASSWORD")))

	err := orm.RegisterDataBase("default", "postgres", postgresConnectionSettings)
	if err != nil {
		log.Fatal("Cannot connect to database.\nQUITTING.\nREASON:", err)
	}
	web.Run()
}
