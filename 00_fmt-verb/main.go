package main

import "fmt"

func main() {
	i := 3
	answer := "Go言語"
	fmt.Printf("%d問の答えは「%s」です。\n", i, answer)

	s1 := "文字列"
	fmt.Printf("%s\n", s1)
	s2 := `バッククォートを使った、改行や
	タブが入った
		文字列`
	fmt.Printf("%s\n", s2)
	fmt.Printf("%q\n", s2)

	i = 254
	fmt.Printf("d: %d b: %b o: %o x:%x X: %X\n", i, i, i, i, i)

	f := 3.14159265358
	fmt.Printf("f: %f F: %F\ne: %e E: %E\n", f, f, f, f)
	fmt.Printf("v: %v %v %v\n", s1, i, f)
	fmt.Printf("T: %T %T %T\n", s1, i, f)
	fmt.Printf(" 「%%%%」で「%%」をひとつ出力/\n")

	fmt.Printf("%%fでs1を指定: %f\n", s1)

}
