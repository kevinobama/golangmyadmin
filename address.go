package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Address struct {
	ID   int    `json:"address_id"`
	Name string `json:"address_realname"`
}

type Fields struct {
	Field   string `json:"Field"`
	Comment string `json:"Comment"`
}

func conn() *sql.DB {
	db, err := sql.Open("mysql", "root:654321@tcp(127.0.0.1:3306)/mall")
	if err != nil {
		panic(err.Error())
	}

	//defer db.Close()
	return db
}

func insert() {
	var db = conn()

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO ds_wechatlog(content,created_at) VALUES ('test','2021-10-30 15:22')")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	fmt.Println("insert success")
	defer insert.Close()
}

func query() {
	//fmt.Println("query")
	var db = conn()

	// Execute the query
	results, err := db.Query("SELECT address_id,address_realname  FROM ds_address order by address_id desc")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var address Address

		err = results.Scan(&address.ID, &address.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		fmt.Println(strconv.Itoa(address.ID) + "," + address.Name)
	}
	defer db.Close()
}

func queryColumn() {
	//fmt.Println("query")
	// Open up our database connection.
	var db = conn()

	// Execute the query
	var sql string = "SHOW FULL COLUMNS FROM ds_address"

	results, err := db.Query(sql)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var fields Fields

		/**
		err = results.Scan(&fields.Field, &fields.Comment)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		} **/

		fmt.Println(fields.Field + "," + fields.Comment)
	}
	defer db.Close()
}
