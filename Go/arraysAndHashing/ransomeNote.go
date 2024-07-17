package arraysAndHashing

func CanConstruct(ransomNote string, magazine string) bool {
	alphabetFrequencyTracker := make([]int, 26)

	for _, char := range magazine {
		alphabetFrequencyTracker[char-97]++
	}

	for _, char := range ransomNote {
		alphabetFrequencyTracker[char-97]--
	}

	for _, char := range alphabetFrequencyTracker {
		if char < 0 {
			return false
		}
	}

	return true
}
