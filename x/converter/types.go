package converter

type Field struct {
	Key  string
	Type string
}

type Table struct {
	Key    string
	Values []Field
}

type Enum struct {
	Key    string
	Values []string
}
