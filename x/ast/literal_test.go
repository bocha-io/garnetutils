package ast

import "testing"

func TestLiteral_number(t *testing.T) {
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

	val, err := NewASTConverter().processLiteral([]byte(testData))
	if err != nil {
		t.Error(err)
	}
	if val != expected {
		t.Errorf("failed: %s, %s", val, expected)
	}
}

func TestLiteral_string(t *testing.T) {
	testData := `{
  "hexValue": "41747461636b44616d616765",
  "id": 62317,
  "isConstant": false,
  "isLValue": false,
  "isPure": true,
  "kind": "string",
  "lValueRequested": false,
  "nodeType": "Literal",
  "src": "923:14:107",
  "typeDescriptions": {
    "typeIdentifier": "t_stringliteral_8b5fa4535314e12990cc739dd9674ca957ccf3b48ae918283da38f6a0bc84a60",
    "typeString": "literal_string \"AttackDamage\""
  },
  "value": "AttackDamage"
}`

	expected := "\"AttackDamage\""

	val, err := NewASTConverter().processLiteral([]byte(testData))
	if err != nil {
		t.Error(err)
	}
	if val != expected {
		t.Errorf("failed: %s, %s", val, expected)
	}
}

func TestLiteral_bool(t *testing.T) {
	testData := `{
  "hexValue": "74727565",
  "id": 82814,
  "isConstant": false,
  "isLValue": false,
  "isPure": true,
  "kind": "bool",
  "lValueRequested": false,
  "nodeType": "Literal",
  "src": "1777:4:152",
  "typeDescriptions": {
    "typeIdentifier": "t_bool",
    "typeString": "bool"
  },
  "value": "true"
}
`

	expected := "true"

	val, err := NewASTConverter().processLiteral([]byte(testData))
	if err != nil {
		t.Error(err)
	}
	if val != expected {
		t.Errorf("failed: %s, %s", val, expected)
	}
}
