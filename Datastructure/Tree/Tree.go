package tree

type TreeNode struct {
	value     int
	LeftNode  *Tree
	RightNode *Tree
}
type Tree struct {
	lenght int
	Root   *TreeNode
}
