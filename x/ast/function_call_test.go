package ast

import "testing"

func TestFunctionCall_TypeConversion(t *testing.T) {
	testData := `
{
    "arguments": [
      {
        "hexValue": "3130",
        "id": 67747,
        "isConstant": false,
        "isLValue": false,
        "isPure": true,
        "kind": "number",
        "lValueRequested": false,
        "nodeType": "Literal",
        "src": "1422:2:123",
        "typeDescriptions": {
          "typeIdentifier": "t_rational_10_by_1",
          "typeString": "int_const 10"
        },
        "value": "10"
      }
    ],
    "expression": {
      "argumentTypes": [
        {
          "typeIdentifier": "t_rational_10_by_1",
          "typeString": "int_const 10"
        }
      ],
      "id": 67746,
      "isConstant": false,
      "isLValue": false,
      "isPure": true,
      "lValueRequested": false,
      "nodeType": "ElementaryTypeNameExpression",
      "src": "1416:5:123",
      "typeDescriptions": {
        "typeIdentifier": "t_type$_t_int32_$",
        "typeString": "type(int32)"
      },
      "typeName": {
        "id": 67745,
        "name": "int32",
        "nodeType": "ElementaryTypeName",
        "src": "1416:5:123",
        "typeDescriptions": {}
      }
    },
    "id": 67748,
    "isConstant": false,
    "isLValue": false,
    "isPure": true,
    "kind": "typeConversion",
    "lValueRequested": false,
    "names": [],
    "nodeType": "FunctionCall",
    "src": "1416:9:123",
    "tryCall": false,
    "typeDescriptions": {
      "typeIdentifier": "t_int32",
      "typeString": "int32"
    }
  }
  `

	expected := "int32(10)"
	val, err := processFunctionCall([]byte(testData))
	if err != nil {
		t.Error(err)
	}
	if val != expected {
		t.Errorf("failed: %s, %s", val, expected)
	}
}
