package main

import (
	"fmt"
	_ "gocms/routers"

	"gocms/models"

	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("gocms run...")
	models.Init()
	beego.Run()
}
