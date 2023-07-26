package ast

import (
	_ "embed"
	"testing"
)

//go:embed testdata/TicSystem.json
var ticFile []byte

func TestAST(t *testing.T) {
	_ = ProcessAST(ticFile)
	t.Fatal()
	// if err == nil {
	// 	t.Fatal()
	// }
}
