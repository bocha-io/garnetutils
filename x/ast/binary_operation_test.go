package ast

import "testing"

func TestProcessBinaryOperation(t *testing.T) {
	testData := `
{
  "commonType": {
    "typeIdentifier": "t_int32",
    "typeString": "int32"
  },
  "id": 67698,
  "isConstant": false,
  "isLValue": false,
  "isPure": false,
  "lValueRequested": false,
  "leftExpression": {
    "commonType": {
      "typeIdentifier": "t_int32",
      "typeString": "int32"
    },
    "id": 67696,
    "isConstant": false,
    "isLValue": false,
    "isPure": false,
    "lValueRequested": false,
    "leftExpression": {
      "id": 67694,
      "name": "targetX",
      "nodeType": "Identifier",
      "overloadedDeclarations": [],
      "referencedDeclaration": 67678,
      "src": "1034:7:123",
      "typeDescriptions": {
        "typeIdentifier": "t_int32",
        "typeString": "int32"
      }
    },
    "nodeType": "BinaryOperation",
    "operator": "+",
    "rightExpression": {
      "id": 67695,
      "name": "targetSize",
      "nodeType": "Identifier",
      "overloadedDeclarations": [],
      "referencedDeclaration": 67682,
      "src": "1044:10:123",
      "typeDescriptions": {
        "typeIdentifier": "t_int32",
        "typeString": "int32"
      }
    },
    "src": "1034:20:123",
    "typeDescriptions": {
      "typeIdentifier": "t_int32",
      "typeString": "int32"
    }
  },
  "nodeType": "BinaryOperation",
  "operator": ">=",
  "rightExpression": {
    "id": 67697,
    "name": "x",
    "nodeType": "Identifier",
    "overloadedDeclarations": [],
    "referencedDeclaration": 67672,
    "src": "1058:1:123",
    "typeDescriptions": {
      "typeIdentifier": "t_int32",
      "typeString": "int32"
    }
  },
  "src": "1034:25:123",
  "typeDescriptions": {
    "typeIdentifier": "t_bool",
    "typeString": "bool"
  }
}
`

	expected := "targetX + targetSize >= x"
	val, err := NewConverter().processBinaryOperation([]byte(testData))
	if err != nil {
		t.Error(err)
	}
	if val != expected {
		t.Errorf("failed binary operation: %s, %s", val, expected)
	}
}
