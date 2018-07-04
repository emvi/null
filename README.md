# Nullable Go types

This package provides nullable Go types that replace sql.NullString, sql.NullInt64, ... that can be marshalled/unmarshalled to/from JSON.

## Usage

```
import (
    "encoding/json"
    "database/sql"
    "fmt"

    "github.com/emvicom/null"
)

type NullableString struct {
    Value null.String `json:"value"`
}

func main() {
    str := NullableString{null.String{sql.NullString{String: "nullable string", Valid: true}}}
    data, _ := json.Marshal(str)
    fmt.Println(string(data)) // -> {"value": "nullable"}
}
```

## License

MIT
