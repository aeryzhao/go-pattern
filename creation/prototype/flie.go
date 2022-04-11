package main

import "fmt"

type file struct {
	name string
}

func (f *file) print(split string) {
	fmt.Println(split + f.name)
}

func (f *file) clone() inode {
	return &file{name: f.name + "_clone"}
}
