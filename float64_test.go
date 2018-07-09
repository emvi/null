package null

import (
	"database/sql"
	"encoding/json"
	"testing"
)

type testFloat64 struct {
	Value Float64 `json:"value"`
}

func TestNewFloat64(t *testing.T) {
	value := NewFloat64(123.45, true)

	if value.Float64 != 123.45 || !value.Valid {
		t.Fatal("New Float64 must have value and be valid")
	}
}

func TestMarshalFloat64(t *testing.T) {
	value := Float64{sql.NullFloat64{Float64: 123.45, Valid: true}}

	if data, err := json.Marshal(value); err != nil || string(data) != "123.45" {
		t.Fatalf("Float64 must be marshalled to value, but was %v %v", err, string(data))
	}

	value.Valid = false

	if data, err := json.Marshal(value); err != nil || string(data) != "null" {
		t.Fatalf("Float64 must be marshalled to null, but was %v %v", err, string(data))
	}
}

func TestUnmarshalFloat64(t *testing.T) {
	str := `{"value": 123.45}`
	var value testFloat64

	if err := json.Unmarshal([]byte(str), &value); err != nil {
		t.Fatalf("Float64 must be unmarshalled to value, but was %v", err)
	}

	if !value.Value.Valid || value.Value.Float64 != 123.45 {
		t.Fatalf("Unmarshalled null float64 must be valid, but was %v", value.Value)
	}

	str = `{"value": null}`

	if err := json.Unmarshal([]byte(str), &value); err != nil {
		t.Fatalf("String must be unmarshalled to null, but was %v", err)
	}

	if value.Value.Valid {
		t.Fatal("Unmarshalled null float64 must be invalid")
	}
}

func TestGettersSettersFloat64(t *testing.T) {
	value := NewFloat64(123.45, true)
	value.SetNil()

	if value.Float64 != 0 || value.Valid {
		t.Fatal("Float64 must be nil")
	}

	value.SetValid(123.45)

	if value.Float64 != 123.45 || !value.Valid {
		t.Fatal("Float64 must be valid")
	}
}
