package day7

import (
	"log"

	"aoc22go/get_data"
)

func getResults(root_size int, directories map[string]Directory) (int, int) {

	const system_total_size int = 70000000
	const space_needed int = 30000000

	var (
		sum_size_smaller_than_1kk int
		to_delete_size            int
	)

	needs_to_delete_at_least := space_needed - (system_total_size - root_size)

	log.Println("Needs to delete at least:", needs_to_delete_at_least)

	for dkey := range directories {

		dir, _ := directories[dkey]

		d_size := dir.getTotalSize()

		if d_size <= 100000 {
			sum_size_smaller_than_1kk += d_size
		}

		if d_size >= needs_to_delete_at_least &&
			(to_delete_size == 0 || d_size < to_delete_size) {
			to_delete_size = d_size
		}
	}
	return sum_size_smaller_than_1kk, to_delete_size
}

func Resolve() {

	data, err := getdata.GetInputData("./day7/input")
	if err != nil {
		log.Fatal(err)
	}

	fs, err := createFileSystem(data)
	if err != nil {
		log.Fatal(err)
	}

	root, _ := fs.directories["/"]

	res1, res2 := getResults(root.getTotalSize(), fs.directories)

	log.Println(res1) // task1: 2031851
	log.Println(res2) // task2: 2568781

}
