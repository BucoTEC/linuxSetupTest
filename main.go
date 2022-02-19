package main

import "fmt"

func sayHy(name string) {
	fmt.Printf("Hello dear friend your name is: %v\n", name)
}
func hello() {
	fmt.Println("hello from my function")
}
func main() {
	fmt.Println("Hello world from golang")
	sayHy("Advan")
	hello()
}
