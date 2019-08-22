package bst

import (
	"fmt"
	"sync"
)

type Item interface{}

type Node struct {
	key int
	value Item
	left *Node
	right *Node
	onLeftSide bool
	onRightSide bool
}

type BinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

func (bst *BinarySearchTree) Insert(key int, value Item) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	n := &Node{key, value, nil, nil, false, false}

	if bst.root == nil {
		bst.root = n
	} else {
		if bst.root.key < n.key {
			n.onLeftSide = true
		} else if bst.root.key > n.key {
			n.onRightSide = true
		}

		insertNode(bst.root, n)
	}
}

func insertNode(node, newNode *Node) {
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

func (bst *BinarySearchTree) InorderTraverse(f func(Item)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	inOrderTraverse(bst.root, f)
}

func inOrderTraverse(n *Node, f func(Item)) {
	if n != nil {
		inOrderTraverse(n.left, f)
		f(n.value)
		inOrderTraverse(n.right, f)
	}
}

func (bst *BinarySearchTree) PreOrderTraverse(f func(Item)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	preOrderTraverse(bst.root, f)
}

func preOrderTraverse(n *Node, f func(Item)) {
	if n != nil {
		f(n.value)
		preOrderTraverse(n.left, f)
		preOrderTraverse(n.right, f)
	}
}

func (bst *BinarySearchTree) PostOrderTraverse(f func(Item)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	postOrderTraverse(bst.root, f)
}

func postOrderTraverse(n *Node, f func(Item)) {
	if n != nil {
		postOrderTraverse(n.left, f)
		postOrderTraverse(n.right, f)
		f(n.value)
	}
}

func (bst *BinarySearchTree) Min() *Item {
	bst.lock.RLock()
	defer bst.lock.RUnlock()

	n := bst.root
	if n == nil {
		return nil
	}

	for {
		if n.left == nil {
			return &n.value
		}
		n = n.left
	}
}

func (bst *BinarySearchTree) Max() *Item {
	bst.lock.RLock()
	defer bst.lock.RUnlock()

	n := bst.root
	if n == nil {
		return nil
	}

	for {
		if n.right == nil {
			return &n.value
		}
		n = n.right
	}
}

func (bst *BinarySearchTree) Search(key int) bool {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	return search(bst.root, key)
}

func search(n* Node, key int) bool {
	if n == nil {
		return false
	}

	if key < n.key {
		return search(n.left, key)
	}

	if key > n.key {
		return search(n.right, key)
	}

	return true
}

func (bst *BinarySearchTree) Remove(key int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	remove(bst.root, key)
}

func remove(node *Node, key int) *Node {
	if node == nil {
		return nil
	}

	if key < node.key {
		node.left = remove(node.left, key)
		return node
	}

	if key > node.key {
		node.right = remove(node.right, key)
		return node
	}

	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}

	if node.left == nil {
		node = node.right
		return node
	}

	if node.right == nil {
		node = node.left
		return node
	}

	leftmostrightside := node.right
	for {
		//find smallest value on the right side
		if leftmostrightside != nil && leftmostrightside.left != nil {
			leftmostrightside = leftmostrightside.left
		} else {
			break
		}
	}

	node.key, node.value = leftmostrightside.key, leftmostrightside.value
	node.right = remove(node.right, node.key)

	return node
}

func (bst *BinarySearchTree) String() {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		/*if n.onLeftSide {
			color.HiRedString(format + "%d\n", n.key)
		} else if n.onRightSide {
			color.HiGreenString(format + "%d\n", n.key)
		} else {*/
			fmt.Printf(format + "%d\n", n.key)
		/*}*/
		stringify(n.right, level)
	}
}

func (bst* BinarySearchTree) PrintNodes() {
	bst.lock.RLock()
	defer bst.lock.RUnlock()

	printNodes(bst.root)

}

func printNodes(node *Node) {
	if node != nil {
		printNodes(node.left)
		fmt.Printf("Middle node: %v\n", node.value)
		if node.left != nil {
			fmt.Printf("\tLeft: %v\n", node.left.value)
		}
		if node.right != nil {
			fmt.Printf("\tRight: %v\n", node.right.value)
		}
		printNodes(node.right)
	}
}