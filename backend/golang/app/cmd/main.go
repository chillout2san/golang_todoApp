package main

import (
	"flag"
	"crud/controller"
	"crud/model"
	"net/http"
	"fmt"
)

var tm = model.CreateTodoModel()
var tc = controller.CreateTodoController(tm)
var ro = controller.CreateRouter(tc)

func migrate() {
	sql := `INSERT INTO todos(id, name, status) VALUES('00000000000000000000000000','買い物', '作業中'),('00000000000000000000000001','洗濯', '作業中'),('00000000000000000000000002','皿洗い', '完了');`

	_, err := model.Db.Exec(sql)

	if err != nil {
		// TODO: ログの出力方法を検討する
		fmt.Println(err)
		return
	}

	fmt.Println("Migration is success!")
}

func main() {
	f := flag.String("option", "", "migrate database or not")
	flag.Parse()
	if *f == "migrate" {
		migrate()
	}
	http.HandleFunc("/fetch-todos", ro.FetchTodos)
	http.ListenAndServe(":8080", nil)
}
