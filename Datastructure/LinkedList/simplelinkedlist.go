package main

type SimpleLinkedList struct {
	value int
	next  *SimpleLinkedList
}

var (
	container *SimpleLinkedList = nil
)

func add(value int) {
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
	tmpContainer := container
	for {
		if tmpContainer == nil {
			break
		}
		println("value : ", tmpContainer.value)
		tmpContainer = tmpContainer.next
	}
}

func delete(value int) {
	tmpContainer := container
	for {
		if tmpContainer == nil {
			break
		}
		if tmpContainer.value == value {
			for {
				println("---")
				if tmpContainer == nil {

					break
				}
				tmpContainer = tmpContainer.next
			}
		}
		// if tmpContainer.next !=nil{
		// 	if tmpContainer
		// }
	}
}

func main() {
	add(1)
	add(2)
	add(3)
	add(4)
	list()
	add(3)
	add(4)
	list()
	add(1)
	add(2)
	delete(1)
	list()
}
