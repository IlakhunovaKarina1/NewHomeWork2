package main

func main() {

}

func Merge(right, left []int) []int {

	size, i, j := len(right)+len(left), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(right)-1 && j <= len(left)-1 {
			slice[k] = left[j]
			j++
		} else if j > len(left)-1 && i <= len(right)-1 {
			slice[k] = right[i]
			i++
		} else if right[i] < left[j] {
			slice[k] = right[i]
			i++
		} else {
			slice[k] = left[j]
			j++
		}
	}
	return slice
}

