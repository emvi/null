<p align="center">
    <img src="nullgopher.svg" width="300px" />
</p>

# Nullable Go types

[![GoDoc](https://godoc.org/github.com/emvi/null?status.svg)](https://godoc.org/github.com/emvi/null)
[![CircleCI](https://circleci.com/gh/emvi/null.svg?style=svg)](https://circleci.com/gh/emvi/null)
[![Go Report Card](https://goreportcard.com/badge/github.com/emvi/null)](https://goreportcard.com/report/github.com/emvi/null)
<a href="https://discord.gg/5UWdRh"><img src="https://img.shields.io/discord/739184135649886288?logo=discord" alt="Chat on Discord"></a>

## Description

This package provides nullable Go types for bool, float64, int64, int32, string and time.Time replacing sql.NullString, sql.NullInt64, ... that can be marshalled/unmarshalled to/from JSON.

## Installation

To install "null", run go get within your project:

```
go get github.com/emvi/null
```

Note that from 1.3 on "null" requires Go version 1.13 or newer.

## Usage

Here is a short example demonstrating the string type. The other types (int64, float64 and bool) work in the same manner.

```
package main

import (
    "encoding/json"
    "database/sql"
    "fmt"

    "github.com/emvi/null"
)

type NullableString struct {
    Value null.String `json:"value"`
}

func main() {
    str := NullableString{null.NewString("nullable string", true)}
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
