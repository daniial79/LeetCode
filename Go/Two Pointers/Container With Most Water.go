package twoPointers

func MaxArea(height []int) int {
	l, r := 0, len(height)-1
	maximumArea := 0

	for l < r {
		area := min(height[l], height[r]) * (r - l)
		maximumArea = max(maximumArea, area)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maximumArea
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
