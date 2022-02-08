package main

import (
	"algorithm/item"
	"fmt"
)

func main() {
	var itype1 item.Int = 10
	var itype2 item.Int = 12

	fmt.Println(itype1.Less(itype2))

	strtype1 := item.NewString("sola")
	strtype2 := item.NewString("so")

	fmt.Println(strtype1.Equal(strtype2))
}