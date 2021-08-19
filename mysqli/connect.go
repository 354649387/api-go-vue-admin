package mysqli

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *sqlx.DB

func Connect() *sqlx.DB {

	db, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_vue_admin")

	if err != nil {
		fmt.Println("mysql connect failed", err)
	}

	Db = db

	return Db
}

func GormConnect() *gorm.DB {

	//连接数据库
	dsn := "root:root@tcp(127.0.0.1:3306)/go_vue_admin?charset=utf8&parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("mysql connect failed", err)
	}

	return db
}
