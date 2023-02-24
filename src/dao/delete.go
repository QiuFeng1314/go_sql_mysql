package dao

import (
	"log"
	"mysql/src/driver"
)

func Delete(id int) (res int64) {
	delete_sql := "delete from go_table where id = ?"

	r, err := driver.DB.Exec(delete_sql, id)
	if err != nil {
		log.Panic("数据删除失败")
	}

	res, err2 := r.RowsAffected()
	if err2 != nil {
		log.Panic("数据删除失败")
	}

	return
}
