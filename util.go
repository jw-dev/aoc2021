package main

import (
	"strconv"
	"strings"
)

func Split(input string) []string {
	i := strings.Split(input, "\n")
	for k := range i {
		i[k] = strings.Trim(i[k], " ")
	}

	return i
}

func SplitToNum(input string) []int {
	var nums []int

	i := strings.Split(input, "\n")
	for k := range i {
		str := strings.Trim(i[k], " ")
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		nums = append(nums, n)
	}

	return nums
}

func Parse(s string, base int, n int) int {
	res, err := strconv.ParseInt(s, base, n)
	if err != nil {
		panic(err)
	}
	return int(res)
}
