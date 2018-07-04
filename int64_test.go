package null

import (
	"database/sql"
	"encoding/json"
	"testing"
)

type testInt64 struct {
	Value Int64 `json:"value"`
}

func TestMarshalInt64(t *testing.T) {
	value := Int64{sql.NullInt64{Int64: 123, Valid: true}}

	if data, err := json.Marshal(value); err != nil || string(data) != "123" {
		t.Fatalf("Int64 must be marshalled to value, but was %v %v", err, string(data))
	}

	value.Valid = false

	if data, err := json.Marshal(value); err != nil || string(data) != "null" {
		t.Fatalf("Int64 must be marshalled to null, but was %v %v", err, string(data))
	}
}

func TestUnmarshalInt64(t *testing.T) {
	str := `{"value": 123}`
	var value testInt64

	if err := json.Unmarshal([]byte(str), &value); err != nil {
		t.Fatalf("Int64 must be unmarshalled to value, but was %v", err)
	}

	if !value.Value.Valid || value.Value.Int64 != 123 {
		t.Fatalf("Unmarshalled null int64 must be valid, but was %v", value.Value)
	}

	str = `{"value": null}`

	if err := json.Unmarshal([]byte(str), &value); err != nil {
		t.Fatalf("Int64 must be unmarshalled to null, but was %v", err)
	}

	if value.Value.Valid {
		t.Fatal("Unmarshalled null int64 must be invalid")
	}
}
