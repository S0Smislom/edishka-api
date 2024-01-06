package utils

import (
	"fmt"
	"reflect"
)

func PrintMapTypes(data map[string]interface{}) {
	for key, value := range data {
		valueType := reflect.TypeOf(value).String()
		fmt.Printf("Key: %s, Type: %s\n", key, valueType)
	}
}
