package tools

import "encoding/json"

//结构体转为json
func StructToJson(obj interface{}) string {
	str, _ := json.Marshal(obj)
	return string(str)
}

//json转为结构体
func JsonToStruct(str string, obj interface{}) {
	_ = json.Unmarshal([]byte(str), obj)
}

// json interface转为结构体
func JsonI2Struct(str interface{}, obj interface{}) {
	// 将json interface转为string
	jsonStr, _ := str.(string)
	JsonToStruct(jsonStr, obj)
}
