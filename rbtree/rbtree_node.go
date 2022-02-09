package rbtree

import (
	"algorithm/item"
	"fmt"
)

type Node struct {
	leftChild  *Node
	rightChild *Node
	parent     *Node
	value      item.Item
	color      bool
}

func NewNode(val item.Item) *Node {
	return &Node{
		leftChild:  nil,
		rightChild: nil,
		parent:     nil,
		value:      val,
		color:      RED,
	}
}

func (n *Node) grandNode() *Node {
	if n.parent == nil {
		return nil
	}
	return n.parent.parent
}
func (n *Node) uncleNode() *Node {
	if n.grandNode() == nil {
		return nil
	}
	if n.parent.whichChild() == RIGHTCHILD {
		return n.grandNode().leftChild
	}
	return n.grandNode().rightChild
}
func (n *Node) whichChild() int {
	if n.parent == nil {
		return NIL
	}
	if n == n.parent.leftChild {
		return LEFTCHILD
	}
	if n == n.parent.rightChild {
		return RIGHTCHILD
	}
	return NIL
}
func (n *Node) GetVal() string {
	return n.value.GetValue()
}
func (n *Node) GetColor() string {
	if n.color == RED {
		return "RED"
	} else {
		return "BLACK"
	}
}
func (n *Node) PrintNode() {
	fmt.Printf("value: %v, color: %v | ", n.GetVal(), n.GetColor())
}