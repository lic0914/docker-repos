package main

import (
    "os"
    "fmt"
)
func main(){
	name, _ := os.Hostname()
	fmt.Println(name,"\n")
}