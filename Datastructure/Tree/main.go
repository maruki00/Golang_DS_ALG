package main

import (
	"Tree/tree"
	"fmt"
)

func main() {
	fmt.Println("Hello Tree DS.")

	var root tree.Tree
	root.Init(101)
	root.Insert(102)
	root.Insert(103)
	root.Insert(104)
	root.Insert(105)
	root.Insert(106)
	root.Insert(107)
	root.Insert(108)
	root.Insert(109)
	root.Print("Root")
}
