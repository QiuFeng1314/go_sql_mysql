package dao

import (
	"log"
	"mysql/src/driver"
	"mysql/src/model"
)

func Add(people model.People) (id int64) {
	conn, _ := driver.DB.Begin()
	insert_sql := "insert into go_table (id, name, age) value(?, ?, ?)"

	r, err1 := conn.Exec(insert_sql, people.Id, people.Name, people.Age)
	if err1 != nil {
		log.Panic("数据新增异常")
		conn.Rollback()
	}

	id, err2 := r.RowsAffected()
	if err2 != nil {
		log.Panic("数据新增异常")
		conn.Rollback()
	}

	conn.Commit()
	return
}
