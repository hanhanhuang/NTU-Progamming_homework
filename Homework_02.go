package main

import (
	"fmt"
	"strings"
)

func main() {
	correct_pass := "11009030"
	var password string

	fmt.Println(strings.Repeat("=", 20), "台科銀行ATM", strings.Repeat("=", 20))

	fmt.Print("請輸入密碼：")
	fmt.Scanln(&password)

	var check_pass = correct_pass == password
	if check_pass {
		fmt.Println("\n歡迎使用")

		var now_money int = 0
		var minus_money int = 0
		var total int = 0

		fmt.Print("請輸入整數存款金額：")
		fmt.Scanln(&now_money)
		fmt.Print("請輸入整數提款金額：")
		fmt.Scanln(&minus_money)

		total = now_money - minus_money
		fmt.Printf("最終金額：%d\n", total)
	} else {
		println("\n登入失敗")
	}
	fmt.Println("\n\n謝謝使用~")
	fmt.Println(strings.Repeat("=", 51))
}
