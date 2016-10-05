package main

import "fmt"

func main() {
	fmt.Println("Hello...")
	loop()
}

func loop() {
	i := 10

	if i >= 5 {
		fmt.Println("5 up")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

}