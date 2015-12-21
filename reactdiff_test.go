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
}

func TestDiff(t *testing.T) {
}
