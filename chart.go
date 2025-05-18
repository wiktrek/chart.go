package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	size := 10;
	generate_columns("column 1, column 2, column 3, column 4", "10,12,14,16", size);
}
func generate_columns(column_names string, values string, size int) {
	names := strings.Split(column_names, ",")
	v := strings.Split(values, ",")
	v2 := []int{}
	for _, i := range v {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		v2 = append(v2, j)
	}
	amount := len(names)
	if amount != len(v) {
		fmt.Println("The number of columns isn't equal to the amount of values")
		return
	}
	max := slices.Max(v2)
	for i := 0; i < amount; i++ {
		fmt.Print(names[i])
		make_bar(v2[i], max, size)
		fmt.Print(v2[i])
	}
}
func make_bar(value int, max_value int, size int) {
	percent := float64(value) / float64(max_value)
	squares := math.Round(float64(size) * percent)
	i := 0
	for i < int(squares) {
		fmt.Println("   ", "x")
		i++
	}
	
	

}