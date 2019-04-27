package dbops

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

func openConn() *sql.DB {
	dbConn, err := sql.Open("mysql", "root:wangchen18606632@#@tcp(localhost:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}

	return dbConn
}

func AddUserCredentital(loginName string, pwd string) error {
	db := openConn()
}

func GetUserCredential(loginName string) (string, error) {
	db := openConn()
}

