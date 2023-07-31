package ast

import "testing"

func TestProcessAssignment(t *testing.T) {
	testData := `
{
    "id": 67860,
    "isConstant": false,
    "isLValue": false,
    "isPure": false,
    "lValueRequested": false,
    "leftHandSide": {
      "id": 67858,
      "name": "enemyX",
      "nodeType": "Identifier",
      "overloadedDeclarations": [],
      "referencedDeclaration": 67839,
      "src": "2230:6:123",
      "typeDescriptions": {
        "typeIdentifier": "t_int32",
        "typeString": "int32"
      }
    },
    "nodeType": "Assignment",
    "operator": "+=",
    "rightHandSide": {
      "hexValue": "35",
      "id": 67859,
      "isConstant": false,
      "isLValue": false,
      "isPure": true,
      "kind": "number",
      "lValueRequested": false,
      "nodeType": "Literal",
      "src": "2240:1:123",
      "typeDescriptions": {
        "typeIdentifier": "t_rational_5_by_1",
        "typeString": "int_const 5"
      },
      "value": "5"
    },
    "src": "2230:11:123",
    "typeDescriptions": {
      "typeIdentifier": "t_int32",
      "typeString": "int32"
    },
}
`

	expected := "enemyX += 5"

	val, err := processAssignment([]byte(testData))
	if err != nil {
		t.Error(err)
	}
	if val != expected {
		t.Errorf("failed: %s, %s", val, expected)
	}
}
