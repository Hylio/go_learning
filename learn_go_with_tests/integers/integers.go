package integers

func Add(x, y int) int {
	return x + y
}

func ShowSeqs(size int) []int {
	temp := make([]int, size)
	for i := range temp {
		temp = append(temp, i)
	}
	return temp
}
