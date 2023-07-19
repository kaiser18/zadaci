package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	input := flag.Arg(0)

	str := strings.Split(input, " ")
	arr := make([]int, len(str))
	for i := range arr {
		arr[i], _ = strconv.Atoi(str[i])
	}

	result := deduplicate(arr)

	result = sort(result)

	fmt.Println(result)
}

func deduplicate(data []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}

	for _, item := range data {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func sort(data []int) []int {
	for i := 0; i <= len(data)-1; i++ {
		for j := 0; j < len(data)-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}

	return data
}
