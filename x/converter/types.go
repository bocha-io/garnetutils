package converter

type Field struct {
	Key  string
	Type string
}

type Table struct {
	Key       string
	Values    []Field
	Singleton bool
}

type Enum struct {
	Key    string
	Values []string
}

func GetEnumKeys(enums []Enum) []string {
	ret := []string{}
	for _, v := range enums {
		ret = append(ret, v.Key)
	}
	return ret
}
