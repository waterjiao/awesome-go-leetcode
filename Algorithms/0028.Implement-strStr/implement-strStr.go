package _028_Implement_strStr

func strStr(haystack string, needle string) int {
	var exist bool
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i < len(haystack)-len(needle)+1; i ++ {
		if haystack[i] == needle[0] {
			exist = true
			for j := 1; j < len(needle); j++ {
				if haystack[i+j] != needle[j] {
					exist = false
				}
			}
		}
		if exist {
			return i
		}
	}
	return -1
}
