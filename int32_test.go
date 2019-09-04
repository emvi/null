package null

import (
	"database/sql"
	"encoding/json"
	"testing"
)

type testInt32 struct {
	Value Int32 `json:"value"`
}

func TestNewInt32(t *testing.T) {
	value := NewInt32(123, true)

	if value.Int32 != 123 || !value.Valid {
		t.Fatal("New Int32 must have value and be valid")
	}
}

func TestMarshalInt32(t *testing.T) {
	value := Int32{sql.NullInt32{Int32: 123, Valid: true}}

	if data, err := json.Marshal(value); err != nil || string(data) != "123" {
		t.Fatalf("Int32 must be marshalled to value, but was %v %v", err, string(data))
	}

	value.Valid = false

	if data, err := json.Marshal(value); err != nil || string(data) != "null" {
		t.Fatalf("Int32 must be marshalled to null, but was %v %v", err, string(data))
	}
}

func TestUnmarshalInt32(t *testing.T) {
	str := `{"value": 123}`
	var value testInt32

	if err := json.Unmarshal([]byte(str), &value); err != nil {
		t.Fatalf("Int32 must be unmarshalled to value, but was %v", err)
	}

	if !value.Value.Valid || value.Value.Int32 != 123 {
		t.Fatalf("Unmarshalled null Int32 must be valid, but was %v", value.Value)
	}

	str = `{"value": null}`

	if err := json.Unmarshal([]byte(str), &value); err != nil {
		t.Fatalf("Int32 must be unmarshalled to null, but was %v", err)
	}

	if value.Value.Valid {
		t.Fatal("Unmarshalled null Int32 must be invalid")
	}
}

func TestGettersSettersInt32(t *testing.T) {
	value := NewInt32(123, true)
	value.SetNil()

	if value.Int32 != 0 || value.Valid {
		t.Fatal("Int32 must be nil")
	}

	value.SetValid(123)

	if value.Int32 != 123 || !value.Valid {
		t.Fatal("Int32 must be valid")
	}
}
