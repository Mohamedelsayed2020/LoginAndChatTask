package server

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func Conn() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=mohammed dbname=mydb password=root sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
