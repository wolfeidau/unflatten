# unflatten 

This library can "flatten" and "unflatten" a hierarchy stored in a map[string]interface{}. 

[![GitHub Actions status](https://github.com/wolfeidau/unflatten/workflows/Go/badge.svg?branch=master)](https://github.com/wolfeidau/unflatten/actions?query=workflow%3AGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/wolfeidau/unflatten)](https://goreportcard.com/report/github.com/wolfeidau/unflatten)
[![Documentation](https://godoc.org/github.com/wolfeidau/unflatten?status.svg)](https://godoc.org/github.com/wolfeidau/unflatten) [![Coverage Status](https://coveralls.io/repos/github/wolfeidau/unflatten/badge.svg?branch=master)](https://coveralls.io/github/wolfeidau/unflatten?branch=master)

# usage

```go
var m = map[string]interface{}{
	"cpu.usage.0.user": map[string]interface{}{
		"value": 2.3,
	},
	"cpu.usage.0.system": map[string]interface{}{
		"value": 1.2,
	},
}

tree := Unflatten(m, func(k string) []string { return strings.Split(k, ".") })

```

# contributions

Thanks to [Andrew Leap](https://github.com/andyleap) for rewriting this library and reminding me I need to use functions more in golang.

# License

This code is Copyright (c) 2014 Mark Wolfe and licenced under the MIT licence. All rights not explicitly granted in the MIT license are reserved. See the included LICENSE.md file for more details.
