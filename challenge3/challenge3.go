package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type objectValue struct {
	ID    string `json:"id"`
	Value int    `json:"value"`
	Next  string `json:"next"`
}

type ObjectToList struct {
	LinkedList []objectValue `json:"linkedList"`
	Top        string        `json:"top"`
}

func main() {
	jsonFile, err := os.Open("puzzle.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var objectToList ObjectToList

	json.Unmarshal(byteValue, &objectToList)

	var bufferList []objectValue
	for len(objectToList.LinkedList) != 0 {
		for i, v := range objectToList.LinkedList {
			if v.ID == objectToList.Top {
				bufferList = append(bufferList, v)
				objectToList.LinkedList = append(objectToList.LinkedList[:i], objectToList.LinkedList[i+1:]...)
			}
			if len(bufferList) == 0 {
				continue
			}
			if v.ID == bufferList[len(bufferList)-1].Next {
				bufferList = append(bufferList, v)
				if len(objectToList.LinkedList) == i {
					objectToList.LinkedList = objectToList.LinkedList[:i-1]
					continue
				}
				if len(objectToList.LinkedList) > 1 {
					objectToList.LinkedList = append(objectToList.LinkedList[:i], objectToList.LinkedList[i+1:]...)
				}

			}
		}
	}

	var result []int

	for _, v := range bufferList {
		result = append(result, v.Value)
	}

	fmt.Println(result)
}
