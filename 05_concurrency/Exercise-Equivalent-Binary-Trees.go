package main

import (
	"fmt"
	"math/rand"
)

func Walk(t *Tree, ch chan int) {
	WalkRecursive(t, ch)
	close(ch)
}

func WalkRecursive(t *Tree, ch chan int) {
	if t == nil {
		return
	}

	if t.Left != nil {
		WalkRecursive(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		WalkRecursive(t.Right, ch)
	}
}

func Same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v1 := range ch1 {
		v2 := <-ch2
		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(New(1), ch)
	fmt.Println(New(1))
	fmt.Println(New(1))
	fmt.Println(Same(New(1), New(1)))
	fmt.Println(Same(New(1), New(2)))
	fmt.Println(Same(New(2), New(1)))
}

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}
