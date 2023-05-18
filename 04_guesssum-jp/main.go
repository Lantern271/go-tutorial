package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	正解 := 当てる数字を決定()
loop:
	for カウンタ := 1; ; カウンタ++ {
		プロンプトを表示(カウンタ)

		switch 解答, err := ユーザからの答えを取得(カウンタ); {
		case err != nil || 解答 < 1 || 10 < 解答:
			fmt.Println("1以上10以下の整数ではないのでハズレです。")
		case 解答 != 正解:
			fmt.Println("残念でした、ハズレです")
		default:
			成功時のメッセージを表示(カウンタ)
			break loop
		}
	}
}

func ユーザからの答えを取得(count int) (int, error) {
	var inp string
	fmt.Scanln(&inp)
	return strconv.Atoi(inp) //整数に変換
}

func 当てる数字を決定() int { //戻り値はint
	rand.Seed(time.Now().UnixNano()) // 乱数のseed
	num := rand.Intn(10) + 1         // 0以上10未満の整数をもらって+1する
	//fmt.Println("答えは: ", num)
	return num
}

func プロンプトを表示(count int) {
	if count == 1 {
		fmt.Print("数当てゲームです")
	}
	fmt.Printf("1以上10以下の整数を入力してください(%v回目): ", count)
}

func 成功時のメッセージを表示(count int) {
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
