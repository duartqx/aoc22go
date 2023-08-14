package main

import (
	"slices"
	"strings"
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
