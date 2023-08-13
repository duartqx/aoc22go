package main

import (
	"bufio"
	"errors"
	"log"
	"os"

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

func processSignal(signal *string, length int) (int, error) {
	for start := 0; start <= (len(*signal) - length); start++ {

		end := (start + length)

		stream := make(map[string]bool)

		for _, c := range strings.Split((*signal)[start:end], "") {
			stream[c] = true
		}

		if len(stream) == length {
			return end, nil
		}
	}
	return -1, errors.New("Stream not found!")
}

func main() {

	data, err := getInputData("./day6/input")
	if err != nil {
		log.Fatal(err)
	}

	signal := (*data)[0]

	signal_marker, err := processSignal(&signal, 4) // 1578
	if err != nil {
		log.Fatal(err)
	}

	signal_message, err := processSignal(&signal, 14) // 2178
	if err != nil {
		log.Fatal(err)
	}

	log.Println(signal_marker, signal_message)
}
