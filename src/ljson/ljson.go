package ljson

import (
	"encoding/json"
)

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

	// MarshalIndent 看上去更加格式化
	// Marshal 不需要格式化
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