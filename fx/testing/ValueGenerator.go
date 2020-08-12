package testing

import "strconv"

func GenerateStrings(prefix string, n int) []string{
	values := []string{}
	for i:=1;i<=n;i++ {
		values = append(values, prefix + strconv.Itoa(i))
	}
	return values
}

func GenerateInts(val int, increment int, n int) []int{
	values := []int{}
	current := val
	for i:=1;i<=n;i++ {
		values = append(values, current)
		current = current + increment
	}
	return values
}

func GenerateFloats(val float64, increment float64, n int) []float64 {
	values := []float64{}
	current := val
	for i:=1;i<=n;i++ {
		values = append(values, current)
		current = current + increment
	}
	return values
}
