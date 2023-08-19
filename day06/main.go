package main

import (
	"errors"
	"log"
	"strings"
	"sync"

	"aoc22go/getdata"
)

func processSignal(signal *string, length int, wg *sync.WaitGroup) (int, error) {
	for start := 0; start <= (len(*signal) - length); start++ {

		end := (start + length)

		stream := make(map[string]bool)

		for _, c := range strings.Split((*signal)[start:end], "") {
			stream[c] = true
		}

		if len(stream) == length {
			log.Println(end)
			wg.Done()
			return end, nil
		}
	}
	return -1, errors.New("Stream not found!")
}

func main() {

	ch, err := getdata.GetInputChannel("./day6/input")
	if err != nil {
		log.Fatal(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(2)

	for signal := range ch {
		go processSignal(&signal, 4, &wg)  // 1578
		go processSignal(&signal, 14, &wg) // 2178
	}

	wg.Wait()
}
