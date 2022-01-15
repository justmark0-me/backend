package main

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	//_ "github.com/lib/pq"
	"log"
	_ "main/src/models"
	_ "main/src/routers"
	"os"
)

func main() {
	err := orm.RegisterDataBase("default", "postgres", os.Getenv("POSTGRES_CONNECTION_CONFIGURATION"))
	if err != nil {
		log.Fatal("Cannot connect to database.\n QUITTING")
	}
	//o := orm.NewOrm()
	//a := models.IpRequest{}
	//id, err := o.Insert(&a)
	//fmt.Println(id)
	web.Run()
}
