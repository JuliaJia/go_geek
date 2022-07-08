package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var ErrNoRows = sql.ErrNoRows

func main() {
	/**
	   sql.ErrNoRows不应该Wrap往上抛，应该直接处理
	**/
	line := 0
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_geek")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sq := "select * from test"
	rows, err := db.Query(sq)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan()
		line++
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	if line == 0 {
		log.Fatal(ErrNoRows)
	}
}
