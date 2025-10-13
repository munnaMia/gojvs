# gojvs
### Simple JS feature implemention in Golang.

## How to use ?
---------------empty


## *Example*

```go
package main

import (
	"fmt"

	"github.com/munnaMia/gojvs"
)

func main() {
    slice := []int{3, 2, 32, 23, 423, 3, 42, 3, 243, 3}

	fmt.Println(gojvs.Map(slice, func(elem, idx int) int { return elem * 2 }))
    // Output : [6 4 64 46 846 6 84 6 486 6]
}

```



## *Available Methods*
- **Map**(*slice, callback()*)
	- Returns a new modify slice based on callback function.
	- Doesn't touch or modify orignal slice.	
	- Callback function takes one args and that is each elements from the slice. and return an value 
- **ForEach**(*slice, callback()*)
	- Doesn't touch or modify orignal slice.	
	- Callback function takes two args one is each elements from the slice and another one is the indexs of each elements. and return a new value 
	- callback function exicute for each elemens of the slice
- **Find**(*slice, callback()*)
	- Find Function takes an slice and a callback function and returns true & the first element in the provided slice that satisfies the provided testing function else return false & a zerovalue.


