package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	tree := create(10)

	fmt.Println("The value of the root of the tree is", tree.Value)

	traverse(tree)
	fmt.Println()

	// tree = insert(tree, -10)
	// tree = insert(tree, -2)

	// traverse(tree)

	fmt.Println(tree)
}

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Printf("Left: %v; Value: %v; Right: %v;\n", t.Left, t.Value, t.Right)
	traverse(t.Right)
}

func create(n int) *Tree {
	var t *Tree

	rand.Seed(time.Now().Unix())

	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
		t = insert(t, temp)
	}

	return t
}

func insert(t *Tree, v int) *Tree {
	// it it is empty tree, v is its root
	if t == nil {
		return &Tree{nil, v, nil}
	}

	// if v already exists in tree, return input tree with NOOP
	if v == t.Value {
		return t
	}

	// input val goes to left (is lower than node value)?
	if v < t.Value {
		// recursively call insert() with left tree as input
		// and assign its result to left tree
		t.Left = insert(t.Left, v)

		//return the resulting tree
		return t
	}

	// recursively call insert() with right tree as input
	// and assign its result to right tree
	t.Right = insert(t.Right, v)

	//return resulting tree
	return t
}


// The value of the root of the tree is 15
// Left: <nil>; Value: 0; Right: <nil>;
// Left: &{<nil> 0 <nil>}; Value: 2; Right: &{<nil> 3 0xc000070160};
// Left: <nil>; Value: 3; Right: &{<nil> 4 <nil>};
// Left: <nil>; Value: 4; Right: <nil>;
// Left: &{0xc0000701a0 2 0xc000070140}; Value: 6; Right: &{0xc000070180 9 0xc0000701c0};
// Left: <nil>; Value: 8; Right: <nil>;
// Left: &{<nil> 8 <nil>}; Value: 9; Right: &{<nil> 12 <nil>};
// Left: <nil>; Value: 12; Right: <nil>;
// Left: &{0xc000070120 6 0xc000070100}; Value: 13; Right: &{<nil> 14 <nil>};
// Left: <nil>; Value: 14; Right: <nil>;
// Left: &{0xc0000700c0 13 0xc0000701e0}; Value: 15; Right: &{<nil> 16 0xc0000700a0};
// Left: <nil>; Value: 16; Right: &{0xc0000700e0 19 <nil>};
// Left: <nil>; Value: 17; Right: <nil>;
// Left: &{<nil> 17 <nil>}; Value: 19; Right: <nil>;


