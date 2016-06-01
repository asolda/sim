package utils

func GetMax(pool []float64, n int) []int {
	maxIndexes := make([]int, n)
	maxValues := make([]float64, n)

	for i := 0; i < n; i++ {
		maxIndexes[i] = i
		maxValues[i] = pool[i]
	}

	for i := n; i < len(pool); i++ {
		newMax, index := pushMax(maxValues, pool[i])
		if newMax {
			maxIndexes[index] = i
		}
	}

	return maxIndexes
}

func pushMax(target []float64, value float64) (bool, int) {
	i := 0

	for i < len(target) && target[i] >= value {
		i++
	}

	if i < len(target) {
		target[i] = value
		return true, i
	}

	return false, -1
}
