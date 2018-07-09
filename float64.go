package null

import (
	"database/sql"
	"encoding/json"
)

// Nullable float64 type based on sql.NullFloat64, that supports parsing to/from JSON.
type Float64 struct {
	sql.NullFloat64
}

// Returns a new nullable Float64 object.
// This is equivalent to `null.Float64{sql.NullFloat64{Float64: f, Valid: valid}}`.
func NewFloat64(f float64, valid bool) Float64 {
	return Float64{sql.NullFloat64{Float64: f, Valid: valid}}
}

func (this Float64) MarshalJSON() ([]byte, error) {
	if this.Valid {
		return json.Marshal(this.Float64)
	}

	return json.Marshal(nil)
}

func (this *Float64) UnmarshalJSON(data []byte) error {
	var value *float64

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if value != nil {
		this.Valid = true
		this.Float64 = *value
	} else {
		this.Valid = false
	}

	return nil
}

// Sets the value and valid to true.
func (this *Float64) SetValid(f float64) {
	this.Float64 = f
	this.Valid = true
}

// Sets the value to default and valid to false.
func (this *Float64) SetNil() {
	this.Float64 = 0
	this.Valid = false
}
