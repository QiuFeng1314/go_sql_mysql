package dao

import (
	"log"
	"mysql/src/driver"
	"mysql/src/model"
)

func Update(people model.People) (res int64) {
	update_sql := "update go_table set age = ? where id = ?"

	r, err := driver.DB.Exec(update_sql, people.Age, people.Id)
	if err != nil {
		log.Panic("数据修改异常")
	}

	res, err2 := r.RowsAffected()
	if err2 != nil {
		log.Panic("数据修改异常")
	}
	return
}
