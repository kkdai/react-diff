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
		fmt.Println("Parent is not exist:", nodeIndex, nodeIndex/2)
		return false
	}

	//Check if value already exist
	if _, exist := r.NodeSet[val]; exist {
		fmt.Println("Element duplicated:", val)
		return false
	}

	r.NodeList[nodeIndex] = val
	r.NodeSet[val] = true
	return true
}

func (r *ReactDiff) deleteNode(nodeIndex int) {
	if r.NodeList[nodeIndex] == nil {
		return
	}

	nextIndex := nodeIndex*2 + 1
	if nextIndex < len(r.NodeList) && r.NodeList[nextIndex] != nil {
		r.deleteNode(nextIndex)
	}

	if nextIndex < len(r.NodeList) && r.NodeList[nextIndex+1] != nil {
		r.deleteNode(nextIndex + 1)
	}

	r.deleteSingleNode(nodeIndex)
}

func (r *ReactDiff) deleteSingleNode(nodeIndex int) {
	if nodeIndex > len(r.NodeList) {
		return
	}

	if r.NodeList[nodeIndex] == nil {
		return
	}

	val := r.NodeList[nodeIndex]
	r.NodeList[nodeIndex] = nil
	delete(r.NodeSet, val)
}

func (r *ReactDiff) RemoveNote(val interface{}) bool {
	if len(r.NodeSet) == 0 {
		fmt.Println("Empty tree deletion")
		return false
	}

	if _, exist := r.NodeSet[val]; !exist {
		fmt.Println("value not exist for deletion")
		return false
	}

	//Remove node and its child nodes
	for index, v := range r.NodeList {
		if v == val {
			r.deleteNode(index)
		}
	}

	return true
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

//New a React Diff Tree with define size
//The binary tree with basic alignment with array
//0-> 1, 2
//1-> 3, 4
//2-> 5, 6 ....
func NewReactDiffTree(treeSize int) *ReactDiff {
	nodes := make([]interface{}, treeSize)
	newRD := new(ReactDiff)
	newRD.NodeList = nodes
	newRD.NodeSet = make(map[interface{}]bool)
	return newRD
}
