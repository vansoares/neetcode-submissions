/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    values := []int{}
	binarySearch(root, &values)

	sort.Ints(values) // O(N log N)

	return values[k-1]
}

func binarySearch(node *TreeNode, values *[]int){
	left := node.Left
	right := node.Right
	
	*values = append(*values, node.Val)

	if left != nil {
		binarySearch(left, values)
	} 

	if right != nil {
		binarySearch(right, values)
	}
}
