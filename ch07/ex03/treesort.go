package ex03

import (
	"bytes"
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) *tree {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}

	appendValues(values[:0], root)
	return root
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	arr := make([]int, 0)
	arr = appendValues(arr, t)
	if len(arr) == 0 {
		return fmt.Sprint("[]")
	}
	var b bytes.Buffer
	fmt.Fprint(&b, "[")
	for idx, c := range arr {
		if len(arr)-1 != idx {
			fmt.Fprintf(&b, "%d ", c)
		} else {
			fmt.Fprintf(&b, "%d", c)
		}
	}
	fmt.Fprint(&b, "]")

	return b.String()
}
