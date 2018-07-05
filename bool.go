package null

import (
	"database/sql"
	"encoding/json"
)

// Nullable boolean type based on sql.NullBool, that supports parsing to/from JSON.
type Bool struct {
	sql.NullBool
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
