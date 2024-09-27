package learnings

import "strconv"

/**
Valid Sudoku
 Leet Code URL - https://leetcode.com/problems/valid-sudoku/description/
*/
func CountAndSay(n int) string {

	i := 1
	tmp := "1"

	if n == 1 {
		return "1"
	}
	for i <= n {
		//println(string(init))
		j := 0
		prev := string("")
		counter := 0
		init := ""

		if i == 1 {
			i++
			continue
		}

		//println(string(tmp))
		for j <= len(tmp)-1 {

			if prev != string("") && prev != string(tmp[j]) {
				// then there is a change so we need to reset the counter now
				init += strconv.Itoa(int(counter)) + string(prev)
				counter = 0
			}

			prev = string(tmp[j])
			j++
			counter++
		}
		init += strconv.Itoa(int(counter)) + prev
		tmp = init
		//println(i, "==", string(init))
		i++
	}
	//println(string(tmp))
	return string(tmp)

}
