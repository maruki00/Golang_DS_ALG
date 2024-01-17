package main

type SimpleLinkedList struct {
	value int
	next  *SimpleLinkedList
}

var (
	container *SimpleLinkedList = nil
)

func (obj *SimpleLinkedList) add(value int) {
	tmp := *&SimpleLinkedList{
		value: value,
		next:  nil,
	}
	if container == nil {
		container = &tmp
		return
	}
	tmpContainer := container
	for {
		if tmpContainer.next == nil {
			tmpContainer.next = &tmp
			break
		} else {
			tmpContainer = tmpContainer.next
		}
	}
}

func list() {
	tmpContainer = container

}
