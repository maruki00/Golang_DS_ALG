package tree

//	type TreeNode struct {
//		value     int
//		LeftNode  *Tree
//		RightNode *Tree
//	}
type Tree struct {
	value     int
	LeftNode  *Tree
	RightNode *Tree
}

func (tree *Tree) Init(value int) {
	tree = &Tree{
		LeftNode:  nil,
		value:     value,
		RightNode: nil,
	}
}

// func (tree *Tree) Insert(value int) {

// 	if tree == nil {
// 		tree = &Tree{
// 			value:     value,
// 			LeftNode:  nil,
// 			RightNode: nil,
// 		}
// 		return
// 	}
// 	if tree.LeftNode == nil {
// 		tree.LeftNode = &Tree{
// 			value:     value,
// 			LeftNode:  nil,
// 			RightNode: nil,
// 		}
// 		return
// 	}
// 	if tree.RightNode == nil {
// 		tree.RightNode = &Tree{
// 			value:     value,
// 			LeftNode:  nil,
// 			RightNode: nil,
// 		}
// 		return
// 	}
// 	tree.LeftNode.Insert(value)
// 	tree.RightNode.Insert(value)

// }

func (tree *Tree) Insert(m int) {
	if tree != nil {

		if tree.LeftNode == nil {
			tree.LeftNode = &Tree{LeftNode: nil, value: m, RightNode: nil}
		} else {
			if tree.RightNode == nil {
				tree.RightNode = &Tree{LeftNode: nil, value: m, RightNode: nil}
			} else {

				if tree.LeftNode != nil {

					tree.LeftNode.Insert(m)
				} else {

					tree.RightNode.Insert(m)
				}

			}

		}

	} else {
		tree = &Tree{LeftNode: nil, value: m, RightNode: nil}
	}
}

func (tree *Tree) Print(msg string) {
	if tree == nil {
		return
	}
	println(msg, " - Value : ", tree.value)
	tree.LeftNode.Print("Left")
	tree.RightNode.Print("Right")

}
