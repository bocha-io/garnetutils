package converter

import (
	"testing"
)

func TestString(t *testing.T) {
	c := Converter{mainStruct: "GameState"}
	generated := c.SingleValueString("PlayerTwo")
	value := `func (g *GameState) getPlayerTwo(rowID string) (data.Field, string, error) {
	return data.GetRowFromIDUsingString(g.db, g.world, rowID, "PlayerTwo")
}`
	if generated != value {
		t.Fatalf("string getter failed: %s, %s", generated, value)
	}
}

func TestInt(t *testing.T) {
	c := Converter{mainStruct: "GameState"}
	generated := c.SingleValueInt("Time")
	value := `func (g *GameState) getTime(key string) (int64, error) {
	return data.GetInt64UsingString(g.db, g.world, key, "Time")
}`
	if generated != value {
		t.Fatalf("int getter failed: %s, %s", generated, value)
	}
}

func TestMultiValueTable(t *testing.T) {
	c := Converter{mainStruct: "GameState"}
	generated := c.MultiValueTable("Projectile", []Field{
		{Key: "test1", Type: "bool"},
		{Key: "test2", Type: "int32"},
		{Key: "test3", Type: "int32"},
		{Key: "test4", Type: "bytes32"},
	}, false)
	value := `
func (g *GameState) getProjectile(key string) (bool, int64, int64, string, error) {
    fields, err := data.GetRowFieldsUsingString(g.db, g.world, key, "Projectile")
    if err != nil {
        return false, 0, 0, "", err
    }

    if len(fields) != 4 {
        return false, 0, 0, "", fmt.Errorf("invalid amount of fields")
    }

    field0 := fields[0].Data.String() == "true"
    field1, err := strconv.ParseInt(fields[1].Data.String(), 10, 32)
    if err != nil {
        return false, 0, 0, "", err
    }
    field2, err := strconv.ParseInt(fields[2].Data.String(), 10, 32)
    if err != nil {
        return false, 0, 0, "", err
    }
    field3 := strings.ReplaceAll(fields[3].Data.String(), "\"", "")
    return field0, field1, field2, field3, nil
}`

	if generated != value {
		t.Fatalf("multivalue failed: %s, %s", generated, value)
	}
}
