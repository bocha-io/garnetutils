package utils

import (
	"strings"
)

const (
	Int64Type   = "int64"
	Uint32Type  = "uint32"
	Int32Type   = "int32"
	Uint8Type   = "uint8"
	AddressType = "address"
	BoolType    = "bool"
	StringType  = "string"
	Bytes32Type = "bytes32"
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
			return Int64Type
		}
	}

	switch val {
	case Bytes32Type:
		return brackets + StringType
	case BoolType:
		return brackets + BoolType
	case Int64Type:
		return brackets + Int64Type
	case Uint32Type:
		return brackets + Int64Type
	case Int32Type:
		return brackets + Int64Type
	case Uint8Type:
		return brackets + Int64Type
	case AddressType:
		return brackets + StringType
	default:
		return brackets + val
	}
}
