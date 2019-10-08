# go_mysqlPool
golang mysqlPool &amp;&amp; yaml config hot reload

------

# Yaml配置文件使用
###### src/config/config.yaml
```yaml
test:
  host: 127.0.0.1   # IP地址
  port: 3306		# 默认
  charset: utf8		# 默认
  username: test	# 用户名
  password: test	# 密码
  databases: test	# 数据库名字
  driverName: mysql	# 默认
  
production:
  host: 127.0.0.1
  port: 3306
  charset: utf8
  username: test
  password: test
  databases: test
  driverName: mysql  
  
debug: true			# true表示使用test配置， false表示production配置
```
配置参数可随时切换，此yaml配置文件支持热重载

------

# 执行sql增删改查
###### src/main/main.go
```go
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
	m.Insert("insert into user (name,password,email,sex,c_time) values ('456','123456','123450912@qq.com','asd','2019-10-08 15:09:15.675000')")
	m.Update("update user set email='1234501912@qq.com' where name='456' and id=41 ")
	m.Delete("delete from user where name='456' and id=41")
	res := m.Select("select c_time from user where name = '456' ")
	for _,v := range res {
		fmt.Println(v["c_time"])
	}
}
```

------

# 数据转换为json格式
###### src/ljson/ljson.go
```go
func StrToJson(str string) ([]byte, error){
	// str 	  string
	// json格式的字串
	// {"servers":[{"serverName":"TianJin","serverIP":"127.0.0.1"},{"serverName":"Beijing","serverIP":"127.0.0.2"}]}
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(str), &dat); err == nil {
		js, _ := json.Marshal(dat)
		return js,nil
	}else {
		return nil,err
	}
}

func MapToJson(marshalIndent bool, maps map[int]map[string]string)  ([]byte, error){
	// maps map[int]map[string]string
	// mysql 数据库直接查询获得的数据
	// map[0:map[c_time:2019-04-25 18:49:15.675000] 1:map[c_time:2019-04-08 18:49:15.675000]]

	// MarshalIndent 数据格式看上去更加美化,有缩进
	// Marshal       没有格式化的数据
	if marshalIndent{
		if jsonByte, err := json.MarshalIndent(maps,""," "); err == nil {
			return jsonByte, err
		} else {
			return nil, err
		}
	}else {
		if jsonByte, err := json.Marshal(maps); err == nil {
			return jsonByte, err
		} else {
			return nil, err
		}
	}
}
```

支持json类型的字符串转json格式，也支持从mysql中直接select的数据格式化为json
> 1. StrToJson("{"servers":[{"serverName":"TianJin","serverIP":"127.0.0.1"},{"serverName":"Beijing","serverIP":"127.0.0.2"}]}")
> 2. str, _ := MapToJson(true, m.Select("select c_time from user where name = '456' ")); fmt.Println(string(str))  # 格式化的数据
> 3. str, _ := MapToJson(false, m.Select("select c_time from user where name = '456' ")); fmt.Println(string(str)) # 没有格式化的数据
