package learnings

/**
Valid Sudoku
 Leet Code URL - https://leetcode.com/problems/valid-sudoku/description/
*/
func ValidSudoku(board [][]byte) bool {

	dot := []byte(".")
	// Firt check the horizontal row
	for rowindex, _ := range board {
		rowindex = int(rowindex)
		var vartints = make(map[string]int)
		for _, value := range board[rowindex] {

			if string(value) != string(dot) {
				vartints[string(value)]++
			}

			if vartints[string(value)] > 1 {
				return false
			}

		}
	}

	// Check the vertical row
	for _, rowIndex := range [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8} {
		var vartints = make(map[string]int)

		for _, colIndex := range [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8} {
			value := board[colIndex][rowIndex]

			if string(value) != string(dot) {
				vartints[string(value)]++
			}

			if vartints[string(value)] > 1 {
				return false
			}
		}
	}

	corners := [][]int{{0, 0}, {0, 3}, {0, 6}, {3, 0}, {3, 3}, {3, 6}, {6, 0}, {6, 3}, {6, 6}}

	for _, corner := range corners {
		x := corner[0]
		y := corner[1]
		count := 1
		var vartintst = make(map[byte]int)
		for count <= 9 {

			if string(board[x][y]) != string(dot) {
				vartintst[board[x][y]]++
			}
			if vartintst[board[x][y]] > 1 {
				return false
			}
			y++
			// reset
			if count%3 == 0 {
				y = corner[1]
				x++
			}
			count++
		}

	}

	return true

}
