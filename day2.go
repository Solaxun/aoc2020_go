package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func isValid(pw string) bool {
	pat, _ := regexp.Compile(`(\d+)-(\d+) (\w+): (.+)`)
	parsed := pat.FindStringSubmatch(pw)
	_, min, max, char, pw := parsed[0], parsed[1], parsed[2], parsed[3], parsed[4]
	mini, _ := strconv.Atoi(min)
	maxi, _ := strconv.Atoi(max)
	occurances := strings.Count(pw, char)
	return occurances >= mini && occurances <= maxi
}

func isValidPt2(pw string) bool {
	pat, _ := regexp.Compile(`(\d+)-(\d+) (\w+): (.+)`)
	parsed := pat.FindStringSubmatch(pw)
	_, min, max, char, pw := parsed[0], parsed[1], parsed[2], parsed[3], parsed[4]
	mini, _ := strconv.Atoi(min)
	maxi, _ := strconv.Atoi(max)
	oneMatch := string(pw[mini-1]) == char || string(pw[maxi-1]) == char
	return oneMatch && string(pw[mini-1]) != string(pw[maxi-1])
}

func main() {
	data, _ := ioutil.ReadFile("inputs/day2.txt")
	text := string(data)
	s := strings.Split(text, "\n")
	valid2 := 0
	valid := 0
	for _, pw := range s {
		if isValidPt2(pw) {
			valid2++
		}
		if isValid(pw) {
			valid++
		}
	}
	fmt.Println(valid, valid2)
}
