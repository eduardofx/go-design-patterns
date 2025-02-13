package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	data string
}

var instance *Singleton
var mutex = &sync.Mutex{}

func GetInstance() *Singleton {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			instance = &Singleton{data: "Lazy singleton Initialized"}
		}
	}
	return instance
}

func main() {

	s1 := GetInstance()
	s2 := GetInstance()

	fmt.Println("S1 and S2 are same instance: ", s1 == s2)
	fmt.Println("Singleton S1: ", s1.data)
}
