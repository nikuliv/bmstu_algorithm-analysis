package scripts

import "sync"


func ParallelMergeSort(s []int) {
	len := len(s)

	if len > 1 {
		if len <= 2048 { // Sequential
			MergeSort(s)
		} else { // Parallel
			middle := len / 2

			var wg sync.WaitGroup
			wg.Add(1)

			go func() {
				defer wg.Done()
				ParallelMergeSort(s[:middle])
			}()

			ParallelMergeSort(s[middle:])

			wg.Wait()
			merge(s, middle)
		}
	}
}