package model

import (
	"os"
	"fmt"
	"database/sql"
)

var Db *sql.DB

func init() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))
	Db, err = sql.Open("mysql", dsn)

	if err != nil {
		// TODO: ログの出力方法を検討する
		fmt.Println(err)
		return
	}

	err = Db.Ping()

	if err != nil {
		// TODO: ログの出力方法を検討する
		fmt.Println(err)
		return
	}

	fmt.Println("Connection has been established!")
}