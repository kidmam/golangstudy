package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// sql.DB 객체 생성
	db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 트랜잭션 시작
	tx, err := db.Begin()
	if err != nil {
		log.Panic(err)
	}
	defer tx.Rollback() //중간에 에러시 롤백하도록 defer 한다

	// INSERT 문 실행
	_, err = tx.Exec("INSERT INTO test1 VALUES (?, ?)", 15, "Jack")
	if err != nil {
		//에러메시지를 출력하고 panic() 호출.
		//panic()은 defer를 실행한다.
		log.Panic(err)
	}

	_, err = tx.Exec("INSERT INTO test2 VALUES (?, ?)", 15, "Data")
	if err != nil {
		log.Panic(err)
	}

	// 트랜잭션 커밋
	err = tx.Commit()
	if err != nil {
		log.Panic(err)
	}
}
