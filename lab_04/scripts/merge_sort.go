package scripts

func MergeSort(s []int) {
	if len(s) > 1 {
		middle := len(s) / 2
		MergeSort(s[:middle])
		MergeSort(s[middle:])
		merge(s, middle)
	}
}