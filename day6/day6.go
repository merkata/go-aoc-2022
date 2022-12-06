package day6

func MarkerPoint(input string, distinct int) int {
	var result int
	for i := 0; i < len(input); i++ {
		if i+distinct >= len(input) {
			break
		}
		seen := make(map[byte]bool)
		for j := i; j <= i+distinct; j++ {
			seen[input[j]] = true
		}
		if len(seen) == distinct+1 {
			result = i + distinct + 1
			break
		}
	}
	return result
}
