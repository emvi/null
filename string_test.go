package null

import (
	"database/sql"
	"encoding/json"
	"testing"
)

type testString struct {
	Value String `json:"value"`
}

func TestMarshalString(t *testing.T) {
	value := String{sql.NullString{String: "test", Valid: true}}

	if data, err := json.Marshal(value); err != nil || string(data) != "\"test\"" {
		t.Fatalf("String must be marshalled to value, but was %v %v", err, string(data))
	}

	value.Valid = false

	if data, err := json.Marshal(value); err != nil || string(data) != "null" {
		t.Fatalf("String must be marshalled to null, but was %v %v", err, string(data))
	}
}

func TestUnmarshalString(t *testing.T) {
	str := `{"value": "test"}`
	var value testString

	if err := json.Unmarshal([]byte(str), &value); err != nil {
		t.Fatalf("String must be unmarshalled to value, but was %v", err)
	}

	if !value.Value.Valid || value.Value.String != "test" {
		t.Fatalf("Unmarshalled null string must be valid, but was %v", value.Value)
	}

	str = `{"value": null}`

	if err := json.Unmarshal([]byte(str), &value); err != nil {
		t.Fatalf("String must be unmarshalled to null, but was %v", err)
	}

	if value.Value.Valid {
		t.Fatal("Unmarshalled null string must be invalid")
	}
}
