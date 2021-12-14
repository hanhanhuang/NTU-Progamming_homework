package main

import (
	"fmt"
	"strings"
)

/*
	自定義型別
*/
type score struct {
	chinese int
	english int
	math    int
	total   int
	average float64
}

func main() {
	myScore := score{
		chinese: 0,
		english: 0,
		math:    0,
		total:   0,
		average: 0,
	}
	input := ""
	isContinue := true

	for {
		showMenu()

		fmt.Print("請輸入選項：")
		fmt.Scan(&input)

		switch input {
		case "1":
			showScore(myScore)
		case "2":
			setScore(&myScore)
		case "3":
			setTotalAndAverage(&myScore)
		case "0":
			isContinue = false
		default:
			fmt.Println("錯誤。原因：無此選項")
		}

		if isContinue {
			fmt.Println("")
		} else {
			break
		}
	}
}

func showMenu() {
	// 顯示主選單
	fmt.Println(strings.Repeat("=", 20), "帳號管理", strings.Repeat("=", 20))
	fmt.Printf("(%d)\t%s\n", 1, "顯示成績單")
	fmt.Printf("(%d)\t%s\n", 2, "設定成績")
	fmt.Printf("(%d)\t%s\n", 3, "分數統計")
	fmt.Printf("(%d)\t%s\n", 0, "離開")
	fmt.Println(strings.Repeat("=", 46))
}

func showScore(myScore score) {
	// 顯示成績單
	fmt.Println(strings.Repeat("-", 10), "成績單", strings.Repeat("-", 10))
	fmt.Printf("%s：%d\n", "國文成績", myScore.chinese)
	fmt.Printf("%s：%d\n", "英文成績", myScore.english)
	fmt.Printf("%s：%d\n", "數學成績", myScore.math)
	fmt.Println(strings.Repeat(".", 26))
	fmt.Printf("%s：%d\n", "總分", myScore.total)
	fmt.Printf("%s：%.1f\n", "平均", myScore.average)
	fmt.Println(strings.Repeat("=", 26))
}

func setScore(myScore *score) {
	// 設定成績單
	fmt.Print("國文成績：")
	fmt.Scan(&myScore.chinese)

	fmt.Print("英文成績：")
	fmt.Scan(&myScore.english)

	fmt.Print("數學成績：")
	fmt.Scan(&myScore.math)
}

func setTotalAndAverage(myScore *score) {
	// 設定總分與平均
	myScore.total = myScore.chinese + myScore.english + myScore.math
	myScore.average = float64(myScore.total) / float64(3)
}
