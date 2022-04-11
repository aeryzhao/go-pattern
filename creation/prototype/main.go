package main

import "fmt"

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
