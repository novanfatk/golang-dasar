package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "hello bro	"
	} else {
		return "hello " + name
	}
}

func main() {
	result := getHello("novan")
	fmt.Println(result)
	fmt.Println(getHello(""))
}
