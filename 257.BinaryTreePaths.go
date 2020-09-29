package code

/**
 * Definition for a binary tree node.
 */
import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	result := []string{}
	nodes := []*TreeNode{root}

	for len(nodes) > 0 {
		if nodes[len(nodes)-1].Left != nil {
			nodes = append(nodes, nodes[len(nodes)-1].Left)
		} else if nodes[len(nodes)-1].Right != nil {
			nodes = append(nodes, nodes[len(nodes)-1].Right)
		} else {
			temp := []string{}
			for _, node := range nodes {
				temp = append(temp, strconv.Itoa(node.Val))
			}
			result = append(result, strings.Join(temp, "->"))
			x_node := nodes[len(nodes)-1]
			nodes = nodes[:len(nodes)-1]
			for len(nodes) > 0 && (nodes[len(nodes)-1].Right == nil ||
				nodes[len(nodes)-1].Right == x_node) {
				x_node = nodes[len(nodes)-1]
				nodes = nodes[:len(nodes)-1]
			}
			if len(nodes) > 0 {
				nodes = append(nodes, nodes[len(nodes)-1].Right)
			}
		}
	}
	return result
}

func Dfs(result *[]string, cur string, root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil && root.Right != nil {
		*result = append(*result, cur+strconv.Itoa(root.Val))
		return
	}
	cur += strconv.Itoa(root.Val) + "->"
	Dfs(result, cur, root.Left)
	Dfs(result, cur, root.Right)
}

func binaryTreePathsRec(root *TreeNode) []string {
	result := []string{}
	Dfs(&result, "", root)
	return result
}
