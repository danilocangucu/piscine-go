package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...\n")
	return true
}

func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...\n")
	return true
}

func IsDoorOpen(Door Door) bool {
	PrintStr("is the Door opened ?\n")
	return true
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?\n")
	return true
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
	}
	if IsDoorOpen(*door) {
		CloseDoor(door)
	}
}

type Door struct{}
