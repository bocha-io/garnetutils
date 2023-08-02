package ast

import "testing"

func TestIfStatement(t *testing.T) {
	testData := `
{
  "condition": {
    "commonType": {
      "typeIdentifier": "t_int32",
      "typeString": "int32"
    },
    "id": 67849,
    "isConstant": false,
    "isLValue": false,
    "isPure": false,
    "lValueRequested": false,
    "leftExpression": {
      "id": 67847,
      "name": "enemyX",
      "nodeType": "Identifier",
      "overloadedDeclarations": [],
      "referencedDeclaration": 67839,
      "src": "2134:6:123",
      "typeDescriptions": {
        "typeIdentifier": "t_int32",
        "typeString": "int32"
      }
    },
    "nodeType": "BinaryOperation",
    "operator": ">",
    "rightExpression": {
      "id": 67848,
      "name": "playerX",
      "nodeType": "Identifier",
      "overloadedDeclarations": [],
      "referencedDeclaration": 67829,
      "src": "2143:7:123",
      "typeDescriptions": {
        "typeIdentifier": "t_int32",
        "typeString": "int32"
      }
    },
    "src": "2134:16:123",
    "typeDescriptions": {
      "typeIdentifier": "t_bool",
      "typeString": "bool"
    }
  },
  "falseBody": {
    "condition": {
      "commonType": {
        "typeIdentifier": "t_int32",
        "typeString": "int32"
      },
      "id": 67857,
      "isConstant": false,
      "isLValue": false,
      "isPure": false,
      "lValueRequested": false,
      "leftExpression": {
        "id": 67855,
        "name": "enemyX",
        "nodeType": "Identifier",
        "overloadedDeclarations": [],
        "referencedDeclaration": 67839,
        "src": "2198:6:123",
        "typeDescriptions": {
          "typeIdentifier": "t_int32",
          "typeString": "int32"
        }
      },
      "nodeType": "BinaryOperation",
      "operator": "<",
      "rightExpression": {
        "id": 67856,
        "name": "playerX",
        "nodeType": "Identifier",
        "overloadedDeclarations": [],
        "referencedDeclaration": 67829,
        "src": "2207:7:123",
        "typeDescriptions": {
          "typeIdentifier": "t_int32",
          "typeString": "int32"
        }
      },
      "src": "2198:16:123",
      "typeDescriptions": {
        "typeIdentifier": "t_bool",
        "typeString": "bool"
      }
    },
    "id": 67863,
    "nodeType": "IfStatement",
    "src": "2194:58:123",
    "trueBody": {
      "id": 67862,
      "nodeType": "Block",
      "src": "2216:36:123",
      "statements": [
        {
          "expression": {
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
            }
          },
          "id": 67861,
          "nodeType": "ExpressionStatement",
          "src": "2230:11:123"
        }
      ]
    }
  },
  "id": 67864,
  "nodeType": "IfStatement",
  "src": "2130:122:123",
  "trueBody": {
    "id": 67854,
    "nodeType": "Block",
    "src": "2152:36:123",
    "statements": [
      {
        "expression": {
          "id": 67852,
          "isConstant": false,
          "isLValue": false,
          "isPure": false,
          "lValueRequested": false,
          "leftHandSide": {
            "id": 67850,
            "name": "enemyX",
            "nodeType": "Identifier",
            "overloadedDeclarations": [],
            "referencedDeclaration": 67839,
            "src": "2166:6:123",
            "typeDescriptions": {
              "typeIdentifier": "t_int32",
              "typeString": "int32"
            }
          },
          "nodeType": "Assignment",
          "operator": "-=",
          "rightHandSide": {
            "hexValue": "35",
            "id": 67851,
            "isConstant": false,
            "isLValue": false,
            "isPure": true,
            "kind": "number",
            "lValueRequested": false,
            "nodeType": "Literal",
            "src": "2176:1:123",
            "typeDescriptions": {
              "typeIdentifier": "t_rational_5_by_1",
              "typeString": "int_const 5"
            },
            "value": "5"
          },
          "src": "2166:11:123",
          "typeDescriptions": {
            "typeIdentifier": "t_int32",
            "typeString": "int32"
          }
        },
        "id": 67853,
        "nodeType": "ExpressionStatement",
        "src": "2166:11:123"
      }
    ]
  }
}
`

	expected := `if enemyX > playerX {
enemyX -= 5
} else {
if enemyX < playerX {
enemyX += 5
}
}`
	val, err := NewASTConverter().processIfStatement([]byte(testData))
	if err != nil {
		t.Error(err)
	}
	if val != expected {
		t.Errorf("failed: %s, %s", val, expected)
	}
}
