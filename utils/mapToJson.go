package utils

import "encoding/json"

func MapToJson(data interface{}) string {
    byteStr, _ := json.Marshal(data)
    return string(byteStr)
}