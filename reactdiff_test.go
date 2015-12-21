package reactdiff_test

import (
	"fmt"
	"testing"

	. "github.com/kkdai/react-diff"
)

func TestInsert(t *testing.T) {
	nT := NewReactDiffTree(20)

	nT.InsertNote("a", 0)
	fmt.Println(nT.NodeList)
	if nT.GetNodeIndex("a") != 0 {
		t.Error("Basic error: First item index=", nT.GetNodeIndex("a"))
	}

	if nT.InsertNote(nil, 1) == true {
		t.Error("Should not insert nil value")
	}

	if nT.InsertNote("b", 20) == true {
		t.Error("Out of index insertion")
	}

	if nT.InsertNote("b", 3) == true {
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
	if nT.RemoveNote(nil) == true {
		t.Error("Should not remove nil value")
	}

	if nT.RemoveNote("a") == true {
		t.Error("Remove from empty tree should failed.")
	}

	nT.InsertNote("a", 0)
	if nT.RemoveNote("a") == false {
		t.Error("Cannot remove item")
	}

	if nT.GetNodeIndex("a") != -1 {
		t.Error("Try to search item already remove")
	}

	nT.InsertNote("a", 0)
	nT.InsertNote("b", 1)
	nT.InsertNote("c", 3)
	nT.InsertNote("d", 7)
	nT.InsertNote("e", 15)
	fmt.Println("current tree:", nT.NodeList)
	nT.RemoveNote("b")
	if nT.GetNodeIndex("c") != -1 {
		t.Error("Recursive deletion failed,", nT.NodeList)
	}
	fmt.Println("final ", nT.NodeList)
}

func TestDiff(t *testing.T) {
}
