package piscine

func ToLower(s string) string {
	newsslc := make([]rune, len(s))

	for i, v := range s {
		if v >= 'A' && v <= 'Z' {
			newsslc[i] = rune(s[i]) + 32
		} else {
			newsslc[i] = rune(s[i])
		}
	}
	var news string

	for _, v := range newsslc {
		news += string(v)
	}

	return news
}
