package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	jsonFile, err := os.ReadFile("puzzle.json")
	if err != nil {
		fmt.Println(err)
	}

	var total int
	var bufferint string

	for _, v := range jsonFile {

		value, err := strconv.Atoi(string(v))

		if err != nil {
			valuNumber, _ := strconv.Atoi(bufferint)
			total = total + valuNumber
			bufferint = ""
			continue
		}

		bufferint = bufferint + strconv.Itoa(value)
	}

	fmt.Println(total)
}
