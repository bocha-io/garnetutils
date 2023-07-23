package converter

import "testing"

func TestGenerateFiles(t *testing.T) {

	a := GenerateFiles("GameObject", mudConfig, "")
	if len(a) == 0 {
		t.Fail()
	}

}
