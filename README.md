# Sringf [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![npm version](https://badge.fury.io/js/stringf.svg)](https://badge.fury.io/js/stringf)
Dead simple cross-language string formatter

No dependencies, O(N*m)

*Note: Key must only contains alphabet and `_`. Invalid keys will be ignore*
# Usage
## go
```go
import (
	"fmt"
	"github.com/thanhpk/stringf"
)

func main() {
	str := stringf.Format("hi {name}, here is your number {{2108 and {name2}.", map[string]string{
		"name": "Kieu Thanh", // key should not contains spaces
	})
	fmt.Println(str)
}

// hi Kieu Thanh, here is your number {2108 and {name2}.
```

## js
```js
var stringf = require("stringf")

var str = stringf.Format("hi {name}, here is your number {{2108 and {name2}.", {
	name: "Thanh",
})
console.log(str)

// hi Thanh, here is your number {2108 and {name2}.
```

# Test
## Go
```sh
go test
```

## Js
```sh
npm test
```

# Testcase
| String    | Parameter map | Output    |
|-----------|---------------|-----------|
| hi {name}  | name: `Thanh` | hi Thanh  |
| hi {num}   | num: `2108`   | hi 2108   |
| {{abc}     | abc: `bcd`    | {abc      |

# Pseudocode to implement in your own language

### input:
	a string, s
	a map string to string, paramMap
### output:
  a string with param replaced

```
let ESCCHAR ← '{', ESCCHAREND ← '}', i ← 0, output ← ""
while i < length(s) do
	if s[i] = ESCCHAR then
		let j ← i + 1
		while j <= length(s) and s[j] = ESCCHAR do
			j ← j + 1
			if (j - i) mod 2 = 0 then
				output ← output + ESCCHAR
		if (j - i) mod 2 ≠ 0 then
			let param ← "" (parse parameter key)
			while j < length(s) and s[j] ≠ ESCCHAREND do
				param ← param + s[j]
				j ← j + 1
			if length(param) > 0 then
				output ← output + paramMap[param]
			if j = length(s) then
				return output
			else if s[j = ESCCHAREND then
				j ← j + 1
		i ← j
	output ← output + s[i]
	i ← i + 1
return output
```
