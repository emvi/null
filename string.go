package null

import (
	"database/sql"
	"encoding/json"
)

// Nullable string type based on sql.NullString, that supports parsing to/from JSON.
type String struct {
	sql.NullString
}

// Returns a new nullable String object.
// This is equivalent to `null.String{sql.NullString{String: s, Valid: valid}}`.
func NewString(s string, valid bool) String {
	return String{sql.NullString{String: s, Valid: valid}}
}

func (this String) MarshalJSON() ([]byte, error) {
	if this.Valid {
		return json.Marshal(this.String)
	}

	return json.Marshal(nil)
}

func (this *String) UnmarshalJSON(data []byte) error {
	var value *string

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if value != nil {
		this.Valid = true
		this.String = *value
	} else {
		this.Valid = false
	}

	return nil
}

// Sets the value and valid to true.
func (this *String) SetValid(s string) {
	this.String = s
	this.Valid = true
}

// Sets the value to default and valid to false.
func (this *String) SetNil() {
	this.String = ""
	this.Valid = false
}
