package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var Problems = []func(string, bool) string{
	Day01, Day02, Day03,
}

func main() {
	prob := os.Args[1]
	splits := strings.Split(prob, "")

	n, err := strconv.Atoi(splits[0])
	if err != nil {
		log.Fatalln(err)
	}

	first := splits[1] == "a"

	reader := bufio.NewReader(os.Stdin)
	data, err := io.ReadAll(reader)

	if err != nil {
		log.Fatalln(err)
	}

	problem := Problems[n-1]
	fmt.Println(problem(string(data), first))
}
