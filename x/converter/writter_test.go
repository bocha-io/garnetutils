package converter

import "testing"

func TestGenerateFiles(t *testing.T) {
	a := GenerateFiles("GameObject", mudConfig, "")
	if a != nil {
		t.Fail()
	}
}
