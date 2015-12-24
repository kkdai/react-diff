package reactdiff_test

import (
	"strconv"
	"testing"

	. "github.com/kkdai/react-diff"
)

func TestInsert(t *testing.T) {
	nT := NewReactDiffTree(20)

	nT.InsertNote("a", 1)
	if nT.GetNodeIndex("a") != 1 {
		t.Error("Basic error: First item index=", nT.GetNodeIndex("a"))
	}

	if nT.InsertNote("", 2) == true {
		t.Error("Should not insert nil value")
	}

	if nT.InsertNote("b", 20) == true {
		t.Error("Out of index insertion")
	}

	if nT.InsertNote("b", 4) == true {
		t.Error("Parent checking failed")
	}

	if nT.InsertNote("b", 1) == false {
		t.Error("Child insertion failed, b = 1")
	}

	if nT.InsertNote("c", 2) == false {
		t.Error("Child insertion failed, c=2")
	}
}

func TestRemove(t *testing.T) {
	nT := NewReactDiffTree(20)
	if nT.RemoveNote("") == true {
		t.Error("Should not remove nil value")
	}

	if nT.RemoveNote("a") == true {
		t.Error("Remove from empty tree should failed.")
	}

	nT.InsertNote("a", 1)
	if nT.RemoveNote("a") == false {
		t.Error("Cannot remove item")
	}

	if nT.GetNodeIndex("a") != -1 {
		t.Error("Try to search item already remove")
	}

	nT.InsertNote("a", 1)
	nT.InsertNote("b", 2)
	nT.InsertNote("c", 4)
	nT.InsertNote("d", 8)
	nT.InsertNote("e", 16)
	//fmt.Println("current tree:", nT.NodeList)
	nT.RemoveNote("b")
	if nT.GetNodeIndex("c") != -1 {
		t.Error("Recursive deletion failed,", nT.NodeList)
	}
	//fmt.Println("final ", nT.NodeList)
}

func TestDiffMove(t *testing.T) {
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

	nT.DiffTree(nT2, MOVE_EXISTING)
	//fmt.Println("Result: nT=", nT.NodeList)

	if nT.GetNodeIndex("d") != 5 {
		t.Error("Move error on d", nT.GetNodeIndex("d"))
	}

	if nT.GetNodeIndex("e") != 10 {
		t.Error("Move error on e", nT.GetNodeIndex("e"))
	}
}

func TestDiffAdd(t *testing.T) {
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
	//fmt.Println("Result: nT=", nT.NodeList)

	if nT.GetNodeIndex("h") != 7 {
		t.Error("Add error on h")
	}

	if nT.GetNodeIndex("e") != 8 {
		t.Error("Add error on e")
	}
}

func TestDiffDel(t *testing.T) {
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

	nT.DiffTree(nT2, REMOVE_NODE)
	//fmt.Println("Result: nT=", nT.NodeList)

	if nT.GetNodeIndex("f") != -1 {
		t.Error("Del error on f")
	}

	if nT.GetNodeIndex("e") != 8 {
		t.Error("Del error on e")
	}
}

func TestDiffComposite1(t *testing.T) {
	nT := NewReactDiffTree(20)
	nT.InsertNote("a", 1)
	nT.InsertNote("b", 2)
	nT.InsertNote("c", 3)
	nT.InsertNote("d", 4)

	nT2 := NewReactDiffTree(20)
	nT2.InsertNote("a", 1)
	nT2.InsertNote("c", 2)
	nT2.InsertNote("d", 3)

	nT.DiffTree(nT2, MOVE_EXISTING|REMOVE_NODE)
	//fmt.Println("Result: nT=", nT.NodeList)

	if nT.GetNodeIndex("d") != 3 {
		t.Error("Composive 2: error on d:", nT.GetNodeIndex("d"))
	}

	if nT.GetNodeIndex("b") != -1 {
		t.Error("Composive 2:  error on b:", nT.GetNodeIndex("b"))
	}
}

