package main

import "fmt"

func main() {
	// 二分木
	/*
				10
			   /  \
			 20	   30
			/ \      \
		   40  50     60
		  /
		 70
	*/
	root := &Node{10, nil, nil}
	root.left = &Node{20, nil, nil}
	root.right = &Node{30, nil, nil}
	root.left.left = &Node{40, nil, nil}
	root.left.right = &Node{50, nil, nil}
	root.right.right = &Node{60, nil, nil}
	root.left.left.left = &Node{70, nil, nil}

	inorder(root)
	fmt.Print(res)
}

var res []int

type Node struct {
	val   int
	left  *Node
	right *Node
}

func inorder(u *Node) {
	if u != nil {
		inorder(u.left)
		res = append(res, u.val)
		inorder(u.right)
	}
}
