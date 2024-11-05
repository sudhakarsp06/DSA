package learnings

import "strings"

/**
Valid Sudoku
 Leet Code URL - https://leetcode.com/problems/zigzag-conversion/description/
*/
func Convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	j := 0

	i := 0

	arr2D := make([][]string, numRows)

	direction := "south"

	for j < len(s) {

		// Dynamic array

		arr2D[i] = append(arr2D[i], string(s[j]))

		j++

		if i == numRows-1 {
			direction = "north"
		} else if i == 0 {
			direction = "south"
		}

		if direction == "south" {

			i++

		} else if direction == "north" {
			i--
		}
	}

	k := 0

	str := ""

	for k <= numRows-1 {

		str += strings.Join(arr2D[k], "")

		k++
	}

	return str

}
