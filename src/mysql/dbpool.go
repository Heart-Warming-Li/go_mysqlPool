package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"yaml"
)

type db struct {
	driver string
	dataSource string
}

func config(file string) *db {
	yaml.Check(file)
	host := yaml.Get("debug.host")
	port := yaml.Get("debug.port")
	charset := yaml.Get("debug.charset")
	username := yaml.Get("debug.username")
	password := yaml.Get("debug.password")
	databases := yaml.Get("debug.databases")
	driverName :=yaml.Get("debug.driverName")
	dataSourceName := strings.Join([]string{username, ":", password, "@tcp(",host, ":", port, ")/",
		databases, "?charset=", charset},"")
	return &db{driver:driverName, dataSource:dataSourceName}
}

var DB *sql.DB
func Init(file string) {
	db := config(file)
	DB, _ = sql.Open(db.driver, db.dataSource)
	DB.SetMaxOpenConns(65535)  // 最大连接数量
	DB.SetMaxIdleConns(10)	  // 最大空闲连接数量
	dbping()
}

func dbping() {
	if err := DB.Ping(); err != nil{
		fmt.Println("connnect database fail ... ...")
		return
	}
	//fmt.Println("connnect database success")
}

func Select(s string) map[int]map[string]string {
	rows, _ := DB.Query(s)
	//返回所有列
	cols, _ := rows.Columns()
	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))
	//这里表示一行填充数据
	scans := make([]interface{}, len(cols))
	//这里scans引用vals，把数据填充到[]byte里
	for k := range vals {
		scans[k] = &vals[k]
	}
	i := 0
	result := make(map[int]map[string]string)
	for rows.Next() {
		//填充数据
		rows.Scan(scans...)
		//每行数据
		res := make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			//这里把[]byte数据转成string
			res[key] = string(v)
		}
		//放入结果集
		result[i] = res
		i++
	}
	return result
}

func Insert(s string) bool {
	begin, err := DB.Begin()
	errInfo("Insert Begin", err)
	_, execErr := begin.Exec(s)
	if execErr != nil{
		begin.Rollback()
		errInfo("Insert", execErr)
	}
	begin.Commit()
	return true
}

func Update(s string) bool {
	begin, err := DB.Begin()
	errInfo("Update Begin", err)
	_, execErr := begin.Exec(s)
	if execErr != nil{
		begin.Rollback()
		errInfo("Update", execErr)
	}
	begin.Commit()
	return true
}

func Delete(s string) bool {
	begin, err := DB.Begin()
	errInfo("Delete Begin", err)
	_, execErr := begin.Exec(s)
	if execErr != nil{
		begin.Rollback()
		errInfo("Delete", execErr)
	}
	begin.Commit()
	return true
}

func errInfo(info string, err error) {
	if err != nil {
		panic(fmt.Sprintf("%s, Error message: %v", info, err))
	}
}