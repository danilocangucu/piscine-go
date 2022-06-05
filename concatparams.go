package piscine

func ConcatParams(args []string) string {
	var newstr string
	for i, v := range args {
		if i == len(args)-1 {
			newstr += v
			break
		}
		newstr += v + "\n"
	}
	return newstr
}
