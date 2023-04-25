package leetcode

func Generete(numRows int) [][]int {
	pascal_arr := [][]int{}
	row_arr := []int{}

	for i := 0; i < numRows; i++ {
		row_arr = append(row_arr, 1)
		for j := 0; j < i; j++ {
			if i-j == 1 {
				row_arr = append(row_arr, 1)
			} else {
				row_arr = append(row_arr, pascal_arr[i-1][j]+pascal_arr[i-1][j+1])
			}
		}
		pascal_arr = append(pascal_arr, row_arr)
		row_arr = []int{}
	}

	return pascal_arr
}
