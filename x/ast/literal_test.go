package ast

import "testing"

func TestLiteral(t *testing.T) {
	testData := `{
"hexValue": "30",
"id": 67927,
"isConstant": false,
"isLValue": false,
"isPure": true,
"kind": "number",
"lValueRequested": false,
"nodeType": "Literal",
"src": "2756:1:123",
"typeDescriptions": {
  "typeIdentifier": "t_rational_0_by_1",
  "typeString": "int_const 0"
},
"value": "0"
}`

	expected := "0"

	val, err := processLiteral([]byte(testData))
	if err != nil {
		t.Error(err)
	}
	if val != expected {
		t.Errorf("failed: %s, %s", val, expected)
	}
}
