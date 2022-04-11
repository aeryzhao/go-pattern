package main

import "fmt"

func main() {
	normalBuilder := getBuilder("normal")

	director := newDirector(normalBuilder)
	normalHouse := director.builderHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)
}
