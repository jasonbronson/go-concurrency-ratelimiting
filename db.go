package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	conn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True", "root", "admin", "127.0.0.1", "3306", "import")
	var err error
	db, err = gorm.Open("mysql", conn)
	//defer db.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database initialized ", conn)
	db.LogMode(true)
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(100)
}
