package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part1(xmove, ymove int, grid []string) int {
	xloc, yloc := 0, 0
	treecount := 0
	for yloc < len(grid) {
		row := grid[yloc]
		if maybetree := string(row[xloc%len(row)]); maybetree == "#" {
			treecount++
		}
		xloc += xmove
		yloc += ymove
	}
	return treecount
}

func part2(grid []string, moves ...[]int) int {
	res := 1
	for _, m := range moves {
		res *= part1(m[0], m[1], grid)
	}
	return res
}

func main() {
	data, _ := ioutil.ReadFile("inputs/day3.txt")
	text := string(data)
	grid := strings.Split(text, "\n")

	fmt.Println(part1(3, 1, grid))
	moves := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	fmt.Println(part2(grid, moves...))
}
