package piscine

func Index(s string, toFind string) int {
	lens := len(s)
	lentoFind := len(toFind)

	if lens < lentoFind {
		return -1
	} else {
		for i := 0; i <= lens-1; i++ {
			if lentoFind+i < lens {
				if toFind == s[i:lentoFind+i] {
					return i
				}
			} else {
				break
			}
		}
	}
	return -1
}
