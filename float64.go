package null

import (
	"database/sql"
	"encoding/json"
)

type Float64 struct {
	sql.NullFloat64
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
