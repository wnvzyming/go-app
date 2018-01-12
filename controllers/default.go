package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "heheh"
	c.Data["Email"] = "astaxie@gmail.com22222211111"

	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)
	log.Debug("this is a debug message1111221")

	c.TplName = "index.tpl"
}
