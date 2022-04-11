package main

import (
	"fmt"
	"sync"
)

type single struct{}

var lock = &sync.Mutex{}

var instance *single

func getInstance() *single {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("Create Single Now!")
			instance = &single{}
		} else {
			fmt.Println("Single Already Created!")
		}
	} else {
		fmt.Println("Single Already Created!")
	}
	return instance
}
