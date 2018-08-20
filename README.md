# Money

[![Godoc](https://img.shields.io/badge/godoc-money-blue.svg)](https://godoc.org/github.com/mpwalkerdine/money)
[![Go Report Card](https://goreportcard.com/badge/github.com/mpwalkerdine/money)](https://goreportcard.com/report/github.com/mpwalkerdine/money)

```go
import "github.com/mpwalkerdine/money"
```

Package money is a convenience wrapper for [github.com/ericlagergren/decimal]().

It makes decimal values returned from this package immutable, at the expense of reduced memory efficiency.
See [https://golang.org/pkg/math/big/]() for why the API is designed in the way it is.
We forego that benefit in calling code, but where possible this package will attempt to minimise allocations during calculations.
In the majority of cases however, monetary values fit inside the compact uint64 value used by the underlying decimal package.