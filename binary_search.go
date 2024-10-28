package learnings

func RecursiveSearch(low int, high int, query int, item []int) int {
	mid := low + ((high - low) / 2)

	if item[mid] == query {
		return mid
	} else if high >= len(item)-1 || high <= 0 {
		return -1
	} else if query > item[mid] {
		return RecursiveSearch(mid+1, high, query, item)
	} else if query < item[mid] {
		return RecursiveSearch(low, mid-1, query, item)
	} else {
		return -1
	}

}
