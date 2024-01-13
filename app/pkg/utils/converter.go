package utils

import "strconv"

func ConvertIntListToStringList(list []int) []string {
	result := []string{}
	for _, item := range list {
		result = append(result, strconv.Itoa(item))
	}
	return result
}
