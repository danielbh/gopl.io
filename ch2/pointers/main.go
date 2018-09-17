package pointers

import (
	"fmt"
)

func main() {
	x := 1
	p := &x         // p, of type *x, points to x
	fmt.Println(*p) // "1"
	*p = 2          // equivlent to x = 2
	fmt.Println(x)  // 2
}

/**

func incr(p * int) int {
	*p++
	return *p
}

v := 1
incr(&v) // side effect: v is now 2
fmt.Printlin(incr(&v)) // "3" (and v is 3)

**/
