# Sringf
Dead simple cross-language string formatter

No dependencies, O(N*m)

# Usage
## go
```go
import (
	"fmt"
	"github.com/thanhpk/stringf"
)

func main() {
	str := stringf.Format("hi #name, here is your number ##2108.", map[string]string{
		"name": "Kieu Thanh", // key should not contains spaces
	})
	fmt.Println(str)
}

// hi Kieu Thanh, here is your number #2108.
```

## js
```js
var stringf = require("stringf")

var str = stringf.Format("hi #name, here is your number ##2108.", {
	name: "Thanh",
})
console.log(str)

// hi Kieu Thanh, here is your number #2108.
```

# Pseudocode to implement in your own language
```
input:
	a string, s
	a map string to string, paramMap
output:
  a string with param replaced

let ESCCHAR ← '#', i ← 0, output ← ""

while i < length(s) do
	if s[i] = ESCCHAR then
		let j ← i + 1
		while j <= length(s) and s[j] = ESCCHAR do
			j ← j + 1
			if (j - i) mod 2 = 0 then
				output ← output + ESCCHAR
		if (j - i) mod 2 ≠ 0 then
			let param ← ""
			while j < length(s) and s[j] ≠ ' ' do
				param ← param + s[j]
				j ← j + 1
			if length(param) > 0 then
				output ← output + paramMap[param]
			if j = length(s) then
				return output
		i ← j
	output ← output + s[i]
	i ← i + 1
return output
```
