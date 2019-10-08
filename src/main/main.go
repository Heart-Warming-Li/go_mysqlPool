package main

import (
	"fmt"
	m "mysql"
	"os"
)

func path() string {
	dir, _ := os.Getwd()
	return dir
}

func main() {
	file := string(path() + "/config/config.yaml")
	m.Init(file)
	//m.Insert("insert into user (name,password,email,sex,c_time) values ('456','123456','123450912@qq.com','asd','2019-10-08 15:09:15.675000')")
	//m.Update("update user set email='1234501912@qq.com' where name='456' and id=41 ")
	//m.Delete("delete from user where name='456' and id=41")
	//res := m.Select("select c_time from user where name = '456' ")
	//for _,v := range res {
	//	fmt.Println(v["c_time"])
	//}
	fmt.Println(os.Args)
	fmt.Println(os.Args[2])
	sqlExecMod := os.Args[1]
	sqlExecStr := os.Args[2]
	if sqlExecMod == "Select"{
		res := m.Select(sqlExecStr)
		//str, _ := ljson.MapToJson(false, res)
		//fmt.Println(string(str))
		for _,v := range res {
			//fmt.Println(v)
			fmt.Println(v["github_url"])
		}
	} else if sqlExecMod == "Insert" {
		m.Insert(sqlExecStr)
		fmt.Println("Insert: ",sqlExecStr)
	} else if sqlExecMod == "Update" {
		m.Update(sqlExecStr)
		fmt.Println("Update: ",sqlExecStr)
	} else if sqlExecMod == "Delete" {
		m.Delete(sqlExecStr)
		fmt.Println("Delete: ",sqlExecStr)
	}


}