func TestDiffComposite2(t *testing.T) {
	nT := NewReactDiffTree(20)
	nT.InsertNote("a", 1)
	nT.InsertNote("b", 2)
	nT.InsertNote("c", 4)
	nT.InsertNote("d", 8)

	nT2 := NewReactDiffTree(20)
	nT2.InsertNote("a", 1)
	nT2.InsertNote("b", 3)
	nT2.InsertNote("c", 7)
	nT2.InsertNote("d", 15)
	nT2.InsertNote("e", 2)

	nT.DiffTree(nT2, MOVE_EXISTING|INSERT_MARKUP)
	//fmt.Println("Result: nT=", nT.NodeList)

	if nT.GetNodeIndex("d") != 15 {
		t.Error("Composive 1: error on d:", nT.GetNodeIndex("d"))
	}

	if nT.GetNodeIndex("e") != 2 {
		t.Error("Composive 1:  error on e", nT.GetNodeIndex("e"))
	}
}

func TestDiffComposite3(t *testing.T) {
	nT := NewReactDiffTree(20)
	nT.InsertNote("a", 1)
	nT.InsertNote("b", 2)
	nT.InsertNote("c", 3)
	nT.InsertNote("d", 4)
	nT.InsertNote("e", 5)

	nT2 := NewReactDiffTree(20)
	nT2.InsertNote("a", 1)
	nT2.InsertNote("c", 2)
	nT2.InsertNote("f", 3)

	nT.DiffTree(nT2, MOVE_EXISTING|INSERT_MARKUP|REMOVE_NODE)
	//fmt.Println("Result: nT=", nT.NodeList)

	if nT.GetNodeIndex("f") != 3 {
		t.Error("Composive 3: error on f:", nT.GetNodeIndex("f"))
	}

	if nT.GetNodeIndex("e") != -1 {
		t.Error("Composive 3:  error on e", nT.GetNodeIndex("e"))
	}
}

func TestPrintTree(t *testing.T) {
	nT2 := NewReactDiffTree(20)
	for i := 1; i < 20; i++ {
		nT2.InsertNote(strconv.Itoa(i), i)
	}
	nT2.DisplayGraphvizTree()
}

func BenchmarkAdd(b *testing.B) {
	b.ResetTimer()
	big := NewReactDiffTree(b.N)

	for i := 0; i < b.N; i++ {
		big.InsertNote(strconv.Itoa(i), i)
	}
}

func BenchmarkDel(b *testing.B) {
	b.ResetTimer()
	big := NewReactDiffTree(b.N)

	for i := 0; i < b.N; i++ {
		big.InsertNote(strconv.Itoa(i), i)
	}

	b.ResetTimer()
	for i := b.N - 1; i > 0; i-- {
		big.RemoveNote(strconv.Itoa(i))
	}
}

func BenchmarkGet(b *testing.B) {
	b.ResetTimer()
	big := NewReactDiffTree(b.N)

	for i := 0; i < b.N; i++ {
		big.InsertNote(strconv.Itoa(i), i)
	}

	b.ResetTimer()
	for i := b.N - 1; i > 0; i-- {
		big.GetNodeIndex(strconv.Itoa(i))
	}
}

func BenchmarkDiff(b *testing.B) {
	nT := NewReactDiffTree(20)
	nT.InsertNote("a", 1)
	nT.InsertNote("b", 2)
	nT.InsertNote("c", 3)
	nT.InsertNote("d", 4)
	nT.InsertNote("e", 5)

	nT2 := NewReactDiffTree(20)
	nT2.InsertNote("a", 1)
	nT2.InsertNote("c", 2)
	nT2.InsertNote("f", 3)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nTc := nT.Clone()
		nTc.DiffTree(nT2, MOVE_EXISTING|INSERT_MARKUP|REMOVE_NODE)
	}
}
