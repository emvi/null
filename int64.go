package null

import (
	"database/sql"
	"encoding/json"
)

type Int64 struct {
	sql.NullInt64
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
