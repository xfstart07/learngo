package mapreduce

func MapStrToStr(arr []string, fn func(string) string) []string {
	var newArr = []string{}
	for _, it := range arr {
		newArr = append(newArr, fn(it))
	}
	return newArr
}

func MapStrToInt(arr []string, fn func(string) int) []int {
	var newArr = []int{}
	for _, it := range arr {
		newArr = append(newArr, fn(it))
	}
	return newArr
}