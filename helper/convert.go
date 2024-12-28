package helper

import (
	"strconv"
	"strings"
)

func ConvertListOfIntToString(businessIds []int) string {
	businessIdsStr := make([]string, len(businessIds))
	for i, id := range businessIds {
		businessIdsStr[i] = strconv.Itoa(id)
	}
	// Join the array to create a comma-separated string
	return strings.Join(businessIdsStr, ",")
}
