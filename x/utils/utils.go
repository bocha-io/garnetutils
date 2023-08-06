package utils

import (
	"strings"
)

func SolidityTypeToGolang(val string, enums []string) string {
	brackets := ""
	splitted := strings.Split(val, "]")
	if len(splitted) == 2 {
		val = splitted[1]
		length := strings.Split(splitted[0], "[")
		if len(length) == 2 {
			brackets = "[" + length[1] + "]"
		} else {
			brackets = "[]"
		}
	}
	val = strings.Trim(val, " ")

	for _, v := range enums {
		if val == v {
			return "int64"
		}
	}

	switch val {
	case "bytes32":
		return brackets + "string"
	case "bool":
		return brackets + "bool"
	case "int64":
		return brackets + "int64"
	case "uint32":
		return brackets + "int64"
	case "int32":
		return brackets + "int64"
	case "uint8":
		return brackets + "int64"
	default:
		return brackets + val
	}
}
