package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	answer := getTheNumber()
loop:
	for count := 1; ; count++ {
		printPrompt(count)

		switch num, err := readUserAnswer(count); {
		case err != nil || num < 1 || 10 < num:
			fmt.Println("1以上10以下の整数ではないのでハズレです。")
		case num != answer:
			fmt.Println("残念でした、ハズレです。")
		default:
			printSuccessMessage(count)
			break loop
		}
	}
}

func readUserAnswer(count int) (int, error) {
	var inp string
	fmt.Scanln(&inp)
	return strconv.Atoi(inp) //整数に変換
}

func getTheNumber() int { //戻り値はint
	rand.Seed(time.Now().UnixNano()) // 乱数のseed
	num := rand.Intn(10) + 1         // 0以上10未満の整数をもらって+1する
	//fmt.Println("答えは: ", num)
	return num
}

func printPrompt(count int) {
	if count == 1 {
		fmt.Print("数当てゲームです")
	}
	fmt.Printf("1以上10以下の整数を入力してください(%v回目): ", count)
}

func printSuccessMessage(count int) {
	adverb := ""
	switch {
	case count == 1:
		fmt.Printf("ビンゴ! おめでとうございます、1発であたりましたね\n")
	case 7 < count:
		adverb = "やっと"
		fallthrough
	default:
		fmt.Printf("おめでとうございます、%v回目で%sあたりましたね\n", count, adverb)
	}
}
