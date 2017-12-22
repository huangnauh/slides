package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func connect2() {
	start := time.Now()
	db, _ := sql.Open("mysql", "runner:runner123456@tcp(127.0.0.1:3251)/upyun")
	defer db.Close()
	err := db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("connected: %s\n", elapsed)
}

func main() {
	for i := 0; i < 3; i++ {
		go connect2()
	}
	time.Sleep(5 * time.Second)
}
