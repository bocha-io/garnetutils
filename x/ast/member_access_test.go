package ast

import "testing"

func TestMemberAccess(t *testing.T) {
	testData := `
{
  "argumentTypes": [
    {
      "typeIdentifier": "t_bytes32",
      "typeString": "bytes32"
    }
  ],
  "expression": {
    "id": 67842,
    "name": "Position",
    "nodeType": "Identifier",
    "overloadedDeclarations": [],
    "referencedDeclaration": 64643,
    "src": "2023:8:123",
    "typeDescriptions": {
      "typeIdentifier": "t_type$_t_contract$_Position_$64643_$",
      "typeString": "type(library Position)"
    }
  },
  "id": 67843,
  "isConstant": false,
  "isLValue": false,
  "isPure": false,
  "lValueRequested": false,
  "memberName": "get",
  "nodeType": "MemberAccess",
  "referencedDeclaration": 64339,
  "src": "2023:12:123",
  "typeDescriptions": {
    "typeIdentifier": "t_function_internal_view$_t_bytes32_$returns$_t_int32_$_t_int32_$",
    "typeString": "function (bytes32) view returns (int32,int32)"
  }
}
`

	expected := "Position.get"
	val, err := NewConverter().processMemberAccess([]byte(testData))
	if err != nil {
		t.Error(err)
	}
	if val != expected {
		t.Errorf("failed: %s, %s", val, expected)
	}
}
