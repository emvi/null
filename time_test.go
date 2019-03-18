package null

import (
	"encoding/json"
	"testing"
	"time"
)

type testTime struct {
	Value Time `json:"value"`
}

func TestNewTime(t *testing.T) {
	value := NewTime(time.Now(), true)

	if value.Time.Equal(time.Time{}) || !value.Valid {
		t.Fatal("New Time must have value and be valid")
	}
}

func TestMarshalTime(t *testing.T) {
	now := time.Now()
	nowStrBytes, _ := json.Marshal(now)
	nowStr := string(nowStrBytes)
	value := Time{Time: now, Valid: true}

	if data, err := json.Marshal(value); err != nil || string(data) != nowStr {
		t.Fatalf("Time must be marshalled to value, but was %v %v", err, string(data))
	}

	value.Valid = false

	if data, err := json.Marshal(value); err != nil || string(data) != "null" {
		t.Fatalf("Time must be marshalled to null, but was %v %v", err, string(data))
	}
}

func TestUnmarshalTime(t *testing.T) {
	now := time.Now()
	nowStrBytes, _ := json.Marshal(now)
	nowStr := string(nowStrBytes)
	str := `{"value": `+nowStr+`}`
	var value testTime

	if err := json.Unmarshal([]byte(str), &value); err != nil {
		t.Fatalf("Time must be unmarshalled to value, but was %v", err)
	}

	if !value.Value.Valid || !value.Value.Time.Equal(now) {
		t.Fatalf("Unmarshalled null Time must be valid, but was %v", value.Value)
	}

	str = `{"value": null}`

	if err := json.Unmarshal([]byte(str), &value); err != nil {
		t.Fatalf("Time must be unmarshalled to null, but was %v", err)
	}

	if value.Value.Valid {
		t.Fatal("Unmarshalled null Time must be invalid")
	}
}

func TestGettersSettersTime(t *testing.T) {
	value := NewTime(time.Now(), true)
	value.SetNil()

	if !value.Time.Equal(time.Time{}) || value.Valid {
		t.Fatal("Time must be nil")
	}

	value.SetValid(time.Now())

	if value.Time.Equal(time.Time{}) || !value.Valid {
		t.Fatal("Time must be valid")
	}
}
