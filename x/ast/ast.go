package ast

import (
	"fmt"

	"github.com/buger/jsonparser"
)

func ProcessAST(data []byte) error {
	err := jsonparser.ObjectEach(
		data,
		func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			fmt.Println(string(key))
			return nil
		},
		"ast",
	)

	return err
}
