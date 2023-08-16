package main

import (
	day8 "aoc22go/day8/src"
	"aoc22go/getdata"
	"log"
)

func getTestData() *[]string {
	return &[]string{"30373", "25512", "65332", "33549", "35390"}
}

func main() {

	patch_of_trees, err := getdata.GetInputData("./input")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(day8.Task1(patch_of_trees)) // 1845
	log.Println(day8.Task2(patch_of_trees)) // 230112

}
