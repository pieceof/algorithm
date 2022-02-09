package rbtree

import (
	"algorithm/item"
	"fmt"
)

const (
	RED        bool = true
	BLACK      bool = false
	LEFTCHILD  int  = -1
	RIGHTCHILD int  = 1
	NIL        int  = 0
)

var (
	midValue NodeSlice
	frontValue NodeSlice
	backValue NodeSlice
)

type RBTree struct {
	root  *Node
	count uint64
}
type NodeSlice []*Node

func NewRBTree() *RBTree {
	return &RBTree{
		root:  nil,
		count: 0,
	}
}

func leftRotate(n *Node) *Node {
	retNode := n.rightChild
	if n.whichChild() == LEFTCHILD{
		n.parent.leftChild = retNode
	}else if n.whichChild() == RIGHTCHILD{
		n.parent.rightChild = retNode
	}
	retNode.parent = n.parent
	n.rightChild = retNode.leftChild
	if n.rightChild != nil{
		n.rightChild.parent = n
	}

	retNode.leftChild = n
	n.parent = retNode

	return retNode
}
func rightRotate(n *Node) *Node {
	retNode := n.leftChild

	if n.whichChild() == LEFTCHILD{
		n.parent.leftChild = retNode
	}else if n.whichChild() == RIGHTCHILD{
		n.parent.rightChild = retNode
	}
	retNode.parent = n.parent

	n.leftChild = retNode.rightChild
	if n.leftChild != nil{
		n.leftChild.parent = n
	}

	retNode.rightChild = n
	n.parent = retNode

	return retNode
}

func check(val item.Item, node *Node) (bool, error) {
	if node == nil {
		return false, nil
	}
	equal, err := node.value.Equal(val)
	if err != nil {
		return false, err
	}
	if equal {
		return true, nil
	}
	less, err := node.value.Less(val)
	if err != nil {
		return false, err
	}
	if less {
		return check(val, node.leftChild)
	} else {
		return check(val, node.rightChild)
	}
}
func (rbt *RBTree) Find(val item.Item) (bool, error) {
	return check(val, rbt.root)
}

func insert(newNode *Node, node *Node) (*Node, error) {
	if node == nil {
		return newNode, nil
	}
	equal, err := node.value.Equal(newNode.value)
	if err != nil {
		return nil, err
	}
	if equal {
		return nil, fmt.Errorf("Node is exist")
	}
	less, err := newNode.value.Less(node.value)
	if err != nil {
		return nil, err
	}
	if less {
		tmpNode, err := insert(newNode, node.leftChild)
		if err != nil {
			return nil, err
		}
		node.leftChild = tmpNode
		tmpNode.parent = node
		return node, nil
	} else {
		tmpNode, err := insert(newNode, node.rightChild)
		if err != nil {
			return nil, err
		}
		node.rightChild = tmpNode
		tmpNode.parent = node
		return node, nil
	}
}
func insertCase1(node *Node) (root *Node) {
	if node.whichChild() == NIL {
		node.color = BLACK
		root = node
		return root
	}
	return insertCase2(node)
}
func insertCase2(node *Node) (root *Node) {
	if node.parent.color == BLACK {
		return
	}
	return insertCase3(node)
}
func insertCase3(node *Node) (root *Node) {
	if uncle := node.uncleNode(); uncle != nil && uncle.color == RED {
		grand := node.grandNode()
		node.parent.color = BLACK
		uncle.color = BLACK
		grand.color = RED
		return insertFixup(grand)
	}
	return insertCase4(node)
}
func insertCase4(node *Node) (root *Node) {
	if node.whichChild() == RIGHTCHILD && node.parent.whichChild() == LEFTCHILD {
		leftRotate(node.parent)
		node = node.leftChild
	} else if node.whichChild() == LEFTCHILD && node.parent.whichChild() == RIGHTCHILD {
		rightRotate(node.parent)
		node = node.rightChild
	}
	return insertCase5(node)
}
func insertCase5(node *Node) (root *Node) {
	grand := node.grandNode()
	node.parent.color = BLACK
	grand.color = RED
	if node.whichChild() == LEFTCHILD && node.parent.whichChild() == LEFTCHILD {
		root = rightRotate(grand)
	} else {
		root = leftRotate(grand)
	}
	return root
}
func insertFixup(node *Node) (root *Node) {
	return insertCase1(node)
}
func (rbt *RBTree) Insert(val item.Item) error {
	newNode := NewNode(val)
	tmpNode, err := insert(newNode, rbt.root)
	if err != nil {
		return err
	}
	rbt.count++
	rbt.root = tmpNode
	tNode := insertFixup(newNode)
	if rbt.root.whichChild() != NIL {
		rbt.root = tNode
	}
	return nil
}

func (rbt *RBTree) GetCount() uint64 {
	return rbt.count
}
func treeGetValue(node *Node) {
	if node == nil {
		return
	}
	frontValue = append(frontValue, node)
	treeGetValue(node.leftChild)
	midValue = append(midValue, node)
	treeGetValue(node.rightChild)
	backValue = append(backValue, node)
}
func (rbt *RBTree) PrintTree() {
	midValue = nil
	frontValue = nil
	backValue = nil
	treeGetValue(rbt.root)
	fmt.Printf("\nfront: \n")
	for _, v := range frontValue {
		v.PrintNode()
	}
	fmt.Printf("\nmid: \n")
	for _, v := range midValue {
		v.PrintNode()
	}
	fmt.Printf("\nback: \n")
	for _, v := range backValue {
		v.PrintNode()
	}
}