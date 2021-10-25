package scripts

func MergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := MergeSort(items[:len(items)/2])
	second := MergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}


func SelectionSort(list []int) {
	n := len(list)

	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if list[j] < list[min] {
				min = j
			}
		}
		list[i], list[min] = list[min], list[i]
	}
}

func BubbleSort(list []int) {
	len := len(list)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}