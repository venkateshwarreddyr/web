package util

import (
	"fmt"
	"reflect"
)

func GetJsonLikeOutput(data interface{}) string {
	jsonLikeOutput := "{"

	dataValue := reflect.ValueOf(data)
	dataType := reflect.TypeOf(data)
	for i := 0; i < dataValue.NumField(); i++ {
		field := dataValue.Field(i)
		fieldName := dataType.Field(i).Name
		fieldValue := fmt.Sprintf("%v", field.Interface())

		jsonLikeOutput += "\"" + fieldName + "\":" + "\"" + fieldValue + "\","

	}

	jsonLikeOutput = jsonLikeOutput[:len(jsonLikeOutput)-1] + "}"

	return jsonLikeOutput
}
