package main

import "fmt"

func main() {
	fmt.Println("hello Go")
	fmt.Println(greeting("MuxueAvid"))
}

func greeting(name string) string {
	return "Hello, " + name
}
