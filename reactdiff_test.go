package reactdiff_test

import (
	"strconv"
	"testing"

	. "github.com/kkdai/react-diff"
)

func TestInsert(t *testing.T) {
	nT := NewReactDiffTree(20)

	nT.InsertNode("a", 1)
	if nT.GetNodeIndex("a") != 1 {
		t.Error("Basic error: First item index=", nT.GetNodeIndex("a"))
	}

	if nT.InsertNode("", 2) == true {
		t.Error("Should not insert nil value")
	}

	if nT.InsertNode("b", 20) == true {
		t.Error("Out of index insertion")
	}

	if nT.InsertNode("b", 4) == true {
		t.Error("Parent checking failed")
	}

	if nT.InsertNode("b", 1) == false {
		t.Error("Child insertion failed, b = 1")
	}

	if nT.InsertNode("c", 2) == false {
		t.Error("Child insertion failed, c=2")
	}
}

func TestRemove(t *testing.T) {
	nT := NewReactDiffTree(20)
	if nT.RemoveNode("") == true {
		t.Error("Should not remove nil value")
	}

	if nT.RemoveNode("a") == true {
		t.Error("Remove from empty tree should failed.")
	}

	nT.InsertNode("a", 1)
	if nT.RemoveNode("a") == false {
		t.Error("Cannot remove item")
	}

	if nT.GetNodeIndex("a") != -1 {
		t.Error("Try to search item already remove")
	}

	nT.InsertNode("a", 1)
	nT.InsertNode("b", 2)
	nT.InsertNode("c", 4)
	nT.InsertNode("d", 8)
	nT.InsertNode("e", 16)
	//fmt.Println("current tree:", nT.NodeList)
	nT.RemoveNode("b")
	if nT.GetNodeIndex("c") != -1 {
		t.Error("Recursive deletion failed,", nT.NodeList)
	}
	//fmt.Println("final ", nT.NodeList)
}

func TestDiffMove(t *testing.T) {
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
	nT.InsertNode("a", 1)
	nT.InsertNode("b", 2)
	nT.InsertNode("c", 3)
	nT.InsertNode("d", 4)

	nT2 := NewReactDiffTree(20)
	nT2.InsertNode("a", 1)
	nT2.InsertNode("c", 2)
	nT2.InsertNode("d", 3)

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
	nT.InsertNode("a", 1)
	nT.InsertNode("b", 2)
	nT.InsertNode("c", 4)
	nT.InsertNode("d", 8)

	nT2 := NewReactDiffTree(20)
	nT2.InsertNode("a", 1)
	nT2.InsertNode("b", 3)
	nT2.InsertNode("c", 7)
	nT2.InsertNode("d", 15)
	nT2.InsertNode("e", 2)

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
	nT.InsertNode("a", 1)
	nT.InsertNode("b", 2)
	nT.InsertNode("c", 3)
	nT.InsertNode("d", 4)
	nT.InsertNode("e", 5)

	nT2 := NewReactDiffTree(20)
	nT2.InsertNode("a", 1)
	nT2.InsertNode("c", 2)
	nT2.InsertNode("f", 3)

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
		nT2.InsertNode(strconv.Itoa(i), i)
	}
	nT2.DisplayGraphvizTree()
}

func BenchmarkAdd(b *testing.B) {
	b.ResetTimer()
	big := NewReactDiffTree(b.N)

	for i := 0; i < b.N; i++ {
		big.InsertNode(strconv.Itoa(i), i)
	}
}

func BenchmarkDel(b *testing.B) {
	big := NewReactDiffTree(10000)

	for i := 0; i < 10000; i++ {
		big.InsertNode(strconv.Itoa(i), i)
	}

	b.ResetTimer()
	for i := b.N - 1; i > 0; i-- {
		big.RemoveNode(strconv.Itoa(i))
	}
}

func BenchmarkGet(b *testing.B) {
	big := NewReactDiffTree(10000)

	for i := 0; i < 10000; i++ {
		big.InsertNode(strconv.Itoa(i), i)
	}

	b.ResetTimer()
	for i := b.N - 1; i > 0; i-- {
		big.GetNodeIndex(strconv.Itoa(i))
	}
}

func BenchmarkDiff(b *testing.B) {
	nT := NewReactDiffTree(20)
	nT.InsertNode("a", 1)
	nT.InsertNode("b", 2)
	nT.InsertNode("c", 3)
	nT.InsertNode("d", 4)
	nT.InsertNode("e", 5)

	nT2 := NewReactDiffTree(20)
	nT2.InsertNode("a", 1)
	nT2.InsertNode("c", 2)
	nT2.InsertNode("f", 3)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nTc := nT.Clone()
		nTc.DiffTree(nT2, MOVE_EXISTING|INSERT_MARKUP|REMOVE_NODE)
	}
}
