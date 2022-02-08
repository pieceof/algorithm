package main

import (
	"algorithm/bitmap"
	"fmt"
)

func main(){
	bmap := bitmap.NewBitmap()
	bmap.AddStr("100")
	bmap.AddStr("120")
	bmap.AddStr("190")
	fmt.Println(bmap.CheckStr("120"))
	fmt.Println(bmap.Min(3))
	fmt.Println(bmap.Max(0))
}