package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// sql.DB 객체 생성
	db, err := sql.Open("mysql", "root:7777@tcp(127.0.0.1:3306)/netivenh")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 하나의 Row를 갖는 SQL 쿼리
	var name string
	err = db.QueryRow("SELECT bo_subject FROM g5_board WHERE bo_table = 'notice'").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}
