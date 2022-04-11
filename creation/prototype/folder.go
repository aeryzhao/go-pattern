package main

import "fmt"

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
