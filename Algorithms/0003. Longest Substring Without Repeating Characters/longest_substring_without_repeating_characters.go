package problem003

func lengthOfLongestSubstring(s string) int {
	var tempList = make([]rune, 0)
	var maxLen = 0
	for _, ch := range s {
		index, v := stringInList(tempList, ch)
		if v {
			if len(tempList) > maxLen {
				maxLen = len(tempList)
			}
			tempList = tempList[index+1:]
			tempList = append(tempList, ch)
		} else {
			tempList = append(tempList, ch)
			if len(tempList) > maxLen {
				maxLen = len(tempList)
			}
		}
	}
	return maxLen
}

func stringInList(tempList []rune, value rune) (int, bool) {
	for index, v := range tempList {
		if v == value {
			return index, true
		}
	}
	return -1, false
}
