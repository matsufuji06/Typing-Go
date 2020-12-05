package main

import "fmt"

func main() {
	ask("dog")
	ask("cat")
	ask("fish")
}

func ask(question string) {
	var input string
	fmt.Printf("次の単語を入力してください: %s\n", question)
	fmt.Scan(&input)

	if question == input {
		fmt.Println("正解!")
	} else {
		fmt.Println("不正解!")
	}
}