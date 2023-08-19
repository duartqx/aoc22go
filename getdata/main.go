package getdata

import (
	"bufio"
	"os"
)

func GetInputData(filename string) (data *[]string, err error) {

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

func GetInputChannel(filename string) (<-chan string, error) {

	ch := make(chan string)

	file, err := os.Open(filename)
	if err != nil {
		return ch, err
	}

	scan := bufio.NewScanner(file)

	go func() {

		defer close(ch)
		defer file.Close()

		for scan.Scan() {
			ch <- scan.Text()
		}
	}()

	return ch, nil
}
