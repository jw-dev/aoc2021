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

func Day01(input string, first bool) string {
	nums := SplitToNum(input)

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

func Day02(input string, first bool) string {
	type com struct {
		dir  string
		amnt int
	}

	commands := []com{}
	for _, line := range Split(input) {
		spl := strings.Split(line, " ")
		n, _ := strconv.Atoi(spl[1])
		commands = append(commands, com{
			dir:  spl[0],
			amnt: n,
		})
	}

	if first {
		horiz := 0
		vert := 0
		for _, val := range commands {
			switch val.dir {
			case "forward":
				horiz = horiz + val.amnt
			case "down":
				vert = vert + val.amnt
			case "up":
				vert = vert - val.amnt
			}
		}

		return strconv.Itoa(horiz * vert)
	} else {
		horiz := 0
		vert := 0
		aim := 0
		for _, val := range commands {
			switch val.dir {
			case "forward":
				horiz = horiz + val.amnt
				vert = vert + aim*val.amnt
			case "down":
				aim = aim + val.amnt
			case "up":
				aim = aim - val.amnt
			}
		}

		return strconv.Itoa(horiz * vert)
	}
}
