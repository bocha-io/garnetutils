package ast

import "testing"

func TestParameterList(t *testing.T) {
	testData := `
{
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
}`

	expected := "x int64, y int64, size int64, targetX int64, targetY int64, targetSize int64"

	val, err := NewConverter().processParameterList([]byte(testData))
	if err != nil {
		t.Error(err)
	}
	if val != expected {
		t.Errorf("failed: %s, %s", val, expected)
	}
}
