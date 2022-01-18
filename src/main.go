package main

import (
	"fmt"
	"log"
	_ "main/src/models"
	_ "main/src/routers"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	postgresConnectionSettings := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DBNAME"))

	log.Println(fmt.Sprintf("Connecting to postgres. With user %s on host %s with port %s on database %s",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DBNAME")))

	err := orm.RegisterDataBase("default", "postgres", postgresConnectionSettings)
	if err != nil {
		log.Fatal("Cannot connect to database.\nQUITTING.\nREASON:", err)
	}
	// CORS allow all
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	web.Run()
}
