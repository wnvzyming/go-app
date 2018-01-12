package controllers

import (
	"gocms/libs"
	"gocms/models"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
	userId    int
	loginName string
	userName  string
	user      *models.Admin
	pageSize  int
}

//获取用户IP地址
func (self *AdminController) getClientIp() string {
	s := strings.Split(self.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

//扩展函数  判断是否登录
func (self *AdminController) Prepare() {
	arr := strings.Split(self.Ctx.GetCookie("auth"), "|")
	self.userId = 0
	if len(arr) == 2 {
		idstr, password := arr[0], arr[1]
		userId, _ := strconv.Atoi(idstr)
		if userId > 0 {
			user, err := models.AdminGetById(userId)
			if err == nil && password == libs.Md5([]byte(self.getClientIp()+"|"+user.Password+user.Salt)) {
				self.userId = user.Id
				self.loginName = user.LoginName
				self.userName = user.RealName
				self.user = user
			}
		}
	}

	//	if self.userId == 0 && (self.controllerName != "login" && self.actionName != "loginin") {
	//		self.Redirect(beego.URLFor("LoginController.LoginIn"), 302)
	//	}
}

func (self *AdminController) Index() {

	self.TplName = "admin/index.html"
}

func (self *AdminController) GetList() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	realName := strings.TrimSpace(self.GetString("realName"))

	StatusText := make(map[int]string)
	StatusText[0] = "<font color='red'>禁用</font>"
	StatusText[1] = "正常"

	self.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	//
	if realName != "" {
		filters = append(filters, "realName", realName)
	}
	result, count := models.AdminGetList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["login_name"] = v.LoginName
		row["real_name"] = v.RealName
		row["phone"] = v.Phone
		row["email"] = v.Email
		row["role_ids"] = v.RoleIds
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		row["status"] = v.Status
		row["status_text"] = StatusText[v.Status]
		list[k] = row
	}

	out := make(map[string]interface{})
	out["code"] = 1
	out["msg"] = "成功"
	out["count"] = count
	out["data"] = list
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()

}
