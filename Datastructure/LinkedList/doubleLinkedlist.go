package main

import "fmt"

type DoubleLinkedList struct {
	val  int
	next *DoubleLinkedList
	prev *DoubleLinkedList
}

func main() {
	fmt.Println("Hello world")
}
