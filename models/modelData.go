package models

import (
	"github.com/astaxie/beego/orm"
)

type MysqlModel struct {
	Id int
	Name string
}

func SelectDB() MysqlModel{

	o := orm.NewOrm()
	//o.Using("default")
	var so MysqlModel

	o.Raw("select id,name from student where id = ?",2).QueryRow(&so)
	return so

}