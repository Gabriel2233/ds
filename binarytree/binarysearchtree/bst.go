package bst

import (
	"fmt"
)

const count = 10

type BinarySearchTree struct {
	root     *Node
	numNodes int
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewNode(val int) *Node {
	return &Node{
		data:  val,
		left:  nil,
		right: nil,
	}
}

func NewBst(val int) *BinarySearchTree {
	node := NewNode(val)
	return &BinarySearchTree{
		root:     node,
		numNodes: 1,
	}
}

func (t *BinarySearchTree) Add(val int) {
    if contains(t.root, val) {
        return
    }
    t.root = add(t.root, val)
}

func (t *BinarySearchTree) Contains(val int) bool {
    return contains(t.root, val)
}

func (t *BinarySearchTree) Remove(val int) bool {
    if contains(t.root, val) {
        t.root = remove(t.root, val)
        return true
    }
    return false
}

func (t *BinarySearchTree) Print() {
    fmt.Println("==================================")
    print(t.root, 0)
    fmt.Println("==================================")
}

/*
   insert 5
       2
     /   \
    1     3

add(2, 5)
    2 < 5
        2.right = add(3, 5)
add(3, 5)
    3 < 5
        3.right = add(nil, 5)
    return 3
add(nil, 5)
    nil
        return node(5)
return 2
*/

func add(root *Node, val int) *Node {
	if root == nil {
		return NewNode(val)
	}

	if val > root.data {
		root.right = add(root.right, val)
	}

	if val < root.data {
		root.left = add(root.left, val)
	}

	return root
}

/*

   contains 0
       4
     /   \
    1    10

contains(4, 0)
    4 > 0
        return contains(1, 0)
contains(1, 0)
    1 > 0
        return contains(nil, 0)
contains(nil, 0)
    nil
        return false
*/

func contains(root *Node, val int) bool {
	if root == nil {
		return false
	}

	if val > root.data {
		return contains(root.right, val)
	} else if val < root.data {
		return contains(root.left, val)
	} else {
		return true
	}
}

/*
 remove 6, 10, and 3
        5
   3        7
1    4    6   9

------------
remove(5, 3)
    5 > 3
        5.left = remove(3, 3)
remove(3, 3)
    3 == 3
        3.l -> ok && 3.r -> ok
            4 = findSmallestNodeInRightSubtree(3.r) // or findBiggestNodeInLeftSubtree(3.l)
            3.val = 4.val

            3.r = remove(4, 4)
    return 3
remove(4, 4)
    4 == 4
        !4.l && !4.r
            return nil
return 5
------------
remove(5, 6)
    5 < 6
        5.right = remove(7, 6)
remove (7, 6)
    7 > 6
        7.left = remove(6, 6)
    return 7
remove(6, 6)
    6 == 6
        6.l && !6.r
            return NULL
return 5
------------
remove(5, 10)
    5 < 10
        5.right = remove(7, 10)
remove(7, 10)
    7 < 10
        7.right = remove(9, 10)
    return 7
remove(9, 10)
    9 < 10
        9.right = remove(nil, 10)
    return 9
remove(nil, 10)
    nil
        return nil
return 5
------------
*/

func remove(root *Node, val int) *Node {
    if root == nil {
        return nil
    }

    if root.data < val {
        root.right = remove(root.right, val)
    } else if root.data > val {
        root.left = remove(root.left, val)
    } else {
        if root.left == nil && root.right == nil {
            return nil
        } else if root.left == nil {
            return root.right
        } else if root.right == nil {
            return root.left
        } else {
            smallest := findSmallestNodeInRightSubtree(root.right)
            root.data = smallest
            root.right = remove(root.right, smallest)
        }
    }

    return root
}

/*
    print
        5
   3        7

(call print(5, 0)
    !nil
    space = 10
    (call print(7, 10)
        !nil
        space = 20
        call print(nil, 20)
            return
        <PRINT 7>
        call print(nil, 20)
            return
    )
    <PRINT 5>
    (call print(3, 10)
        !nil
        space = 20
        call print(nil, 20)
            return
        <PRINT 3>
        call print(nil, 20)
        return
    )
)
*/

func print(root *Node, space int) {
    if root == nil {
        return
    }

    space += count

    print(root.right, space)

    for i := count; i < space; i++ {
        fmt.Print(" ")
    }
    fmt.Printf("%d\n", root.data)

    print(root.left, space)
}

func findSmallestNodeInRightSubtree(root *Node) int {
    cur := root
    for cur.left != nil {
        cur = cur.left
    }
    return cur.data
}
