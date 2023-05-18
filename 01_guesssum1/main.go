package main

import (
	"fmt"
	"strconv"
)

func main() {
	answer := 4
	fmt.Print("数当てゲームです。1以上10以下の整数を入力してください: ")
	var inp string
	fmt.Scanln(&inp) // 文字列として読み込み
	// inpの値を書き換えてもらうためにポインタを渡す
	// Scanlnは読み込んだ単語数を返す
	num, err := strconv.Atoi(inp) // 整数に変換
	if err != nil || num < 1 || 10 < num {
		fmt.Println("1以上10以下の整数ではないのでハズレです。")
	} else if num == answer {
		fmt.Println("ビンゴ")
	} else {
		fmt.Println("残念でした。ハズレです。")
	}
}
