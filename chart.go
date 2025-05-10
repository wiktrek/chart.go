package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello world");
	generate_columns(4, "column 1, column 2, column 3, column 4", "10,12,14,16");
}
	func generate_columns(amount int, column_names string, values string) {
		names := strings.Split(column_names, ",")
		v := strings.Split(values, ",")
		for i := 0; i < amount; i++ {
			fmt.Println(names[i])
			fmt.Println(v[i])
		}
}