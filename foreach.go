// Write a function ForEach that, for an int slice, applies a function on each element of that slice.

package piscine

func ForEach(f func(int), a []int) {
	for _, results := range a {
		f(results)
	}
}
