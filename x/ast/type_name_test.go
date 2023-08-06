package ast

import "testing"

func TestTypeName_elementary(t *testing.T) {
	testData := `{
    "id": 82793,
    "name": "bytes32",
    "nodeType": "ElementaryTypeName",
    "src": "1606:7:152",
    "typeDescriptions": {
      "typeIdentifier": "t_bytes32",
      "typeString": "bytes32"
    }
  }
`

	expected := "bytes32"

	val, err := NewConverter().processElementaryTypeName([]byte(testData))
	if err != nil {
		t.Error(err)
	}
	if val != expected {
		t.Errorf("failed: %s, %s", val, expected)
	}
}

func TestTypeName_userDefined(t *testing.T) {
	testData := `{
    "id": 82853,
    "nodeType": "UserDefinedTypeName",
    "pathNode": {
      "id": 82852,
      "name": "PlacedCardsData",
      "nodeType": "IdentifierPath",
      "referencedDeclaration": 75931,
      "src": "2152:15:152"
    },
    "referencedDeclaration": 75931,
    "src": "2152:15:152",
    "typeDescriptions": {
      "typeIdentifier": "t_struct$_PlacedCardsData_$75931_storage_ptr",
      "typeString": "struct PlacedCardsData"
    }
  }`

	expected := "PlacedCardsData"

	val, err := NewConverter().processUserDefinedTypeName([]byte(testData))
	if err != nil {
		t.Error(err)
	}
	if val != expected {
		t.Errorf("failed: %s, %s", val, expected)
	}
}

func TestTypeName_arrayTypeName(t *testing.T) {
	testData := `{
    "baseType": {
      "id": 83021,
      "name": "bytes32",
      "nodeType": "ElementaryTypeName",
      "src": "3496:7:152",
      "typeDescriptions": {
        "typeIdentifier": "t_bytes32",
        "typeString": "bytes32"
      }
    },
    "id": 83022,
    "nodeType": "ArrayTypeName",
    "src": "3496:9:152",
    "typeDescriptions": {
      "typeIdentifier": "t_array$_t_bytes32_$dyn_storage_ptr",
      "typeString": "bytes32[]"
    }
  }
`

	expected := "[]bytes32"

	val, err := NewConverter().processArrayTypeName([]byte(testData))
	if err != nil {
		t.Error(err)
	}
	if val != expected {
		t.Errorf("failed: %s, %s", val, expected)
	}
}
