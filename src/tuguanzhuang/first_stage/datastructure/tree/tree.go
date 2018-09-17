/*
 * @Author: Mr Bian
 * @Date: 2018-08-26 18:15:08
 * @LastEditors: Mr Bian
 * @LastEditTime: 2018-08-26 18:22:55
 * @Description:
 * @version:
 */

package tree

import (
	"fmt"
	"math/rand"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(100) {
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

//  walks the tree t sending all values
// from the tree to the channel ch.
func WalkRecursion(t *Tree, ch chan int) {
	if t == nil {
		fmt.Println("This is an empty tree.")
		return
	}
	if t.Left != nil {
		WalkRecursion(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		WalkRecursion(t.Right, ch)
	}
}

// control the goroutine of walking, enabling closing the goroutine
func Walk(t *Tree, ch chan int) {
	WalkRecursion(t, ch)
	close(ch)
}

/**   test Walk *******************************
ch := make(chan int)
go tree.Walk(tree.New(1), ch)
for value := range ch {
	fmt.Printf("%v ", value)
}
*/

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for value := range ch1 {
		if value != <-ch2 {
			return false
		}
	}
	if _, ok := <-ch2; ok == true {
		return false
	}
	return true
}

func SameBuffered(t1, t2 *Tree, buffer_amount int) bool {
	ch1 := make(chan int, buffer_amount)
	ch2 := make(chan int, buffer_amount)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for value := range ch1 {
		if value != <-ch2 {
			return false
		}
	}
	if _, ok := <-ch2; ok == true {
		return false
	}
	return true
}

/**
// test the buffered and unbuffered Same() function
// advantage of multi-channel is not obvious
	t1 := tree.New(1)
	t2 := tree.New(2)
	for i := 1; i <= 10; i++ {
		starting_time := time.Now()
		tree.SameBuffered(t1, t2, i)
		time_consumption := time.Since(starting_time)
		fmt.Printf("%2v buffered channel use %v \n", i, time_consumption)
	}
*/
