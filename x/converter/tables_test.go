package converter

import (
	_ "embed"
	"testing"
)

func TestTablesGeneration(t *testing.T) {
	tables := GetTablesFromJSON(mudConfigJSON)
	tableMatch := Table{
		Key: "Match",
		Values: []Field{
			{
				Key:  "created",
				Type: "bool",
			},
			{
				Key:  "gameType",
				Type: "GameType",
			},
		},
	}

	tablePlayerTwo := Table{
		Key: "PlayerTwo",
		Values: []Field{
			{
				Key:  "value",
				Type: "bytes32",
			},
		},
	}

	tableUser := Table{
		Key: "User",
		Values: []Field{
			{
				Key:  "value",
				Type: "bytes32",
			},
		},
	}

	tableCurrentGameState := Table{
		Key: "CurrentGameState",
		Values: []Field{
			{
				Key:  "value",
				Type: "GameState",
			},
		},
	}

	tableCurrentHp := Table{
		Key: "CurrentHp",
		Values: []Field{
			{
				Key:  "value",
				Type: "uint32",
			},
		},
	}

	tableInmune := Table{
		Key: "Inmune",
		Values: []Field{
			{
				Key:  "value",
				Type: "uint32",
			},
		},
	}

	tableTime := Table{
		Key: "Time",
		Values: []Field{
			{
				Key:  "value",
				Type: "uint32",
			},
		},
	}

	tableProjectil := Table{
		Key: "Projectil",
		Values: []Field{
			{
				Key:  "spawned",
				Type: "bool",
			},
			{
				Key:  "x",
				Type: "int32",
			},

			{
				Key:  "y",
				Type: "int32",
			},

			{
				Key:  "shotDir",
				Type: "int32",
			},
		},
	}

	tablePosition := Table{
		Key: "Position",
		Values: []Field{
			{
				Key:  "x",
				Type: "int32",
			},

			{
				Key:  "y",
				Type: "int32",
			},
		},
	}

	tablesRes := []Table{
		tableMatch,
		tablePlayerTwo,
		tableUser,
		tableCurrentGameState,
		tableCurrentHp,
		tableInmune,
		tableTime,
		tableProjectil,
		tablePosition,
	}
	if len(tablesRes) != len(tables) {
		t.Fatalf("incorrect table len")
	}
	for k := range tables {
		if tables[k].Key != tablesRes[k].Key {
			t.Fatalf("table keys are different")
		}
		for i := range tables[k].Values {
			if tables[k].Values[i] != tablesRes[k].Values[i] {
				t.Fatalf("table values are different")
			}
		}
	}
}

func TestEnmusGeneration(t *testing.T) {
	enums := GetEnumsFromJSON(mudConfigJSON)

	enumsRes := []Enum{
		{
			Key:    "GameState",
			Values: []string{"Playing", "Victory", "Defeat"},
		},
		{
			Key:    "GameType",
			Values: []string{"Solo", "Online"},
		},
	}

	if len(enums) != len(enumsRes) {
		t.Fatalf("incorrect enum len")
	}

	for k := range enums {
		if enums[k].Key != enumsRes[k].Key {
			t.Fatalf("enums keys are different")
		}
		for i := range enums[k].Values {
			if enums[k].Values[i] != enumsRes[k].Values[i] {
				t.Fatalf("enums values are different")
			}
		}
	}
}
