package main

import (
	"fmt"
	"mysql/src/dao"
	"mysql/src/model"
)

func main() {

	fmt.Println("新增用户：")
	p := model.People{Id: 10004, Name: "李一", Age: 22}
	fmt.Printf("dao.Add(): %v\n", dao.Add(p))
	fmt.Printf("查询所有用户: %+v\n", dao.Query())

	// p.Age = 30
	// fmt.Println("修改用户：")
	// fmt.Printf("dao.Update(p): %v\n", dao.Update(p))
	// fmt.Printf("查询所有用户: %+v\n", dao.Query())

	// fmt.Println("删除用户：")
	// fmt.Printf("dao.Delete(p.Id): %v\n", dao.Delete(p.Id))
	// fmt.Printf("查询所有用户: %+v\n", dao.Query())
}
