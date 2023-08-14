package main

import (
	"errors"
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func getData() string {
	return `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`
}

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

func main() {

	fs := FileSystem{
		directories: make(map[string]Directory),
	}

	var pwd *Directory

	pwd, err := fs.makeDirectory("/", &Directory{})
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range strings.Split(getData(), "\n")[1:] {

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
						log.Fatal(err)
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
				log.Fatal(err)
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

	sizes := []int{}

	for dkey := range fs.directories {
		dir, exists := fs.directories[dkey]
		if !exists {
			log.Fatal(dir, "Does exists")
		}
		d_size := dir.getTotalSize()
		if d_size <= 100000 {
			sizes = append(sizes, d_size)
		}
	}

	var total int
	for _, s := range sizes {
		total += s
	}

	log.Println(sizes, total)
}
