package reactdiff

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/awalterschulze/gographviz"
)

type DiffOption int

const (
	INSERT_MARKUP DiffOption = 1 << iota
	MOVE_EXISTING DiffOption = 1 << iota
	REMOVE_NODE   DiffOption = 1 << iota
)

//React Diff is a binary unsort tree to represent the concept of React Diff
//React Diff has optimize tree diff algorithm to optimize original tree diff O(n^3) -> O(n)
type ReactDiff struct {
	//Major node structure
	NodeList []string

	//Node set target to store all node item in this tree
	//It help to determine if any element is exist in this tree or not
	NodeSet map[string]bool
}

//Insert node into ReactDiff tree below to Node Index
//It will return the node index and success or not
//Note: If parent node not exist, will return false
func (r *ReactDiff) InsertNote(val string, nodeIndex int) bool {
	if nodeIndex > len(r.NodeList) || nodeIndex <= 0 {
		//fmt.Println("length too big or too small")
		return false
	}

	if val == "" {
		//fmt.Println("Val is nil")
		return false //cannot insert nil value
	}

	//Check if parent exist
	if nodeIndex > 1 && r.NodeList[nodeIndex/2] == "" {
		//fmt.Println("Parent is not exist:", nodeIndex, nodeIndex/2)
		return false
	}

	//Check if value already exist
	if _, exist := r.NodeSet[val]; exist {
		//fmt.Println("Element duplicated:", val)
		return false
	}

	//Reserve zero for other usage, indexing start from 1
	r.NodeList[nodeIndex] = val
	r.NodeSet[val] = true
	return true
}

func (r *ReactDiff) deleteNode(nodeIndex int) {
	if r.NodeList[nodeIndex] == "" {
		return
	}

	nextIndex := nodeIndex * 2
	if nextIndex < len(r.NodeList) && r.NodeList[nextIndex] != "" {
		r.deleteNode(nextIndex)
	}

	if nextIndex < len(r.NodeList) && r.NodeList[nextIndex+1] != "" {
		r.deleteNode(nextIndex + 1)
	}

	r.deleteSingleNode(nodeIndex)
}

func (r *ReactDiff) deleteSingleNode(nodeIndex int) {
	if nodeIndex > len(r.NodeList) {
		return
	}

	if r.NodeList[nodeIndex] == "" {
		return
	}

	val := r.NodeList[nodeIndex]
	r.NodeList[nodeIndex] = ""
	delete(r.NodeSet, val)
}

func (r *ReactDiff) RemoveNote(val string) bool {
	if len(r.NodeSet) == 0 {
		//fmt.Println("Empty tree deletion")
		return false
	}

	if _, exist := r.NodeSet[val]; !exist {
		//fmt.Println("value not exist for deletion")
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

	//fmt.Println("option=", option, " it is match with ", option&REMOVE_NODE)
	for newIndex, value := range targetTree.NodeList {
		if value == "" {
			continue
		}

		oldIndex := r.GetNodeIndex(value)

		//INSERT_MARKUP
		if (option&INSERT_MARKUP) == INSERT_MARKUP && oldIndex == -1 {
			//new node
			//fmt.Println("Insert mode: ready to insert")
			r.InsertNote(value, newIndex)
		}

		//MOVE_EXISTING
		if option&MOVE_EXISTING == MOVE_EXISTING {
			//fmt.Println("Enter move:", oldIndex, newIndex)
			if oldIndex != -1 && oldIndex < newIndex {
				//Change its address
				r.NodeList[oldIndex] = ""
				r.NodeList[newIndex] = value
			}
		}
	}

	//REMOVE_NODE
	if option&REMOVE_NODE == REMOVE_NODE {
		//fmt.Println("Enter remove node")
		for k, _ := range r.NodeSet {
			//fmt.Println("Remove check ", k)
			if _, exist := targetTree.NodeSet[k]; !exist {

				//fmt.Println("Remove =>", k)
				r.RemoveNote(k)
			}
		}
	}

	return false
}

//Print out tree structure via Graphviz
func (r *ReactDiff) DisplayGraphvizTree() {
	_, err := exec.LookPath("dot")
	if err != nil {
		fmt.Println("Error: You need to install Graphviz to display tree")
		return
	}

	graphAst, _ := gographviz.Parse([]byte(`digraph G{}`))
	graph := gographviz.NewGraph()
	gographviz.Analyse(graphAst, graph)

	r.recursiveTree2Graphviz(graph, 1)
	//graph.AddNode(defaultGraph, "a", nil)
	//graph.AddNode(defaultGraph, "b", nil)
	//graph.AddEdge("a", "b", true, nil)
	//fmt.Println(graph.String())

	ioutil.WriteFile("out.gv", []byte(graph.String()), 0666)

	system("dot out.gv -T png -o out.png")
	system("open out.png")
}

func (r *ReactDiff) recursiveTree2Graphviz(g *gographviz.Graph, index int) {
	if index >= len(r.NodeList) || r.NodeList[index] == "" {
		return
	}

	//Add self and its parent node
	//fmt.Println("Add node=", r.NodeList[index])
	g.AddNode("G", r.NodeList[index], nil)
	if index/2 != 0 {
		//fmt.Println("Add edge:", r.NodeList[index/2], "->", r.NodeList[index])
		g.AddEdge(r.NodeList[index/2], r.NodeList[index], true, nil)
	}

	r.recursiveTree2Graphviz(g, index*2)
	r.recursiveTree2Graphviz(g, index*2+1)
}

func system(s string) {
	cmd := exec.Command(`/bin/sh`, `-c`, s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", out.String())
}

//New a React Diff Tree with define size
//The binary tree with basic alignment with array
//0-> 1, 2
//1-> 3, 4
//2-> 5, 6 ....
func NewReactDiffTree(treeSize int) *ReactDiff {
	nodes := make([]string, treeSize)
	newRD := new(ReactDiff)
	newRD.NodeList = nodes
	newRD.NodeSet = make(map[string]bool)
	return newRD
}
