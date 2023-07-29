package ast

import (
	_ "embed"
	"testing"
)

//go:embed testdata/TicSystem.json
var ticFile []byte

func TestAST(t *testing.T) {
	err := ProcessAST(ticFile)
	if err != nil {
		t.Errorf(err.Error())
	}
}
