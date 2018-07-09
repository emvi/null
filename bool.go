package null

import (
	"database/sql"
	"encoding/json"
)

// Nullable boolean type based on sql.NullBool, that supports parsing to/from JSON.
type Bool struct {
	sql.NullBool
}

// Returns a new nullable Bool object.
// This is equivalent to `null.Bool{sql.NullBool{Bool: b, Valid: valid}}`.
func NewBool(b, valid bool) Bool {
	return Bool{sql.NullBool{Bool: b, Valid: valid}}
}

func (this Bool) MarshalJSON() ([]byte, error) {
	if this.Valid {
		return json.Marshal(this.Bool)
	}

	return json.Marshal(nil)
}

func (this *Bool) UnmarshalJSON(data []byte) error {
	var value *bool

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if value != nil {
		this.Valid = true
		this.Bool = *value
	} else {
		this.Valid = false
	}

	return nil
}

// Sets the value and valid to true.
func (this *Bool) SetValid(b bool) {
	this.Bool = b
	this.Valid = true
}

// Sets the value to default and valid to false.
func (this *Bool) SetNil() {
	this.Bool = false
	this.Valid = false
}
