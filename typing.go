package main

import "fmt"
import "time"

func main() {
	start := time.Now();

	totalScore := ask(1, "html")
	totalScore += ask(2, "css")
	totalScore += ask(3, "ruby")

	end := time.Now()

	fmt.Printf("%f秒かかりました\n", (end.Sub(start)).Seconds())
	fmt.Println("スコア", totalScore)
	
}

func ask(number int, question string) int {
	var input string
	fmt.Printf("[質問%d] 次の単語を入力してください: %s\n", number, question)
	fmt.Scan(&input)

	if question == input {
		fmt.Println("正解!")
		return 10

	} else {
		fmt.Println("不正解!")
		return 0
	}
}