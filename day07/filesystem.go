package day7

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type FileSystem struct {
	directories map[string]Directory
}

func (f *FileSystem) makeDirectory(
	dir_name string,
	parent *Directory,
) (d *Directory, err error) {
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

func (f *FileSystem) changeDirectory(
	dir_name string,
	parent *Directory,
) (d *Directory, err error) {

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

	var cwd *Directory

	cwd, err := fs.makeDirectory("/", &Directory{})
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

				old_pwd_name := cwd.name

				if arg == ".." {
					cwd = cwd.parent
				} else {
					cwd, err = fs.changeDirectory(arg, cwd)
					if err != nil {
						return &fs, err
					}
				}

				log.Println(
					"Changed directory from:",
					old_pwd_name,
					"to:",
					cwd.name,
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
			cwd.getOrCreateFile(file_size, file[1])

		case c == "d":
			dir_data := strings.Split(line, " ")
			cwd.getOrCreateDir(dir_data[1])
		}
	}
	return &fs, err
}
