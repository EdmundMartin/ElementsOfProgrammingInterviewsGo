package ch11_searching

type MinMax struct {
	Smallest int
	Largest  int
}

func minMax(a, b int) MinMax {
	if a < b {
		return MinMax{Smallest: a, Largest: b}
	}
	return MinMax{Smallest: b, Largest: a}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func FindMinMax(arr []int) MinMax {

	if len(arr) == 1 {
		return MinMax{arr[0], arr[0]}
	}

	globalMinMax := MinMax{
		Smallest: arr[0],
		Largest:  arr[1],
	}

	for idx := 2; idx < len(arr)-1; idx += 2 {
		localMinMax := minMax(arr[idx], arr[idx+1])

		globalMinMax = MinMax{
			Smallest: min(localMinMax.Smallest, globalMinMax.Smallest),
			Largest:  max(localMinMax.Largest, globalMinMax.Largest),
		}
	}

	// If odd number of elements we need to check final value
	if len(arr)%2 != 0 {
		globalMinMax = MinMax{
			Smallest: min(globalMinMax.Smallest, arr[len(arr)-1]),
			Largest:  max(globalMinMax.Largest, arr[len(arr)-1]),
		}
	}

	return globalMinMax
}