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

	signal := (*data)[0]

	var characters_processed_until_packet_marker int // 1578

	for start := 0; start <= (len(signal) - 4); start++ {
		compact_stream := make(map[string]bool)
		for _, c := range strings.Split(signal[start:start+4], "") {
			compact_stream[c] = true
		}
		if len(compact_stream) == 4 {
			characters_processed_until_packet_marker = start + 4
			break
		}
	}

	log.Println(characters_processed_until_packet_marker)

	var characters_processed_until_message int // 2178

	for start := 0; start <= (len(signal) - 14); start++ {
		possible_message := make(map[string]bool)
		for _, c := range strings.Split(signal[start:start+14], "") {
			possible_message[c] = true
		}
		if len(possible_message) == 14 {
			characters_processed_until_message = start + 14
			break
		}
	}

	log.Println(characters_processed_until_message)
}
