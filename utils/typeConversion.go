package utils

import (
	"encoding/json"
	"strconv"
)

// 字符数组转int数字
func StringToInt(strArr []string) []int {
	res := make([]int, len(strArr))

	for index, val := range strArr {
		res[index], _ = strconv.Atoi(val)
	}

	return res
}

func MapToJson(data interface{}) string {
	byteStr, _ := json.Marshal(data)
	return string(byteStr)
}
