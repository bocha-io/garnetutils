package ast

import "testing"

func TestIndexAccess(t *testing.T) {
	testData := `{
  "baseExpression": {
    "id": 62343,
    "name": "_schema",
    "nodeType": "Identifier",
    "overloadedDeclarations": [],
    "referencedDeclaration": 62335,
    "src": "1162:7:107",
    "typeDescriptions": {
      "typeIdentifier": "t_array$_t_enum$_SchemaType_$44035_$dyn_memory_ptr",
      "typeString": "enum SchemaType[] memory"
    }
  },
  "id": 62345,
  "indexExpression": {
    "hexValue": "30",
    "id": 62344,
    "isConstant": false,
    "isLValue": false,
    "isPure": true,
    "kind": "number",
    "lValueRequested": false,
    "nodeType": "Literal",
    "src": "1170:1:107",
    "typeDescriptions": {
      "typeIdentifier": "t_rational_0_by_1",
      "typeString": "int_const 0"
    },
    "value": "0"
  },
  "isConstant": false,
  "isLValue": true,
  "isPure": false,
  "lValueRequested": true,
  "nodeType": "IndexAccess",
  "src": "1162:10:107",
  "typeDescriptions": {
    "typeIdentifier": "t_enum$_SchemaType_$44035",
    "typeString": "enum SchemaType"
  }
}`

	expected := "_schema[int64(0)]"

	val, err := NewConverter().processIndexAccess([]byte(testData))
	if err != nil {
		t.Error(err)
	}
	if val != expected {
		t.Errorf("failed: %s, %s", val, expected)
	}
}
