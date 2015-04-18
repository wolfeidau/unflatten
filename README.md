# unflatten [![GoDoc](https://img.shields.io/badge/godoc-Reference-brightgreen.svg?style=flat)](http://godoc.org/github.com/wolfeidau/unflatten)

This library a simple way to "unflatten" a map[string]interface{} where the keys represent some flattened structure. 

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

# License

This code is Copyright (c) 2014 Mark Wolfe and licenced under the MIT licence. All rights not explicitly granted in the MIT license are reserved. See the included LICENSE.md file for more details.