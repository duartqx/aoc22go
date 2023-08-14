package main

import (
	"errors"
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"aoc22go/get_data"
)

type File struct {
	name string
	size int
}

type Directory struct {
	name       string
	children   map[string]*Directory
	parent     *Directory
	filesystem *FileSystem
	files      map[string]File
}

type FileSystem struct {
	directories map[string]Directory
}

func (d *Directory) getFullPath() (full_name string) {

	full_name_slice := []string{}

	for d_parent := d; d_parent != nil; d_parent = d_parent.parent {
		if d_parent.name != "" {
			full_name_slice = slices.Insert(full_name_slice, 0, d_parent.name)
		}
	}

	full_name = strings.Replace(
		strings.Join(full_name_slice, "/"),
		"//", "/", 1,
	)

	return full_name
}

func (d *Directory) getOrCreateFile(file_size int, file_name string) (File, bool) {
	f, exists := d.files[file_name]
	if !exists {
		f = File{
			size: file_size,
			name: file_name,
		}
		d.files[file_name] = f
		return f, true
	}
	return f, false
}

func (d *Directory) getOrCreateDir(dir_name string) (Directory, bool) {

	dir, err := d.filesystem.makeDirectory(dir_name, d)
	if err != nil {
		return *dir, false
	}
	return *dir, true
}

func (d *Directory) getTotalSize() (total_size int) {
	for fkey := range d.files {
		total_size += d.files[fkey].size
	}
	for dkey := range d.children {
		total_size += d.children[dkey].getTotalSize()
	}
	return total_size
}

func (f *FileSystem) makeDirectory(dir_name string, parent *Directory) (d *Directory, err error) {
	if d, exists := parent.children[dir_name]; exists {
		return d, errors.New("Dir with this name exists already!")
	}
	d = &Directory{
		name:       dir_name,
		parent:     parent,
		children:   make(map[string]*Directory),
		filesystem: f,
		files:      make(map[string]File),
	}
	f.directories[d.getFullPath()] = *d
	return d, nil
}

func (f *FileSystem) changeDirectory(dir_name string, parent *Directory) (d *Directory, err error) {

	d, exists := parent.children[dir_name]
	if !exists {
		d, err := f.makeDirectory(dir_name, parent)
		if err != nil {
			return d, err
		}
	}
	return d, err
}

func createFileSystem(data *[]string) (*FileSystem, error) {

	fs := FileSystem{
		directories: make(map[string]Directory),
	}

	var pwd *Directory

	pwd, err := fs.makeDirectory("/", &Directory{})
	if err != nil {
		return &fs, err
	}

	for _, line := range (*data)[1:] {

		c := string(line[0])
		switch {
		case regexp.MustCompile(`\$`).MatchString(c):

			command := line[2:4]

			switch command {
			case "cd":

				arg := strings.Trim(line[5:], " ")

				old_pwd_name := pwd.name

				if arg == ".." {
					pwd = pwd.parent
				} else {
					pwd, err = fs.changeDirectory(arg, pwd)
					if err != nil {
						return &fs, err
					}
				}

				log.Println(
					"Changed directory from:",
					old_pwd_name,
					"to:",
					pwd.name,
				)

			case "ls":
				continue
			}
		case regexp.MustCompile(`[0-9]`).MatchString(c):
			file := strings.Split(line, " ")
			file_size, err := strconv.Atoi(file[0])
			if err != nil {
				return &fs, err
			}
			pwd.getOrCreateFile(file_size, file[1])

		case c == "d":
			dir_data := strings.Split(line, " ")
			dir, created := pwd.getOrCreateDir(dir_data[1])
			if created {
				pwd.children[dir.name] = &dir
			}
		}
	}
	return &fs, err
}

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

func main() {

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
