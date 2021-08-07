package admin

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Admin struct {
	Id int `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}



var Db *sqlx.DB


func init(){

	database,err := sqlx.Open("mysql","root:root@tcp(127.0.0.1:3306)/go_vue_admin")

	if err != nil {
		fmt.Println("mysql connect failed:",err)
		return
	}

	Db = database

}

/*
获取管理员列表
返回一个管理员结构体的切片
 */
func List() []Admin {

	//通过切片存储
	//admins := make([]Admin,0)
	var admins []Admin

	rows,_ := Db.Query("select * from admin")


	//fmt.Println(rows)

	//遍历
	var admin Admin

	for rows.Next(){

		rows.Scan(&admin.Id,&admin.Username,&admin.Password)

		admins = append(admins,admin)
	}

	return admins

}


func Add() bool {

	username := "mila"
	password := "123456"

	res,err := Db.Exec("insert into admin(username,password) values(?,?)",username,password)

	if err != nil {
		fmt.Println("insert failed:",err)
	}

	num,_ := res.RowsAffected()

	if num > 0 {
		return true
	}else {
		return false
	}

}
