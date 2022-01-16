package main

import (
	"log"
	_ "main/src/models"
	_ "main/src/routers"
	"os"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	err := orm.RegisterDataBase("default", "postgres", os.Getenv("POSTGRES_CONNECTION_CONFIGURATION"))
	if err != nil {
		log.Fatal("Cannot connect to database.\nQUITTING.\nREASON:", err)
	}
	web.Run()
}
