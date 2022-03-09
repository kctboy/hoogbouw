package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Building struct {
	Name   string
	Height int
}

var buildings []Building

func main() {
	var buildingArray [10]string

	file, err := os.Open("buildings.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	s := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		buildingArray[s] = (scanner.Text())
		//fmt.Printf(gebouw[s])
		s++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	//fmt.Println(len(buildingArray))

	//count := len(buildingArray)
	for i := 0; i <= 3; i++ {

		buildingInfo := strings.Split(buildingArray[i], " ")
		building := Building{
			Name:   buildingInfo[0],
			Height: 100, //buildingInfo[1],
		}
		buildings = append(buildings, building)
	}

	for index, building := range buildings {
		fmt.Println(index, building)

	}
}
