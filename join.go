package piscine

func Concat(str1 string, str2 string) string {
	str3 := str1 + str2
	return str3
}

func Join(strs []string, sep string) string {
	result := ""
	for i := range strs {
		if i < len(strs)-1 {
			result = Concat(result, strs[i]) + sep
		} else {
			result = Concat(result, strs[i])
		}
	}
	return result
}
