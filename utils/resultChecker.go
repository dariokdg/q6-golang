package utils

func GetNumberOfMatches(user []int, system []int) int {
	matches := 0
	for _, userNumber := range user {
		for _, systemNumber := range system {
			if userNumber == systemNumber {
				matches++
			}
		}
	}
	return matches
}
