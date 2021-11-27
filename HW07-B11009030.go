package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Score struct {
	name  string
	chi   int
	en    int
	math  int
	total int
	aver  float64
}

var scores []Score

func main() {
	var choose int
	for {
		showMenu()
		fmt.Print("請輸入選項：")
		fmt.Scan(&choose)

		if choose == 1 {
			if len(scores) > 0 {
				fmt.Println("[Error] 已經匯入過了\n")
			} else {
				importStudentName()
				fmt.Println()
				//fmt.Println("執行成功\n")
			}
		} else if choose == 2 {
			if len(scores) > 0 {
				enterScores()
			} else {
				fmt.Println("[Error] 尚未匯入學生資料\n")
			}
		} else if choose == 3 {
			if len(scores) > 0 {
				err := exportScore()

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("執行成功\n")
				}
			} else {
				fmt.Println("[Error] 尚未匯入學生資料\n")
			}
		} else if choose == 0 {
			break
		} else {
			fmt.Println("[Error] 無此選項\n")
		}
	}
}

func showMenu() {
	fmt.Println("=============== NTUST Score System ===============")
	fmt.Printf("%d\t%s\n", 1, "匯入學生")
	fmt.Printf("%d\t%s\n", 2, "輸入成績")
	fmt.Printf("%d\t%s\n", 3, "匯出成績")
	fmt.Printf("%d\t%s\n", 0, "結束")
	fmt.Println("================================================")
}

func importStudentName() {
	importByte, err := ioutil.ReadFile("StudentList.txt")
	if err != nil {
		//error := errors.New(string(err.Error()))

		//err2 := errors.New(";kjf")
		fmt.Print(err.Error())
		//fmt.Print(err2) // 取得錯誤訊息
		//er2 := err.Error()
		//fmt.Println(er2)
		//err2.Error()
		//log.Fatal(err)
		//fmt.Println()
		//fmt.Print(err)

	}
	//fmt.Print("byte:\n", importByte)
	importString := string(importByte)
	//fmt.Print("string\n", importString)
	importArray_line := strings.Split(importString, "\r\n")
	//fmt.Print("array\n", importArray_line)

	for i := 0; i < len(importArray_line); i++ {
		importArray_tab := strings.Split(importArray_line[i], "\t")

		var tempScore Score
		tempScore.name = importArray_tab[0]

		scores = append(scores, tempScore)
	}
	//fmt.Print(scores[1].name)
}

func checkScore(scoreStr string) int {
	scoreInt, _ := strconv.Atoi(scoreStr)
	if scoreInt < 0 || scoreInt > 100 {
		scoreInt = 0
	}
	return scoreInt
}

func enterScores() {
	for i := 0; i < len(scores); i++ {
		fmt.Printf("[%s]\n", scores[i].name)

		subject := [...]string{"國文", "英文", "數學"}
		subEn := [...]string{"chi", "en", "math"}
		subPtr := [...]*int{&scores[i].chi, &scores[i].en, &scores[i].math}

		for j := 0; j < len(subject); j++ {
			fmt.Printf("\t請輸入%s成績：", subject[j])
			fmt.Scan(&subEn[j])

			*subPtr[j] = checkScore(subEn[j])
			scores[i].total += *subPtr[j]
		}
		scores[i].aver = float64(scores[i].total) / 3
	}
	fmt.Println("\n執行成功")
	//fmt.Print(scores)
}

func exportScore() error {
	var exportString string

	for i := 0; i < len(scores); i++ {
		exportString += fmt.Sprintf("[%s]\n", scores[i].name)
		exportString += fmt.Sprintf("\t請輸入國文成績：%d\n", scores[i].chi)
		exportString += fmt.Sprintf("\t請輸入英文成績：%d\n", scores[i].en)
		exportString += fmt.Sprintf("\t請輸入數學成績：%d\n", scores[i].math)
		exportString += fmt.Sprintf("總分：%d / 平均：%.2f\n\n", scores[i].total, scores[i].aver)
	}
	fmt.Print(exportString)
	data_byte := []byte(exportString)

	err := ioutil.WriteFile("ScoreBook.txt", data_byte, 0644)

	return err
}
