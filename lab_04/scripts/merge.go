package scripts

func merge(a []int, middle int) {
	temp := make([]int, len(a))
	copy(temp, a)
	i := 0
	j := middle
	indx := 0

	for i < middle && j < len(a) {
		if a[i] < a[j] {
			temp[indx] = a[i]
			i++
		} else {
			temp[indx] = a[j]
			j++
		}
		indx++
	}
	for ; i < middle; i++ {
		temp[indx] = a[i]
	}
	for ; j < len(a); j++ {
		temp[indx] = a[j]
	}
	copy(a, temp)
}