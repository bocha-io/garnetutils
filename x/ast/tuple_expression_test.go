package ast

import "testing"

func TestTupleExpression(t *testing.T) {
	testData := `
{
    "components": [
      {
        "id": 67891,
        "name": "enemyX",
        "nodeType": "Identifier",
        "overloadedDeclarations": [],
        "referencedDeclaration": 67839,
        "src": "2448:6:123",
        "typeDescriptions": {
          "typeIdentifier": "t_int32",
          "typeString": "int32"
        }
      },
      {
        "id": 67892,
        "name": "enemyY",
        "nodeType": "Identifier",
        "overloadedDeclarations": [],
        "referencedDeclaration": 67841,
        "src": "2456:6:123",
        "typeDescriptions": {
          "typeIdentifier": "t_int32",
          "typeString": "int32"
        }
      }
    ],
    "id": 67893,
    "isConstant": false,
    "isInlineArray": false,
    "isLValue": false,
    "isPure": false,
    "lValueRequested": false,
    "nodeType": "TupleExpression",
    "src": "2447:16:123",
    "typeDescriptions": {
      "typeIdentifier": "t_tuple$_t_int32_$_t_int32_$",
      "typeString": "tuple(int32,int32)"
    }
  }
`

	expected := "enemyX, enemyY"

	val, err := processTupleExpression([]byte(testData))
	if err != nil {
		t.Error(err)
	}
	if val != expected {
		t.Errorf("failed: %s, %s", val, expected)
	}
}
