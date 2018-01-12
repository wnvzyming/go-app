package controllers

import (
	"fmt"
	"gocms/libs"
	"gocms/models"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type LoginController struct {
	beego.Controller
	userId int
}

//ajax返回
func (self *LoginController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["message"] = msg
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

//获取用户IP地址
func (self *LoginController) getClientIp() string {
	s := strings.Split(self.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

func (self *LoginController) LoginIn() {

	if self.userId > 0 {
		self.Redirect(beego.URLFor("HomeController.Index"), 302)
	}
	beego.ReadFromRequest(&self.Controller)
	if self.Ctx.Request.Method == "POST" {

		username := strings.TrimSpace(self.GetString("username"))
		password := strings.TrimSpace(self.GetString("password"))
		rememberMe := strings.TrimSpace(self.GetString("rememberMe"))

		if username != "" && password != "" {
			user, err := models.AdminGetByName(username)
			fmt.Println(user)
			flash := beego.NewFlash()
			errorMsg := ""
			if err != nil {
				errorMsg = "帐号或密码错误"
			} else if user.Status == -1 {
				errorMsg = "该帐号已禁用"
			} else {
				user.LastIp = self.getClientIp()
				user.LastLogin = time.Now().Unix()
				user.Update()

				if rememberMe == "true" {
					authkey := libs.Md5([]byte(self.getClientIp() + "|" + user.Password + user.Salt))
					self.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 7*86400)
				}
				self.ajaxMsg(beego.URLFor("HomeController.Index"), 1)
				//self.Redirect(beego.URLFor("HomeController.Index"), 302)
			}

			self.ajaxMsg(beego.URLFor("LoginController.LoginIn"), 0)

			flash.Error(errorMsg)
			flash.Store(&self.Controller)

			//self.Redirect(beego.URLFor("LoginController.LoginIn"), 302)
		}
	}
	self.TplName = "login/login.html"

	//控制台log 输出
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)
	log.Debug("this is a debug message1111221")

}

//登出
func (self *LoginController) LoginOut() {
	self.Ctx.SetCookie("userId", "")
	self.Redirect(beego.URLFor("LoginController.LoginIn"), 302)
}
