package tool

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitSQL() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/message_board?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	DB = db
	return nil
}
