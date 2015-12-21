package reactdiff

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
	NodeList []Node

	//Node set target to store all node item in this tree
	//It help to determine if any element is exist in this tree or not
	NodeSet map[interface{}]bool
}

//Insert node into ReactDiff tree below to Parent Index
//It will return the node index and success or not
//Note: It is binary tree, insert more than 2 node below to one parent node will return false
func (r *ReactDiff) InsertNote(val interface{}, parentIndex int) (int, bool) {
	return 0, false
}

func (r *ReactDiff) RemoveNote(val interface{}) bool {
	return false
}

//Return node index via node value, return -1 if node is not exist
func (r *ReactDiff) GetNodeIndex(searchTarget interface{}) int {
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
