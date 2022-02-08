package rbtree

import (
	"algorithm/item"
	"fmt"
)

const (
	RED   bool = true
	BLACK bool = false
	LEFTCHILD int = -1
	RIGHTCHILD int = 1
	NIL int = 0
)

type Node struct {
	leftChild  *Node
	rightChild *Node
	parent     *Node
	value      item.Item
	color      bool
}
type RBTree struct {
	root  *Node
	count uint64
}
type NodeSlice []*Node
func NewNode(val item.Item) *Node {
	return &Node{
		leftChild:  nil,
		rightChild: nil,
		parent:     nil,
		value:      val,
		color:      RED,
	}
}
func NewRBTree() *RBTree{
	return &RBTree{
		root: nil,
		count: 0,
	}
}
func (n *Node) whichChild() int {
	if n.parent == nil{
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
func (n *Node) GetVal() item.Item {
	return n.value
}
func (n *Node) GetColor() string {
	if n.color == RED{
		return "RED"
	}else{
		return "BLACK"
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

func leftRotate(n *Node) *Node{
	retNode := n.rightChild
	n.rightChild = retNode.leftChild
	retNode.leftChild = n
	return retNode
}
func rightRotate(n *Node) *Node{
	retNode := n.leftChild
	n.leftChild = retNode.rightChild
	retNode.rightChild = n
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
	less, err := node.value.Less(newNode.value)
	if err != nil {
		return nil, err
	}
	if less {
		tmpNode, err := insert(newNode, node.leftChild)
		if err != nil {
			return nil, err
		}
		node.leftChild = tmpNode
		return node, nil
	} else {
		tmpNode, err := insert(newNode, node.rightChild)
		if err != nil {
			return nil, err
		}
		node.rightChild = tmpNode
		return node, nil
	}
}
func insertCase1(node *Node) (root *Node){
	if node.parent == nil{
		node.color = BLACK
		root = node
		return root
	}
	return insertCase2(node)
}
func insertCase2(node *Node) (root *Node){
	if node.parent.color == BLACK{
		return
	}
	return insertCase3(node)
}
func insertCase3(node *Node) (root *Node){
	if uncle := node.uncleNode();uncle != nil && uncle.color == RED{
		grand := node.grandNode()
		node.parent.color = BLACK
		uncle.color = BLACK
		grand.color = RED
		return insertFixup(grand)
	}
	return  insertCase4(node)
}
func insertCase4(node *Node) (root *Node){
	if node.whichChild() == RIGHTCHILD && node.parent.whichChild() == LEFTCHILD{
		leftRotate(node.parent)
		node = node.leftChild
	}else if node.whichChild() == LEFTCHILD && node.parent.whichChild() == RIGHTCHILD{
		rightRotate(node.parent)
		node = node.rightChild
	}
	return insertCase5(node)
}
func insertCase5(node *Node) (root *Node){
	grand := node.grandNode()
	node.parent.color = BLACK
	grand.color = RED
	if node.whichChild() == LEFTCHILD && node.parent.whichChild() == LEFTCHILD{
		root = rightRotate(grand)
	}else{
		root = leftRotate(grand)
	}
	return root
}
func insertFixup(node *Node) (root *Node){
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
	tmpNode = insertFixup(newNode)
	if rbt.root.parent != nil {
		rbt.root = tmpNode
	}
	return nil
}

var midValue NodeSlice
var frontValue NodeSlice
var backValue NodeSlice
func treeGetValue(node *Node){
	if node == nil{
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
	fmt.Printf("front: \n")
	for _,v := range frontValue {fmt.Printf("value: %v, color: %v | ",v.GetVal(),v.GetColor())}
	fmt.Printf("\nmid: \n")
	for _,v := range midValue {fmt.Printf("value: %v, color: %v | ",v.GetVal(),v.GetColor())}
	fmt.Printf("\nback: \n")
	for _,v := range backValue {fmt.Printf("value: %v, color: %v | ",v.GetVal(),v.GetColor())}
}

func (rbt *RBTree) GetCount() uint64{
	return rbt.count
}