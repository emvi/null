package null

import (
	"database/sql"
	"encoding/json"
)

// Nullable int64 type based on sql.NullInt64, that supports parsing to/from JSON.
type Int64 struct {
	sql.NullInt64
}

// Returns a new nullable Int64 object.
// This is equivalent to `null.Int64{sql.NullInt64{Int64: i, Valid: valid}}`.
func NewInt64(i int64, valid bool) Int64 {
	return Int64{sql.NullInt64{Int64: i, Valid: valid}}
}

func (this Int64) MarshalJSON() ([]byte, error) {
	if this.Valid {
		return json.Marshal(this.Int64)
	}

	return json.Marshal(nil)
}

func (this *Int64) UnmarshalJSON(data []byte) error {
	var value *int64

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if value != nil {
		this.Valid = true
		this.Int64 = *value
	} else {
		this.Valid = false
	}

	return nil
}

// Sets the value and valid to true.
func (this *Int64) SetValid(i int64) {
	this.Int64 = i
	this.Valid = true
}

// Sets the value to default and valid to false.
func (this *Int64) SetNil() {
	this.Int64 = 0
	this.Valid = false
}
