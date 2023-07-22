package converter

import (
	_ "embed"
	"testing"
)

//go:embed testdata/mud.config.ts
var mudConfig []byte

//go:embed testdata/mud.config.json
var mudConfigJSON []byte

func TestMudConfigToJSON(t *testing.T) {
	res := MudConfigToJSON(mudConfig)
	if string(res) != string(mudConfigJSON) {
		t.Fatalf("invalid conversion from ts to json")
	}
}
