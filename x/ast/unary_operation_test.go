package ast

import "testing"

func TestUnaryOperation(t *testing.T) {
	testData := `
{
  "id": 67925,
  "isConstant": false,
  "isLValue": false,
  "isPure": false,
  "lValueRequested": false,
  "nodeType": "UnaryOperation",
  "operator": "!",
  "prefix": true,
  "src": "2726:8:123",
  "subExpression": {
    "id": 67924,
    "name": "spawned",
    "nodeType": "Identifier",
    "overloadedDeclarations": [],
    "referencedDeclaration": 67912,
    "src": "2727:7:123",
    "typeDescriptions": {
      "typeIdentifier": "t_bool",
      "typeString": "bool"
    }
  },
  "typeDescriptions": {
    "typeIdentifier": "t_bool",
    "typeString": "bool"
  }
}
`

	expected := "!spawned"

	val, err := processUnaryOperation([]byte(testData))
	if err != nil {
		t.Error(err)
	}
	if val != expected {
		t.Errorf("failed: %s, %s", val, expected)
	}
}
