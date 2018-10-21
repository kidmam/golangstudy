package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// sql.DB 객체 생성
	db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Prepared Statement 생성
	stmt, err := db.Prepare("UPDATE test1 SET name=? WHERE id=?")
	checkError(err)
	defer stmt.Close()

	// Prepared Statement 실행
	_, err = stmt.Exec("Tom", 1) //Placeholder 파라미터 순서대로 전달
	checkError(err)
	_, err = stmt.Exec("Jack", 2)
	checkError(err)
	_, err = stmt.Exec("Shawn", 3)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
