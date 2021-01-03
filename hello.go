package main // All executables have "package main" and "func main" as their entrypoint

import "fmt"

func main() {
	printToConsole("Hello World")
}

func printToConsole(input string) {
	fmt.Println(input)
}
