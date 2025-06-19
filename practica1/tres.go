package main

import "fmt"

func main() {
	/* integers */
	var zz int = 0 /* o para hexadecimal 0xA */
	x := 10
	var z int = x /* x no esta declarada */
	var y int8 = int8(x + 1)
	const n = 5001
	const c int = 5001
	/* float */
	var e float32 = 6
	var f float32 = e

	fmt.Println(zz, z, y, c, n, e, f)
}
