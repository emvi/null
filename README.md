<p align="center">
    <img src="nullgopher.svg" width="300px" />
</p>

# Nullable Go types

[![GoDoc](https://godoc.org/github.com/emvicom/null?status.svg)](https://godoc.org/github.com/emvicom/null)
[![CircleCI](https://circleci.com/gh/emvicom/null.svg?style=svg)](https://circleci.com/gh/emvicom/null)
[![Go Report Card](https://goreportcard.com/badge/github.com/emvicom/null)](https://goreportcard.com/report/github.com/emvicom/null)

## Description

This package provides nullable Go types that replace sql.NullString, sql.NullInt64, ... that can be marshalled/unmarshalled to/from JSON.

## Installation

To install "null", run go get within your project:

```
go get github.com/emvicom/null
```

## Usage

Here is a short example demonstrating the string type. The other types (int64, float64 and bool) work in the same way.

```
package main

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
    str := null.NewString("nullable string", true)
    // or long version: str := NullableString{null.String{sql.NullString{String: "nullable string", Valid: true}}}
    data, _ := json.Marshal(str)
    fmt.Println(string(data)) // -> {"value": "nullable"}

    str.SetNil() // use str.SetValid("value") to set a value again
    data, _ = json.Marshal(str)
    fmt.Println(string(data)) // -> {"value": null}
}
```

## Contribute

[See CONTRIBUTING.md](CONTRIBUTING.md)

## License

MIT
