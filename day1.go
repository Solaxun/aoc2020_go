package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func part1(s []int) int {
	for _, num := range s {
		for _, num2 := range s {
			if num+num2 == 2020 {
				return num * num2
			}
		}
	}
	return 0
}

func part2(s []int) int {
	for _, num := range s {
		for _, num2 := range s {
			for _, num3 := range s {
				if num+num2+num3 == 2020 {
					return num * num2 * num3
				}
			}

		}
	}
	return 0
}

func main() {
	data, _ := ioutil.ReadFile("inputs/day1.txt")
	text := string(data)
	s := strings.Split(text, "\n")
	nums := make([]int, len(s))
	for i, text := range s {
		num, _ := strconv.Atoi(text)
		nums[i] = num
	}
	fmt.Println(part1(nums))
	fmt.Println(part2(nums))
}
