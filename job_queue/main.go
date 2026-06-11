package main

import (
	"sync"
)

func processData(wg *sync.WaitGroup, result *[]int, data int) {
	result.add(data*2)

}

func main() {
	var wg sync.WaitGroup
	input := []int{1, 2, 34, 5}
	result := []int{}

	for _, data := range input {
		wg.Add(1)
		go processData(&wg,&result,data)
	}

}
