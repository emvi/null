package null

import (
	"database/sql"
	"encoding/json"
)

type String struct {
	sql.NullString
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
