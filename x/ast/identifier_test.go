package ast

import "testing"

func TestIdentifier(t *testing.T) {
	testData := `{
"id": 67930,
"name": "x",
"nodeType": "Identifier",
"overloadedDeclarations": [],
"referencedDeclaration": 67914,
"src": "2773:1:123",
"typeDescriptions": {
  "typeIdentifier": "t_int32",
  "typeString": "int32"
}
}
`

	expected := "x"

	val, err := processIdentifier([]byte(testData))
	if err != nil {
		t.Error(err)
	}
	if val != expected {
		t.Errorf("failed: %s, %s", val, expected)
	}
}
