package learnings

func MaxArea(height []int) int {

	leftPoint := 0
	rightPoint := len(height) - 1

	maxContent := 0

	for leftPoint < rightPoint {
		min := height[leftPoint]
		if height[leftPoint] > height[rightPoint] {
			min = height[rightPoint]
		}
		content := min * (rightPoint - leftPoint)

		if content > maxContent {
			maxContent = content
		}

		if height[leftPoint] < height[rightPoint] {
			leftPoint++
		} else {
			rightPoint--
		}
	}
	return maxContent
}
