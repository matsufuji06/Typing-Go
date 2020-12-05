package main

import "fmt"

func main() {
	fmt.Println("Hello")
	ask("dog")
	ask("cat")
	ask("fish")
}

func ask(question string) {
	var input string
	fmt.Printf("次の単語を入力してください: %s\n", question)
	fmt.Scan(&input)
}