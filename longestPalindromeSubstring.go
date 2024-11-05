package learnings

func LongPalindromeSubString(s string) string {

	/**

	https://leetcode.com/problems/longest-palindromic-substring/

	First get to the middle of the string and see if that is palindrom

	if (s - maxstr) < maxStr return maxStr -- passed

	then scan to left side of the pivot point
	   -- call recursive but reduce the string size according to left and rigt side

	check if the length is greater then half then return otherwise

	 scan right side of the string
	   -- call recursive but reduce the string size according to left and rigt side
	*/

	longpalin := ""
	leftPointer := 0
	rightPointer := 0
	if len(s) == 1 {
		return s
	} else if len(s)%2 == 0 {
		leftPointer = (len(s) / 2) - 1
		rightPointer = (len(s) / 2)
		longpalin = evenPalindrom(s, leftPointer, rightPointer)
	} else {
		leftPointer = (len(s) - 1) / 2
		longpalin = oddPalindrom(s, leftPointer)
	}

	/**
	Now move the pointer to the left
	*/
	if len(longpalin) >= len(s) {
		return longpalin
	} else {
		leftptr := 0
		if leftPointer <= 0 {
			leftptr = leftPointer
		}
		println("leftptr", leftptr)

		longpalinLeft := ""

		if leftptr == 0 {
			if s[leftptr] == s[leftptr+1] {
				longpalinLeft = string(s[leftptr]) + string(s[leftptr+1])
			} else {
				longpalinLeft = string(s[leftptr])
			}
		} else {
			longpalinLeft = recurseLeft(s, leftptr, longpalin)
		}

		//println("longpalinLeft", longpalinLeft)

		righptr := 0
		if leftPointer >= (len(s) - 1) {
			righptr = leftPointer
		}

		longpalinRight := ""
		if righptr == (len(s) - 1) {
			if s[righptr] == s[righptr-1] {
				longpalinRight = string(s[righptr]) + string(s[righptr-1])
			} else {
				longpalinRight = string(s[righptr])
			}
		} else {
			longpalinRight = recurseRight(s, righptr, longpalin)
		}

		//println("longpalinRight", longpalinRight)

		if len(longpalinLeft) > len(longpalin) {
			longpalin = longpalinLeft
		}
		if len(longpalinRight) > len(longpalin) {
			longpalin = longpalinRight
		}
		return longpalin
	}

}

func oddPalindrom(s string, pointer int) string {

	leftPointer := (pointer - 1)
	rightPointer := (pointer + 1)
	str := string(s[pointer])
	for leftPointer >= 0 && rightPointer < len(s) {
		//println("rightPointer", rightPointer, string(s[19]))
		if s[leftPointer] == s[rightPointer] {
			str = string(s[leftPointer]) + str + string(s[rightPointer])
		} else {
			return str
		}

		leftPointer--
		rightPointer++
	}
	return str

}

func evenPalindrom(s string, leftPointer int, rightPointer int) string {

	str := ""
	for leftPointer >= 0 && rightPointer < len(s) {
		// println("rightPointer", rightPointer, string(s[19]))
		if s[leftPointer] == s[rightPointer] {
			str = string(s[leftPointer]) + str + string(s[rightPointer])
		} else {
			break
		}

		leftPointer--
		rightPointer++

	}
	return str

}

func recurseLeft(s string, pointer int, longpalin string) string {
	recursePalinOdd := oddPalindrom(s, pointer)
	if len(recursePalinOdd) > len(longpalin) {
		longpalin = recursePalinOdd
	}
	recursePalinEven := evenPalindrom(s, pointer-1, pointer)
	if len(recursePalinEven) > len(longpalin) {
		longpalin = recursePalinEven
	}
	pointer--

	if pointer < 0 {
		return longpalin
	} else {
		return recurseLeft(s, pointer, longpalin)
	}
}

func recurseRight(s string, pointer int, longpalin string) string {
	recursePalinOdd := oddPalindrom(s, pointer)
	if len(recursePalinOdd) > len(longpalin) {
		longpalin = recursePalinOdd
	}
	recursePalinEven := evenPalindrom(s, pointer, pointer+1)
	if len(recursePalinEven) > len(longpalin) {
		longpalin = recursePalinEven
	}
	pointer++

	if pointer > (len(s) - 1) {
		return longpalin
	} else {
		return recurseRight(s, pointer, longpalin)
	}
}
