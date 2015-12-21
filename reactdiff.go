package reactdiff

import "fmt"

type DiffOption int

const (
	INSERT_MARKUP DiffOption = 1 << iota
	MOVE_EXISTING DiffOption = 1 << iota
	REMOVE_NODE   DiffOption = 1 << iota
)

type Node struct {
	Index int
	Val   interface{}
}

func NewReactDiffTree(treeSize int) *ReactDiff {
	nodes := make([]interface{}, treeSize)
	newRD := new(ReactDiff)
	newRD.NodeList = nodes
	newRD.NodeSet = make(map[interface{}]bool)
	return newRD
}

//React Diff is a binary unsort tree to represent the concept of React Diff
//React Diff has optimize tree diff algorithm to optimize original tree diff O(n^3) -> O(n)
type ReactDiff struct {
	//Major node structure
	NodeList []interface{}

	//Node set target to store all node item in this tree
	//It help to determine if any element is exist in this tree or not
	NodeSet map[interface{}]bool
}

//Insert node into ReactDiff tree below to Node Index
//It will return the node index and success or not
//Note: If parent node not exist, will return false
func (r *ReactDiff) InsertNote(val interface{}, nodeIndex int) bool {
	if nodeIndex > len(r.NodeList) {
		fmt.Println("length too big")
		return false
	}

	if val == nil {
		fmt.Println("Val is nil")
		return false //cannot insert nil value
	}

	//Check if parent exist
	if nodeIndex != 0 && r.NodeList[nodeIndex/2] == nil {
		fmt.Println("Parent is not exist")
		return false
	}

	//Check if value already exist
	if _, exist := r.NodeSet[val]; exist {
		fmt.Println("Element duplicated")
		return false
	}

	r.NodeList[nodeIndex] = val
	r.NodeSet[val] = true
	return true
}

func (r *ReactDiff) RemoveNote(val interface{}) bool {
	return false
}

//Return node index via node value, return -1 if node is not exist
func (r *ReactDiff) GetNodeIndex(searchTarget interface{}) int {
	for index, value := range r.NodeList {
		if value == searchTarget {
			return index
		}
	}
	return -1
}

// Diff Tree will diff with input target tree, if not identical will replace to new one
// Return true if two tree is identical, false will replace to new one with React Diff Algorithm
func (r *ReactDiff) DiffTree(targetTree *ReactDiff, option DiffOption) bool {
	return false
}

//Print out tree structure
func (r *ReactDiff) DisplayTree() {
}
