package main

import . "github.com/kkdai/react-diff"

func main() {
	nT := NewReactDiffTree(20)
	nT.InsertNote("a", 1)
	nT.InsertNote("b", 2)
	nT.InsertNote("c", 3)
	nT.InsertNote("d", 4)
	nT.InsertNote("f", 6)
	nT.InsertNote("e", 8)

	nT2 := NewReactDiffTree(20)
	nT2.InsertNote("a", 1)
	nT2.InsertNote("b", 2)
	nT2.InsertNote("c", 3)
	nT2.InsertNote("d", 5)
	nT2.InsertNote("h", 7)
	nT2.InsertNote("e", 10)

	nT.DiffTree(nT2, INSERT_MARKUP)
	nT.DisplayGraphvizTree()
}
