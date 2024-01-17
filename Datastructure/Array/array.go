package main

import "fmt"

func main() {
	var values [10]int
	for i := 0; i < 10; i++ {
		values[i] = ((i << 10) * 50) >> 1
	}

	for index, value := range values {
		fmt.Printf("[%d] index  => values %d. \n ", index, value)
	}
}
