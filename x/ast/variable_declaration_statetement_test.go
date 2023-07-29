package ast

import (
	"testing"
)

const expectedVariableDeclarationOne = "bool collisionX := x + size >= targetX && targetX + targetSize >= x"

func TestProcessVariableDeclarationStatement(t *testing.T) {
	testData := `
{
  "assignments": [
    67688
  ],
  "declarations": [
    {
      "constant": false,
      "id": 67688,
      "mutability": "mutable",
      "name": "collisionX",
      "nameLocation": "998:10:123",
      "nodeType": "VariableDeclaration",
      "scope": 67719,
      "src": "993:15:123",
      "stateVariable": false,
      "storageLocation": "default",
      "typeDescriptions": {
        "typeIdentifier": "t_bool",
        "typeString": "bool"
      },
      "typeName": {
        "id": 67687,
        "name": "bool",
        "nodeType": "ElementaryTypeName",
        "src": "993:4:123",
        "typeDescriptions": {
          "typeIdentifier": "t_bool",
          "typeString": "bool"
        }
      },
      "visibility": "internal"
    }
  ],
  "id": 67700,
  "initialValue": {
    "commonType": {
      "typeIdentifier": "t_bool",
      "typeString": "bool"
    },
    "id": 67699,
    "isConstant": false,
    "isLValue": false,
    "isPure": false,
    "lValueRequested": false,
    "leftExpression": {
      "commonType": {
        "typeIdentifier": "t_int32",
        "typeString": "int32"
      },
      "id": 67693,
      "isConstant": false,
      "isLValue": false,
      "isPure": false,
      "lValueRequested": false,
      "leftExpression": {
        "commonType": {
          "typeIdentifier": "t_int32",
          "typeString": "int32"
        },
        "id": 67691,
        "isConstant": false,
        "isLValue": false,
        "isPure": false,
        "lValueRequested": false,
        "leftExpression": {
          "id": 67689,
          "name": "x",
          "nodeType": "Identifier",
          "overloadedDeclarations": [],
          "referencedDeclaration": 67672,
          "src": "1011:1:123",
          "typeDescriptions": {
            "typeIdentifier": "t_int32",
            "typeString": "int32"
          }
        },
        "nodeType": "BinaryOperation",
        "operator": "+",
        "rightExpression": {
          "id": 67690,
          "name": "size",
          "nodeType": "Identifier",
          "overloadedDeclarations": [],
          "referencedDeclaration": 67676,
          "src": "1015:4:123",
          "typeDescriptions": {
            "typeIdentifier": "t_int32",
            "typeString": "int32"
          }
        },
        "src": "1011:8:123",
        "typeDescriptions": {
          "typeIdentifier": "t_int32",
          "typeString": "int32"
        }
      },
      "nodeType": "BinaryOperation",
      "operator": ">=",
      "rightExpression": {
        "id": 67692,
        "name": "targetX",
        "nodeType": "Identifier",
        "overloadedDeclarations": [],
        "referencedDeclaration": 67678,
        "src": "1023:7:123",
        "typeDescriptions": {
          "typeIdentifier": "t_int32",
          "typeString": "int32"
        }
      },
      "src": "1011:19:123",
      "typeDescriptions": {
        "typeIdentifier": "t_bool",
        "typeString": "bool"
      }
    },
    "nodeType": "BinaryOperation",
    "operator": "&&",
    "rightExpression": {
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
    },
    "src": "1011:48:123",
    "typeDescriptions": {
      "typeIdentifier": "t_bool",
      "typeString": "bool"
    }
  },
  "nodeType": "VariableDeclarationStatement",
  "src": "993:66:123"
}
`

	val, err := processVariableDeclarationStatement([]byte(testData))
	if err != nil {
		t.Error(err)
	}
	if val != expectedVariableDeclarationOne {
		t.Errorf("failed variable declaration: %s, %s", val, expectedVariableDeclarationOne)
	}

}
