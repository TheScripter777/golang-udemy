package bst

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

type Attribute int

// Base attributes
const (
	Reset Attribute = iota
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

const (
	MaxTree = 20
)

func fillTree(bst *BinarySearchTree) {
	bst.Insert(8, "8")
	bst.Insert(4, "4")
	bst.Insert(10, "10")
	bst.Insert(2, "2")
	bst.Insert(6, "6")
	bst.Insert(1, "1")
	bst.Insert(3, "3")
	bst.Insert(5, "5")
	bst.Insert(7, "7")
	bst.Insert(9, "9")
}

func TestRandomInsert(t *testing.T) {
	bst := &BinarySearchTree{}
	bst.Insert(MaxTree / 2, strconv.Itoa(MaxTree / 2))

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < MaxTree; i++ {
		r := rand.Intn(MaxTree) + 1
		bst.Insert(r, fmt.Sprintf("%v", r))
	}

	bst.String()
	//bst.PrintNodes()
}

func TestInsert(t *testing.T) {
	bst := &BinarySearchTree{}
	fillTree(bst)

	//bst.String()

	bst.Insert(11, "11")
	bst.Insert(20, "20")
	bst.Insert(15, "15")
	bst.Insert(14, "14")
	bst.Insert(19, "19")
	bst.Insert(13, "13")
	bst.String()
	bst.PrintNodes()
}

func TestInOrderTraverse(t *testing.T) {
	bst := &BinarySearchTree{}
	fillTree(bst)
	var result []string

	bst.InorderTraverse(func(i Item) {
		result = append(result, fmt.Sprintf("%s", i))
	})

	if !isSameSlice(result, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}) {
		t.Errorf("Traversal order incorrect, got %v", result)
	}
}

func TestMin(t *testing.T) {
	bst := &BinarySearchTree{}
	fillTree(bst)
	if fmt.Sprintf("%s", *bst.Min()) != "1" {
		t.Errorf("min should be 1")
	}
}

func TestMax(t *testing.T) {
	bst := &BinarySearchTree{}
	fillTree(bst)
	if fmt.Sprintf("%s", *bst.Max()) != "10" {
		t.Errorf("max should be 10")
	}
}

func isSameSlice(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a{
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
