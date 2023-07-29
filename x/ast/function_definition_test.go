package ast

import "testing"

func TestProcessFunctionDefinition(t *testing.T) {
	testData := `
{
    "id": 67720,
    "nodeType": "FunctionDefinition",
    "src": "825:400:123",
    "nodes": [],
    "body": {
      "id": 67719,
      "nodeType": "Block",
      "src": "983:242:123",
      "nodes": [],
      "statements": [
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
        },
        {
          "assignments": [
            67702
          ],
          "declarations": [
            {
              "constant": false,
              "id": 67702,
              "mutability": "mutable",
              "name": "collisionY",
              "nameLocation": "1074:10:123",
              "nodeType": "VariableDeclaration",
              "scope": 67719,
              "src": "1069:15:123",
              "stateVariable": false,
              "storageLocation": "default",
              "typeDescriptions": {
                "typeIdentifier": "t_bool",
                "typeString": "bool"
              },
              "typeName": {
                "id": 67701,
                "name": "bool",
                "nodeType": "ElementaryTypeName",
                "src": "1069:4:123",
                "typeDescriptions": {
                  "typeIdentifier": "t_bool",
                  "typeString": "bool"
                }
              },
              "visibility": "internal"
            }
          ],
          "id": 67714,
          "initialValue": {
            "commonType": {
              "typeIdentifier": "t_bool",
              "typeString": "bool"
            },
            "id": 67713,
            "isConstant": false,
            "isLValue": false,
            "isPure": false,
            "lValueRequested": false,
            "leftExpression": {
              "commonType": {
                "typeIdentifier": "t_int32",
                "typeString": "int32"
              },
              "id": 67707,
              "isConstant": false,
              "isLValue": false,
              "isPure": false,
              "lValueRequested": false,
              "leftExpression": {
                "commonType": {
                  "typeIdentifier": "t_int32",
                  "typeString": "int32"
                },
                "id": 67705,
                "isConstant": false,
                "isLValue": false,
                "isPure": false,
                "lValueRequested": false,
                "leftExpression": {
                  "id": 67703,
                  "name": "y",
                  "nodeType": "Identifier",
                  "overloadedDeclarations": [],
                  "referencedDeclaration": 67674,
                  "src": "1087:1:123",
                  "typeDescriptions": {
                    "typeIdentifier": "t_int32",
                    "typeString": "int32"
                  }
                },
                "nodeType": "BinaryOperation",
                "operator": "+",
                "rightExpression": {
                  "id": 67704,
                  "name": "size",
                  "nodeType": "Identifier",
                  "overloadedDeclarations": [],
                  "referencedDeclaration": 67676,
                  "src": "1091:4:123",
                  "typeDescriptions": {
                    "typeIdentifier": "t_int32",
                    "typeString": "int32"
                  }
                },
                "src": "1087:8:123",
                "typeDescriptions": {
                  "typeIdentifier": "t_int32",
                  "typeString": "int32"
                }
              },
              "nodeType": "BinaryOperation",
              "operator": ">=",
              "rightExpression": {
                "id": 67706,
                "name": "targetY",
                "nodeType": "Identifier",
                "overloadedDeclarations": [],
                "referencedDeclaration": 67680,
                "src": "1099:7:123",
                "typeDescriptions": {
                  "typeIdentifier": "t_int32",
                  "typeString": "int32"
                }
              },
              "src": "1087:19:123",
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
              "id": 67712,
              "isConstant": false,
              "isLValue": false,
              "isPure": false,
              "lValueRequested": false,
              "leftExpression": {
                "commonType": {
                  "typeIdentifier": "t_int32",
                  "typeString": "int32"
                },
                "id": 67710,
                "isConstant": false,
                "isLValue": false,
                "isPure": false,
                "lValueRequested": false,
                "leftExpression": {
                  "id": 67708,
                  "name": "targetY",
                  "nodeType": "Identifier",
                  "overloadedDeclarations": [],
                  "referencedDeclaration": 67680,
                  "src": "1110:7:123",
                  "typeDescriptions": {
                    "typeIdentifier": "t_int32",
                    "typeString": "int32"
                  }
                },
                "nodeType": "BinaryOperation",
                "operator": "+",
                "rightExpression": {
                  "id": 67709,
                  "name": "targetSize",
                  "nodeType": "Identifier",
                  "overloadedDeclarations": [],
                  "referencedDeclaration": 67682,
                  "src": "1120:10:123",
                  "typeDescriptions": {
                    "typeIdentifier": "t_int32",
                    "typeString": "int32"
                  }
                },
                "src": "1110:20:123",
                "typeDescriptions": {
                  "typeIdentifier": "t_int32",
                  "typeString": "int32"
                }
              },
              "nodeType": "BinaryOperation",
              "operator": ">=",
              "rightExpression": {
                "id": 67711,
                "name": "y",
                "nodeType": "Identifier",
                "overloadedDeclarations": [],
                "referencedDeclaration": 67674,
                "src": "1134:1:123",
                "typeDescriptions": {
                  "typeIdentifier": "t_int32",
                  "typeString": "int32"
                }
              },
              "src": "1110:25:123",
              "typeDescriptions": {
                "typeIdentifier": "t_bool",
                "typeString": "bool"
              }
            },
            "src": "1087:48:123",
            "typeDescriptions": {
              "typeIdentifier": "t_bool",
              "typeString": "bool"
            }
          },
          "nodeType": "VariableDeclarationStatement",
          "src": "1069:66:123"
        },
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
      ]
    },
    "implemented": true,
    "kind": "function",
    "modifiers": [],
    "name": "checkCollision",
    "nameLocation": "834:14:123",
    "parameters": {
      "id": 67683,
      "nodeType": "ParameterList",
      "parameters": [
        {
          "constant": false,
          "id": 67672,
          "mutability": "mutable",
          "name": "x",
          "nameLocation": "855:1:123",
          "nodeType": "VariableDeclaration",
          "scope": 67720,
          "src": "849:7:123",
          "stateVariable": false,
          "storageLocation": "default",
          "typeDescriptions": {
            "typeIdentifier": "t_int32",
            "typeString": "int32"
          },
          "typeName": {
            "id": 67671,
            "name": "int32",
            "nodeType": "ElementaryTypeName",
            "src": "849:5:123",
            "typeDescriptions": {
              "typeIdentifier": "t_int32",
              "typeString": "int32"
            }
          },
          "visibility": "internal"
        },
        {
          "constant": false,
          "id": 67674,
          "mutability": "mutable",
          "name": "y",
          "nameLocation": "864:1:123",
          "nodeType": "VariableDeclaration",
          "scope": 67720,
          "src": "858:7:123",
          "stateVariable": false,
          "storageLocation": "default",
          "typeDescriptions": {
            "typeIdentifier": "t_int32",
            "typeString": "int32"
          },
          "typeName": {
            "id": 67673,
            "name": "int32",
            "nodeType": "ElementaryTypeName",
            "src": "858:5:123",
            "typeDescriptions": {
              "typeIdentifier": "t_int32",
              "typeString": "int32"
            }
          },
          "visibility": "internal"
        },
        {
          "constant": false,
          "id": 67676,
          "mutability": "mutable",
          "name": "size",
          "nameLocation": "873:4:123",
          "nodeType": "VariableDeclaration",
          "scope": 67720,
          "src": "867:10:123",
          "stateVariable": false,
          "storageLocation": "default",
          "typeDescriptions": {
            "typeIdentifier": "t_int32",
            "typeString": "int32"
          },
          "typeName": {
            "id": 67675,
            "name": "int32",
            "nodeType": "ElementaryTypeName",
            "src": "867:5:123",
            "typeDescriptions": {
              "typeIdentifier": "t_int32",
              "typeString": "int32"
            }
          },
          "visibility": "internal"
        },
        {
          "constant": false,
          "id": 67678,
          "mutability": "mutable",
          "name": "targetX",
          "nameLocation": "885:7:123",
          "nodeType": "VariableDeclaration",
          "scope": 67720,
          "src": "879:13:123",
          "stateVariable": false,
          "storageLocation": "default",
          "typeDescriptions": {
            "typeIdentifier": "t_int32",
            "typeString": "int32"
          },
          "typeName": {
            "id": 67677,
            "name": "int32",
            "nodeType": "ElementaryTypeName",
            "src": "879:5:123",
            "typeDescriptions": {
              "typeIdentifier": "t_int32",
              "typeString": "int32"
            }
          },
          "visibility": "internal"
        },
        {
          "constant": false,
          "id": 67680,
          "mutability": "mutable",
          "name": "targetY",
          "nameLocation": "900:7:123",
          "nodeType": "VariableDeclaration",
          "scope": 67720,
          "src": "894:13:123",
          "stateVariable": false,
          "storageLocation": "default",
          "typeDescriptions": {
            "typeIdentifier": "t_int32",
            "typeString": "int32"
          },
          "typeName": {
            "id": 67679,
            "name": "int32",
            "nodeType": "ElementaryTypeName",
            "src": "894:5:123",
            "typeDescriptions": {
              "typeIdentifier": "t_int32",
              "typeString": "int32"
            }
          },
          "visibility": "internal"
        },
        {
          "constant": false,
          "id": 67682,
          "mutability": "mutable",
          "name": "targetSize",
          "nameLocation": "915:10:123",
          "nodeType": "VariableDeclaration",
          "scope": 67720,
          "src": "909:16:123",
          "stateVariable": false,
          "storageLocation": "default",
          "typeDescriptions": {
            "typeIdentifier": "t_int32",
            "typeString": "int32"
          },
          "typeName": {
            "id": 67681,
            "name": "int32",
            "nodeType": "ElementaryTypeName",
            "src": "909:5:123",
            "typeDescriptions": {
              "typeIdentifier": "t_int32",
              "typeString": "int32"
            }
          },
          "visibility": "internal"
        }
      ],
      "src": "848:78:123"
    },
    "returnParameters": {
      "id": 67686,
      "nodeType": "ParameterList",
      "parameters": [
        {
          "constant": false,
          "id": 67685,
          "mutability": "mutable",
          "name": "",
          "nameLocation": "-1:-1:-1",
          "nodeType": "VariableDeclaration",
          "scope": 67720,
          "src": "973:4:123",
          "stateVariable": false,
          "storageLocation": "default",
          "typeDescriptions": {
            "typeIdentifier": "t_bool",
            "typeString": "bool"
          },
          "typeName": {
            "id": 67684,
            "name": "bool",
            "nodeType": "ElementaryTypeName",
            "src": "973:4:123",
            "typeDescriptions": {
              "typeIdentifier": "t_bool",
              "typeString": "bool"
            }
          },
          "visibility": "internal"
        }
      ],
      "src": "972:6:123"
    },
    "scope": 68505,
    "stateMutability": "pure",
    "virtual": false,
    "visibility": "private"
  }
`
	expected := `func checkCollision (x int32, y int32, size int32, targetX int32, targetY int32, targetSize int32)  ( bool) {
` + expectedBlock + `
}`

	val, err := processFunctionDefinition([]byte(testData))
	if err != nil {
		t.Error(err)
	}
	if val != expected {
		t.Errorf("failed: %s, %s", val, expected)
	}

}
