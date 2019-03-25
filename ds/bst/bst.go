package bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var stack []*TreeNode
    stack = append(stack, root)
    for stack[len(stack) - 1].Left != nil {
        stack = append(stack, stack[len(stack) - 1].Left)
	}
	smallest := stack[len(stack) - 1].Val - 1

	stack = append(stack, stack[len(stack) - 1].Left)

	for len(stack) > 0 {
        node := stack[len(stack) - 1]
        stack = stack[0:len(stack) - 1]
        if node.Val <= smallest {
            return false
        }
        smallest = node.Val
        if node.Right != nil {
            stack = append(stack, node.Right)
            for stack[len(stack) - 1].Left != nil {
                stack = append(stack, stack[len(stack) - 1].Left)
            }
        }
    }
	
}
