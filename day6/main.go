package main

import (
	"bufio"
	"log"
	"os"
	// "slices"
	"strings"
)

func getInputData(filename string) (data *[]string, err error) {

	data = &[]string{}

	file, err := os.Open(filename)
	if err != nil {
		return data, err
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		*data = append(*data, scan.Text())
	}

	return data, nil
}

func main() {

	data, err := getInputData("./day6/input")
	if err != nil {
		log.Fatal(err)
	}

	// var start_of_packet int

	for _, d := range *data {
		for i := 0; i <= (len(d) - 4); i++ {
			compact_stream := make(map[string]bool)
			for _, c := range strings.Split(d[i:i+4], "") {
				compact_stream[c] = true
			}
			if len(compact_stream) == 4 {
				log.Println(i+4, ":", compact_stream)
				break
			}
		}
	}
}

// 1578
