package utils

func SolidityTypeToGolang(val string) string {
	switch val {
	case "bytes32":
		return "string"
	case "bool":
		return "bool"
	default:
		return "int64"
	}
}
