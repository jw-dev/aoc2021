package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

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

func Day03(input string, first bool) string {
	lines := Split(input)
	bits := len(lines[0])

	getFreq := func(l []int, bits int) int {
		freq := 0
		for i := 0; i < bits; i++ {
			sum := 0
			for _, n := range l {
				if n&(1<<i) > 0 {
					sum = sum + 1
					if sum >= int(math.Ceil(float64(len(l))/2.0)) {
						freq = freq | (1 << i)
					}
				}
			}
		}
		return freq
	}

	filter := func(l []int, f func(n int) bool) []int {
		res := make([]int, 0)
		for _, n := range l {
			if f(n) {
				res = append(res, n)
			}
		}
		return res
	}

	nums := make([]int, 0)
	for _, l := range lines {
		nums = append(nums, Parse(l, 2, 64))
	}

	if first {
		gamma := getFreq(nums, bits)
		epsilon := (^gamma) & ((1 << bits) - 1)
		return fmt.Sprintf("%v", gamma*epsilon)
	} else {
		oxy := make([]int, len(nums))
		scrubber := make([]int, len(nums))
		copy(oxy, nums)
		copy(scrubber, nums)

		for i := bits - 1; i >= 0; i-- {
			if len(oxy) > 1 {
				freq := getFreq(oxy, bits)
				oxy = filter(oxy, func(n int) bool { return (freq & (1 << i)) == (n & (1 << i)) })
			}

			if len(scrubber) > 1 {
				freq := getFreq(scrubber, bits)
				scrubber = filter(scrubber, func(n int) bool { return (freq & (1 << i)) != (n & (1 << i)) })
			}
		}
		return fmt.Sprintf("%v", oxy[0]*scrubber[0])
	}
}
