package models

import (
	"strconv"

	"github.com/astaxie/beego/orm"
)

type Admin struct {
	Id         int
	LoginName  string
	RealName   string
	Password   string
	RoleIds    string
	Phone      string
	Email      string
	Salt       string
	LastLogin  int64
	LastIp     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (a *Admin) TableName() string {
	return TableName("admin")
}

func AdminAdd(a *Admin) (int64, error) {

	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	return o.Insert(a)
}

func AdminGetByName(loginName string) (*Admin, error) {
	a := new(Admin)
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	err := o.QueryTable(TableName("admin")).Filter("login_name", loginName).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func AdminGetList(page, pageSize int, filters ...interface{}) ([]*Admin, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Admin, 0)
	query := orm.NewOrm()
	query.Using("default")
	realName := ""
	sql := ""
	var total int64
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			if filters[k].(string) == "realName" {
				realName = filters[k+1].(string)
			}
		}
	}
	if realName == "" {
		sql = "SELECT * FROM cms_admin ORDER BY id DESC LIMIT ?,?"
		total, _ = query.Raw(sql, strconv.Itoa(offset), strconv.Itoa(pageSize)).QueryRows(&list)
	} else {
		sql = "SELECT * FROM cms_admin WHERE real_name like ?  ORDER BY id DESC LIMIT ?,?"
		total, _ = query.Raw(sql, "%"+realName+"%", strconv.Itoa(offset), strconv.Itoa(pageSize)).QueryRows(&list)
	}
	return list, total
}

func AdminGetById(id int) (*Admin, error) {
	r := new(Admin)
	o := orm.NewOrm()
	o.Using("default")
	err := o.QueryTable(TableName("admin")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (a *Admin) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}

// func RoleAuthDelete(id int) (int64, error) {
// 	query := orm.NewOrm().QueryTable(TableName("role_auth"))
// 	return query.Filter("role_id", id).Delete()
// }

// func RoleAuthMultiAdd(ras []*RoleAuth) (n int, err error) {
// 	query := orm.NewOrm().QueryTable(TableName("role_auth"))
// 	i, _ := query.PrepareInsert()
// 	for _, ra := range ras {
// 		_, err := i.Insert(ra)
// 		if err == nil {
// 			n = n + 1
// 		}
// 	}
// 	i.Close() // 别忘记关闭 statement
// 	return n, err
// }
