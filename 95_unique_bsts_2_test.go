package leetcode

func genTrees(low, high int) []*TreeNode {
	var trees []*TreeNode
	for i := low; i < high; i++ {
		lefts := genTrees(low, i)
		rights := genTrees(i+1, high)

		if len(lefts) == 0 && len(rights) == 0 {
			node := &TreeNode{Val: i}
			trees = append(trees, node)
		} else if len(lefts) == 0 {
			for k := 0; k < len(rights); k++ {
				node := &TreeNode{Val: i, Right: rights[k]}
				trees = append(trees, node)
			}
		} else if len(rights) == 0 {
			for j := 0; j < len(lefts); j++ {
				node := &TreeNode{Val: i, Left: lefts[j]}
				trees = append(trees, node)
			}
		} else {
			for j := 0; j < len(lefts); j++ {
				for k := 0; k < len(rights); k++ {
					node := &TreeNode{Val: i, Left: lefts[j], Right: rights[k]}
					trees = append(trees, node)
				}
			}
		}
	}
	return trees
}

func genTreesDP(low, high int, cache [][][]*TreeNode) []*TreeNode {
	if cache[low][high] != nil {
		return cache[low][high]
	}

	var trees []*TreeNode
	for i := low; i < high; i++ {
		lefts := genTreesDP(low, i, cache)
		rights := genTreesDP(i+1, high, cache)

		if len(lefts) == 0 && len(rights) == 0 {
			node := &TreeNode{Val: i}
			trees = append(trees, node)
		} else if len(lefts) == 0 {
			for k := 0; k < len(rights); k++ {
				node := &TreeNode{Val: i, Right: rights[k]}
				trees = append(trees, node)
			}
		} else if len(rights) == 0 {
			for j := 0; j < len(lefts); j++ {
				node := &TreeNode{Val: i, Left: lefts[j]}
				trees = append(trees, node)
			}
		} else {
			for j := 0; j < len(lefts); j++ {
				for k := 0; k < len(rights); k++ {
					node := &TreeNode{Val: i, Left: lefts[j], Right: rights[k]}
					trees = append(trees, node)
				}
			}
		}
	}

	cache[low][high] = trees
	return trees
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	cache := make([][][]*TreeNode, n+1)
	for i := 0; i < n+1; i++ {
		cache[i] = make([][]*TreeNode, n+1)
	}
	return genTreesDP(1, n+1, cache)
}
