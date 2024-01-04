package versioncomparer

import (
	"strconv"
	"strings"
)

func Compare(version1 string, version2 string) (res int, err error) {
	v1Array := strings.Split(version1, ".")
	v2Array := strings.Split(version2, ".")
	maxLen := max(len(v1Array), len(v2Array))
	for i := 0; i < maxLen; i++ {
		val1 := 0
		if i < len(v1Array) {
			val1, err = strconv.Atoi(v1Array[i])
			if err != nil {
				return
			}
		}

		val2 := 0
		if i < len(v1Array) {
			val2, err = strconv.Atoi(v2Array[i])
			if err != nil {
				return
			}
		}

		if val1 > val2 {
			res = 1
			return
		} else if val2 > val1 {
			res = -1
			return
		}
	}

	res = 0
	return
}
