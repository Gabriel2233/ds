package main

import (
	bst "github.com/Gabriel2233/ds/binarytree/binarysearchtree"
)

func main() {
    tree := bst.NewBst(5)

    tree.Add(1)
    tree.Add(8)
    tree.Add(6)
    tree.Add(3)
    tree.Add(0)
    tree.Add(10)
    tree.Print()

    tree.Remove(10)
    tree.Print()

    tree.Remove(99)
    tree.Print()

    tree.Remove(5)
    tree.Print()
}
