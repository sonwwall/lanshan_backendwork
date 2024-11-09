package utils

func Reverse(s string) string {
	subs := s[:]
	var result string

	n := len(s)
	for i := 0; i < n; i++ {
		result = result + string(subs[n-i-1])
	}
	return result
}
func IsPalindrome(s string) bool {
	subs := s[:]
	var result string

	n := len(s)
	for i := 0; i < n; i++ {
		result = result + string(subs[n-i-1])
	}
	if s == result {
		return true
	} else {
		return false
	}
}
