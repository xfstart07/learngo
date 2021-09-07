package mapreduce

func Fileter(arr []int, fn func(n int) bool) []int {
	var newArr = []int{}
	for _, it := range arr {
		if fn(it) {
			newArr = append(newArr, it)
		}
	}
	return newArr
}
