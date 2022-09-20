package main

import (
	"fmt"
	"mygo/db1"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(4)
	fmt.Println("test")
	a := db1.Addanddesc(2, 2)
	fmt.Println(a)

}