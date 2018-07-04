package null

import (
	"database/sql"
	"encoding/json"
	"testing"
)

type testBool struct {
	Value Bool `json:"value"`
}

func TestMarshalBool(t *testing.T) {
	value := Bool{sql.NullBool{Bool: true, Valid: true}}

	if data, err := json.Marshal(value); err != nil || string(data) != "true" {
		t.Fatalf("Bool must be marshalled to value, but was %v %v", err, string(data))
	}

	value.Valid = false

	if data, err := json.Marshal(value); err != nil || string(data) != "null" {
		t.Fatalf("Bool must be marshalled to null, but was %v %v", err, string(data))
	}
}

func TestUnmarshalBool(t *testing.T) {
	str := `{"value": true}`
	var value testBool

	if err := json.Unmarshal([]byte(str), &value); err != nil {
		t.Fatalf("Bool must be unmarshalled to value, but was %v", err)
	}

	if !value.Value.Valid || !value.Value.Bool {
		t.Fatalf("Unmarshalled null bool must be valid, but was %v", value.Value)
	}

	str = `{"value": null}`

	if err := json.Unmarshal([]byte(str), &value); err != nil {
		t.Fatalf("Bool must be unmarshalled to null, but was %v", err)
	}

	if value.Value.Valid {
		t.Fatal("Unmarshalled null bool must be invalid")
	}
}
