package daily_practice

/*
54. 螺旋矩阵
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。


示例 1：


输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：


输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]


提示：

m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100
*/

func spiralOrder(matrix [][]int) []int {
	colMin, colMax := 0, len(matrix[0])
	rowMin, rowMax := 0, len(matrix)
	res := []int{}
	order := make(map[string]string, 1)
	order["r"] = "d,j,+"
	order["d"] = "l,i,+"
	order["l"] = "u,j,-"
	order["u"] = "r,i,-"
	p, q := 0, 0
	start := "r"
	for i := 0; i >= rowMin && i < rowMax; {
		for j := 0; j < colMax; {

			res = append(res, matrix[p][q])
		}
		for j := i; j < q; {
			res = append(res, matrix[p][q])
			if j == r-1 {
				i = j
			}
		}
	}
	for i := 0; i < c; i++ {
		p, q := 0, 0
		for j := i; j < r; {
			res = append(res, matrix[p][q])
			if j == r-1 {
				i = j
			}
		}

	}

}
