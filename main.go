package main

import "fmt"

const reset = "\033"

func main() {
	fmt.Println("sijkd " + Color(RED, "I'm red!"))
	fmt.Println(.Color(BLUE, "I'm blue!"))
	fmt.Println(Color(GREEN, "I'm green!"))
	fmt.Println(Color(YELLOW, "I'm yellow!"))
	fmt.Println(Color(PURPLE, "I'm purple!"))

}
