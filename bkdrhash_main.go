package main

import (
	"algorithm/bkdrhash"
	"fmt"
)

func main(){
	str := "fjdaljfjhdal/fdaj/fjdalk"
	hasher := bkdrhash.NewHasherBKDR()
	hasher.AddStr(str)
	hasher.AddInt(1234)
	fmt.Println(hasher.Val())
	hs := bkdrhash.NewHasherBKDR()
	hs.AddStr(str+"1234")
	fmt.Println(hs.Val())
}
