package main

import "fmt"

type inode interface {
	print(string)
	clone() inode
}

type file struct {
	name string
}

func (f *file) print(split string) {
	fmt.Println(split + f.name)
}

func (f *file) clone() inode {
	return &file{name: f.name + "_clone"}
}

type folder struct {
	name     string
	children []inode
}

func (f *folder) print(split string) {
	fmt.Println(split + f.name)
	for _, i := range f.children {
		i.print(split + split)
	}
}

func (f *folder) clone() inode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildren []inode
	for _, i := range f.children {
		clone := i.clone()
		tempChildren = append(tempChildren, clone)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

func main() {
	file1 := &file{"file1"}
	file2 := &file{"file2"}
	file3 := &file{"file3"}

	folder := &folder{
		name:     "folder2",
		children: []inode{file1, file2, file3},
	}

	fmt.Println("\nPrinting hierarchy for Folder")
	folder.print("  ")

	cloneFolder := folder.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
}
