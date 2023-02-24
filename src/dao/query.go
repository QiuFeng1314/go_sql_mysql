package dao

import (
	"log"
	"mysql/src/driver"
	"mysql/src/model"
)

func Query() (peoples []model.People) {
	query_sql := "select id, name, age from go_table order by id"

	rows, err := driver.DB.Query(query_sql)
	defer rows.Close()

	if err != nil {
		log.Panic("数据查询异常")
	}

	for rows.Next() {
		var people model.People = model.People{}
		err = rows.Scan(&people.Id, &people.Name, &people.Age)
		if err != nil {
			log.Fatalln(err)
		}
		peoples = append(peoples, people)
	}
	return peoples
}
