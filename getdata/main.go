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
