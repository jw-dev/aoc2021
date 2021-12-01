package main

import (
	"strconv"
	"strings"
)

func Split(input string) []int {
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

func Day01(input string, first bool) string {
	nums := Split(input)

	if first {
		tot := 0
		for i := 1; i < len(nums); i++ {
			if nums[i] > nums[i-1] {
				tot = tot + 1
			}
		}
		return strconv.Itoa(tot)
	} else {
		tot := 0
		sum := 0
		from := 0
		to := 2
		for to < len(nums)-1 {
			from = from + 1
			to = to + 1

			newSum := 0
			for i := from; i <= to; i++ {
				newSum = newSum + nums[i]
			}

			if newSum > sum {
				tot = tot + 1
			}

			sum = newSum
		}
		return strconv.Itoa(tot)
	}
}
