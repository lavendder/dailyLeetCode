package offer

import (
	"fmt"
)

/*
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。



示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。

给定 target = 20，返回 false。



限制：

0 <= n <= 1000

0 <= m <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestOffer4() {
	a := [][]int{0: {5}, 1: {6}}
	fmt.Println(findNumberIn2DArray(a, 6))
	return
}

/*
执行用时：32 ms, 在所有 Go 提交中击败了67.98%的用户
内存消耗：6.6 MB, 在所有 Go 提交中击败了33.62%的用户
*/
func findNumberIn2DArray1(matrix [][]int, target int) bool {
	if l1 := len(matrix); l1 <= 0 || l1 > 1000 {
		return false
	}
	l2 := len(matrix[0])
	if l2 <= 0 || l2 > 1000 {
		return false
	}
	for _, v := range matrix {
		if target < v[0] || target > v[l2-1] {
			continue
		}
		for _, vv := range v {
			if vv > target {
				break
			}
			if vv == target {
				return true
			}
		}
	}
	return false
}

/*
二叉搜索树
执行用时：
28 ms, 在所有 Go 提交中击败了93.16%的用户
内存消耗：6.6 MB, 在所有 Go 提交中击败了33.62%的用户
*/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	l1 := len(matrix)
	if l1 <= 0 || l1 > 1000 {
		return false
	}
	l2 := len(matrix[0])
	if l2 <= 0 || l2 > 1000 {
		return false
	}
	j := l2 - 1
	i := 0
	for i < l1 {
		if i < 0 || j < 0 {
			break
		}
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			i++
		} else if matrix[i][j] > target {
			j--
		}

	}
	return false
}
