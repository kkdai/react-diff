package main

import . "github.com/kkdai/react-diff"

func main() {
	nT := NewReactDiffTree(20)
	nT.InsertNode("a", 1)
	nT.InsertNode("b", 2)
	nT.InsertNode("c", 3)
	nT.InsertNode("d", 4)
	nT.InsertNode("f", 6)
	nT.InsertNode("e", 8)

	nT2 := NewReactDiffTree(20)
	nT2.InsertNode("a", 1)
	nT2.InsertNode("b", 2)
	nT2.InsertNode("c", 3)
	nT2.InsertNode("d", 5)
	nT2.InsertNode("h", 7)
	nT2.InsertNode("e", 10)

	nT.DiffTree(nT2, INSERT_MARKUP)
	nT.DisplayGraphvizTree()
}
