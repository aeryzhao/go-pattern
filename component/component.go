package main

import "fmt"

type component interface {
	search(string)
}

type file struct {
	name string
}

func (f *file) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *file) getName() string {
	return f.name
}

type folder struct {
	components []component
	name       string
}

func (f *folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, c := range f.components {
		c.search(keyword)
	}
}

func (f *folder) add(c component) {
	f.components = append(f.components, c)
}

func main() {
	f1 := &file{"file1"}
	f2 := &file{"file2"}
	f3 := &file{"file3"}

	folder1 := &folder{
		components: []component{f1},
	}

	folder2 := &folder{
		components: []component{f2, f3},
	}

	folder := &folder{
		components: []component{folder1, folder2, f3},
	}

	folder.search("rose")
}
