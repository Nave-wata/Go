package main

import "fmt"

func main() {
	num := 123
	str := "ABC"

	fmt.Print("num=", num, " str=", str, "\n")	// 文末改行なし，カンマ区切りで空白なし．フォーマットなし
	fmt.Println("num=", num, " str=", str)		// 文末改行あり．カンマ区切りで空白あり．フォーマットなし
	fmt.Printf("num=%d str=%s\n", num, str)		// 文末改行なし．カンマ区切りで空白なし．フォーマットあり
}