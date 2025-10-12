# gojvs
### Simple JS feature implemention in Golang.

## How to use ?
---------------empty

## *Available Methods*
- **Map**(*slice, callback()*)




## *Example*

```go
package main

import (
	"fmt"

	"github.com/munnaMia/gojvs"
)

func main() {
    slice := []int{3, 2, 32, 23, 423, 3, 42, 3, 243, 3}

	fmt.Println(gojvs.Map(slice, func(a int) int { return a * 2 }))
    // Output : [6 4 64 46 846 6 84 6 486 6]
}

```

