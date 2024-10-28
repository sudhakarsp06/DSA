package learnings

func ValidParenthesis(s string) bool {
	var valid bool = true
	leftparanthsis := 0
	leftflower := 0
	leftsquare := 0
	var currentPointer []int
	for pos, _ := range s {

		// If we find a closin one then we remove the last item
		if leftparanthsis <= 0 && rune(s[pos]) == ')' {
			valid = false
			break
		}

		if leftflower <= 0 && rune(s[pos]) == '}' {
			valid = false
			break
		}

		if leftsquare <= 0 && rune(s[pos]) == ']' {
			valid = false
			break
		}

		// Store the left paranthesis
		if rune(s[pos]) == '(' {
			currentPointer = append(currentPointer, pos)
			leftparanthsis++
		}

		if rune(s[pos]) == '{' {
			currentPointer = append(currentPointer, pos)
			leftflower++
		}

		if rune(s[pos]) == '[' {
			currentPointer = append(currentPointer, pos)
			leftsquare++
		}

		// If we find a closin one then we remove the last item
		if leftparanthsis > 0 && rune(s[pos]) == ')' {
			if pos > 0 && rune(s[currentPointer[len(currentPointer)-1]]) != '(' {
				valid = false
				break
			}
			currentPointer = currentPointer[:len(currentPointer)-1]
			leftparanthsis--
		}

		if leftflower > 0 && rune(s[pos]) == '}' {
			if pos > 0 && rune(s[currentPointer[len(currentPointer)-1]]) != '{' {
				valid = false
				break
			}
			currentPointer = currentPointer[:len(currentPointer)-1]
			leftflower--
		}

		if leftsquare > 0 && rune(s[pos]) == ']' {
			if pos > 0 && rune(s[currentPointer[len(currentPointer)-1]]) != '[' {
				valid = false
				break
			}
			currentPointer = currentPointer[:len(currentPointer)-1]
			leftsquare--
		}

	}

	if leftparanthsis > 0 || leftflower > 0 || leftsquare > 0 {
		valid = false
	}

	return valid
}
