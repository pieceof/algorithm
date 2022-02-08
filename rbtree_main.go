package main

import (
	"algorithm/item"
	"algorithm/rbtree"
	"fmt"
)

func main(){
	rbt := rbtree.NewRBTree()
	rbt.Insert(item.Int(10))
	rbt.Insert(item.Int(9))
	rbt.Insert(item.Int(8))
	rbt.Insert(item.Int(6))
	rbt.Insert(item.Int(7))
	fmt.Println("count: ",rbt.GetCount())
	rbt.PrintTree()
}