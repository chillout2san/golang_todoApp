package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "crud/model"
)

func main() {
	fmt.Println("main関数が実行された")
}
