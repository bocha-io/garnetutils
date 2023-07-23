package converter

import "testing"

func TestCreateEventFunction(t *testing.T) {
	c := Converter{mainStruct: "GameObject"}
	generated := c.CreateEventFunction("Projectile", []Field{
		{Key: "spawned", Type: "bool"},
		{Key: "x", Type: "uint32"},
		{Key: "y", Type: "uint32"},
		{Key: "game", Type: "bytes32"},
	})
	value := `
func CreateProjectileEvent(ID string, spawned bool, x int64, y int64, game []byte) data.MudEvent {
    return data.MudEvent{
        Table: "Projectile",
        Key:   ID,
        Fields: []data.Field{
            {Key: "spawned", Data: data.BoolField{Data: spawned}},
            {Key: "x", Data: data.UintField{Data: *big.NewInt(x)}},
            {Key: "y", Data: data.UintField{Data: *big.NewInt(y)}},
            {Key: "game", Data: data.NewBytesField(game)},
        },
    }
}`
	if generated != value {
		t.Fatalf("error generating event function: %s, %s", generated, value)
	}
}
