package mysqli

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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
