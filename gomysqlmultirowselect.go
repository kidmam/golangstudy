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

	// 복수 Row를 갖는 SQL 쿼리
	var id int
	var name string
	rows, err := db.Query("SELECT bo_subject FROM g5_board WHERE bo_table >= ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() //반드시 닫는다 (지연하여 닫기)

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
}
