package scripts

func LevenshteinRecursive(s1, s2 string) int {
	s1Rune, s2Rune := []rune(s1), []rune(s2)

	s1_len, s2_len := len(s1Rune), len(s2Rune)

	return distance(s1Rune, s2Rune, s1_len, s2_len)

}


func distance(s1, s2 []rune, i, j int) int {
	if i == 0 {
		return j
	}
	if j == 0 {
		return i
	}
	cost := 1
	if s1[i - 1] == s2[j - 1] {
		cost = 0
	}
	return findMin(distance(s1, s2, i, j - 1) + 1,
		distance(s1, s2, i - 1, j) + 1,
		distance(s1, s2, i - 1, j - 1) + cost)
}


func LevenshteinMatrix(s1, s2 string) (int, intMatrix) {
	var (
		len_s1, len_s2, dist, shDist int
	)

	s1Rune, s2Rune := []rune(s1), []rune(s2)

	len_s1, len_s2 = len(s1Rune), len(s2Rune)

	distMat := make(intMatrix, len_s1+1)

	for i := range distMat {
		distMat[i] = make([]int, len_s2+1)
	}

	for i := 0; i < len_s2+1; i++ {
		distMat[0][i] = i
	}

	for i := 0; i < len_s1+1; i++ {
		distMat[i][0] = i
	}

	for i := 1; i < len_s1+1; i++ {
		for j := 1; j < len_s2+1; j++ {
			cost := 1
			if s1Rune[i-1] == s2Rune[j-1] {
				cost = 0
			}
			dist = findMin(	distMat[i][j-1] + 1,
							distMat[i-1][j] + 1,
							distMat[i-1][j-1] + cost)
			distMat[i][j] = dist
		}
	}

	shDist = distMat[len_s1][len_s2]

	return shDist, distMat
}

type pair struct {
	s1 int
	s2 int
}

func LevenshteinRecursiveCache(s1, s2 string) (int) {
	var (
		len_s1, len_s2 int
	)

	s1Rune, s2Rune := []rune(s1), []rune(s2)

	len_s1, len_s2 = len(s1Rune), len(s2Rune)

	cache := make(map[pair]int)
	SubLevenshteinRecursiveCache(s1Rune, s2Rune, len_s1, len_s2, cache)
	resPair := pair{len_s1, len_s2}

	return cache[resPair]
}

func SubLevenshteinRecursiveCache(s1, s2 []rune, i, j int, cache map[pair]int) int {
	strs := pair{i, j}
	if _, ok := cache[strs]; ok {
		return cache[strs]
	}
	if i == 0 && j == 0 {
		cache[strs] = 0
	} else if j == 0{
		cache[strs] = i
	} else if i == 0 {
		cache[strs] = j
	} else {
		cost := 1
		if s1[i-1] == s2[j-1] {
			cost = 0
		}

		cache[strs] = findMin(
			SubLevenshteinRecursiveCache(s1, s2, i, j-1, cache)+1,
			SubLevenshteinRecursiveCache(s1, s2, i-1, j, cache)+1,
			SubLevenshteinRecursiveCache(s1, s2, i-1, j-1, cache)+cost)
	}

	return cache[strs]
}


func SubDamerauLevenshtein(s1, s2 []rune, i, j int) int {
	var result int
	if i == 0 && j == 0 {
		return 0
	} else if j == 0{
		return i
	} else if i == 0 {
		return j
	} else {
		cost := 1
		if s1[i-1] == s2[j-1] {
			cost = 0
		}
		result = findMin(
			SubDamerauLevenshtein(s1, s2, i, j-1)+1,
			SubDamerauLevenshtein(s1, s2, i-1, j)+1,
			SubDamerauLevenshtein(s1, s2, i-1, j-1)+cost)
		if i > 1 && j > 1 && s1[i-2] == s2[j-1] && s1[i-1] == s2[j-2] {
			result = findMin(SubDamerauLevenshtein(s1, s2, i-2, j-2)+1, result)
		}
	}
	return result
}


func DamerauLevenshtein(s1, s2 string) (int) {
	var (
		len_s1, len_s2 int
	)

	s1Rune, s2Rune := []rune(s1), []rune(s2)

	len_s1, len_s2 = len(s1Rune), len(s2Rune)

	result := SubDamerauLevenshtein(s1Rune, s2Rune, len_s1, len_s2)

	return result
}


func findMin(num ... int) int {
	min := num[0]
	for _, i := range num {
		if i < min {
			min = i
		}
	}
	return min
}