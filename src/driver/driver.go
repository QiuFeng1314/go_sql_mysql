package driver

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func init() {
	log.Printf("driver init success...")
	dbName := "mysql"
	db, err := sql.Open(dbName, "root:123321@tcp(localhost:3306)/golang?charset=utf8mb4")
	if err == nil && db.Ping() == nil {
		// log.Printf("%s数据库连接成功", dbName)
		DB = db
	} else {
		log.Panic(dbName, "数据库连接异常")
	}
}
