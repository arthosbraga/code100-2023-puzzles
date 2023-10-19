package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	jsonFile, err := os.ReadFile("puzzle.json")
	if err != nil {
		fmt.Println(err)
	}

	var arrayStrings []string

	arrayStrings = CreateArry(jsonFile)

	var bufferArraystring []string

	for _, v1 := range arrayStrings {
		for _, v2 := range arrayStrings {
			if validateAnagramy(v1, v2) && !contains(bufferArraystring, v1) {

				validateAnagramy(v1, v2)
				bufferArraystring = append(bufferArraystring, v1)
			}

		}
	}

	fmt.Println(bufferArraystring)
	fmt.Println(len(bufferArraystring))
}

func CreateArry(jsonFile []byte) []string {

	var arrayStrings []string

	var bufferString string

	for _, v := range jsonFile {

		if strings.Contains(string(v), "]") {
			bufferString = ""
			continue
		}
		if strings.Contains(string(v), "[") ||
			strings.Contains(string(v), "\"") ||
			strings.Contains(string(v), "\r") ||
			strings.Contains(string(v), " ") ||
			strings.Contains(string(v), "\n") {
			continue
		}

		if strings.Contains(string(v), ",") {
			arrayStrings = append(arrayStrings, bufferString)
			bufferString = ""
			continue
		}

		bufferString = bufferString + string(v)
	}
	slices.Sort(arrayStrings)
	return arrayStrings
}

func validateAnagramy(initial string, compare string) bool {

	if len(initial) == len(compare) && initial != compare {
		var isValid bool
		isValid = true

		for _, value := range compare {
			var test = string(value)
			if !strings.Contains(string(initial), test) {
				isValid = false
				break
			}
			initial = strings.Replace(initial, test, "", 1)

		}
		return isValid
	}

	return false
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
