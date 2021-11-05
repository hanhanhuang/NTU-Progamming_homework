package main

import (
	"fmt"
	"strings"
)

func main() {
	var wallet int
	fmt.Print("請輸入使用者錢包餘額：")
	fmt.Scanln(&wallet)

	fmt.Printf("%-2s\t%-2s\t%-2s\t\n", "編號", "品名", "價格")
	fmt.Printf("%-2s\t%-2s\t%-2s\t\n", "1", "拉麵", "200")
	fmt.Printf("%-2s\t%-2s\t%-2s\t\n", "2", "泡芙", "30")
	fmt.Printf("%-2s\t%-2s\t%-2s\t\n", "3", "洋芋片", "35")

	var chosen string
	fmt.Print("請輸入要購買的食物編號(1, 2, 3)：")
	fmt.Scanln(&chosen)

	fmt.Println(strings.Repeat("=", 10), "交易結果", strings.Repeat("=", 10))

	var final string = "成功"
	var reason string
	var dish string

	switch chosen {
	case "1":
		wallet -= 200
		dish = "拉麵 ($200)"
	case "2":
		wallet -= 30
		dish = "泡芙 ($30)"
	case "3":
		wallet -= 35
		dish = "洋芋片 ($35)"
	default:
		final = "失敗"
		reason = "無此選項！"
	}
	if wallet < 0 {
		final = "失敗"
		reason = "餘額不足！"
	}
	fmt.Printf("%s。\n", final)
	if final != "成功" {
		fmt.Printf("原因：%s", reason)
	} else {
		fmt.Printf("品名：%s\n", dish)
		fmt.Printf("帳戶餘額：%d", wallet)
	}
}
