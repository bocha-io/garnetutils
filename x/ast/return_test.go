package ast

import "testing"

func TestProcessReturn(t *testing.T) {
	testData := `
{
  "expression": {
    "commonType": {
      "typeIdentifier": "t_bool",
      "typeString": "bool"
    },
    "id": 67717,
    "isConstant": false,
    "isLValue": false,
    "isPure": false,
    "lValueRequested": false,
    "leftExpression": {
      "id": 67715,
      "name": "collisionX",
      "nodeType": "Identifier",
      "overloadedDeclarations": [],
      "referencedDeclaration": 67688,
      "src": "1194:10:123",
      "typeDescriptions": {
        "typeIdentifier": "t_bool",
        "typeString": "bool"
      }
    },
    "nodeType": "BinaryOperation",
    "operator": "&&",
    "rightExpression": {
      "id": 67716,
      "name": "collisionY",
      "nodeType": "Identifier",
      "overloadedDeclarations": [],
      "referencedDeclaration": 67702,
      "src": "1208:10:123",
      "typeDescriptions": {
        "typeIdentifier": "t_bool",
        "typeString": "bool"
      }
    },
    "src": "1194:24:123",
    "typeDescriptions": {
      "typeIdentifier": "t_bool",
      "typeString": "bool"
    }
  },
  "functionReturnParameters": 67686,
  "id": 67718,
  "nodeType": "Return",
  "src": "1187:31:123"
}
`

	expected := "return collisionX && collisionY"
	val, err := NewConverter().processReturn([]byte(testData))
	if err != nil {
		t.Error(err)
	}
	if val != expected {
		t.Errorf("failed return: %s, %s", val, expected)
	}
}
