package levenshtein

// Distance Function that gives Levenshtein Distance
func Distance(str1, str2 string, icost, scost, dcost int) int {
	str1_len := len(str1)							// 1
	str2_len := len(str2)							// 2
	row1 := make([]int, str2_len+1)					// 3
	row2 := make([]int, str2_len+1)					// 4

	for i := 1; i <= str2_len; i++ {					// 5
		row1[i] = i * icost								// 6
	}

	for i := 1; i <= str1_len; i++ {					// 7
		row2[0] = i * dcost								// 8

		for j := 1; j <= str2_len; j++ {				// 9
			if str1[i-1] == str2[j-1] {					// 10
				row2[j] = row1[j-1]						// 11
			} else {
				ins := row2[j-1] + icost				// 12
				del := row1[j] + dcost					// 13
				sub := row1[j-1] + scost				// 14

				if ins < del && ins < sub {				// 15
					row2[j] = ins						// 16
				} else if del < sub {					// 17
					row2[j] = del						// 18
				} else {
					row2[j] = sub						// 19
				}
			}
		}
		row1, row2 = row2, row1							// 20
	}

	return row1[len(row1)-1]
}