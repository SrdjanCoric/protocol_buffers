package main

import (
	"fmt"
	pb "src/simple"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := example.simple.SimpleMessage{
		Id: 12345,
		isSimple: true,
		Name: "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}
	fmt.Println(sm)
}