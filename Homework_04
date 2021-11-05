package main

import (
	"fmt"
	"strings"
)

func main() {
	student_quan := 0
	subject_quan := 0

	fmt.Print("請輸入學生人數：")
	fmt.Scan(&student_quan)

	fmt.Print("請輸入科目數量：")
	fmt.Scan(&subject_quan)

	fmt.Println()
	fmt.Println(strings.Repeat("=", 5), "成績輸入", strings.Repeat("=", 5))

	var scores = make([][]int, student_quan)
	var sum = make([]int, student_quan)

	for i := 0; i < len(scores); i++ {
		scores[i] = make([]int, subject_quan)
		//fmt.Printf("%d ", scores[i])
	}

	for i := 0; i < student_quan; i++ {
		for j := 0; j < subject_quan; j++ {
			fmt.Printf("請輸入第%d位學生的第%d成績：", i+1, j+1)
			fmt.Scan(&scores[i][j])
			if scores[i][j] > 100 || scores[i][j] < 0 {
				fmt.Println("輸入錯誤。原因：成績超出範圍(0~100)")
				fmt.Println()
				j--
			} else {
				sum[i] += scores[i][j]
				//fmt.Printf("score[%d][%d] = %d\n", i, j, scores[i][j])
				//fmt.Printf("sum[%d] = %d\n", i, sum[i])
			}
		}
		println()
	}
	fmt.Println(strings.Repeat("=", 5), "判讀結果", strings.Repeat("=", 5))

	maxium_num := 0
	maxium_sum := 0

	for i, _ := range sum {
		if sum[i] > maxium_sum {
			maxium_sum = sum[i]
			maxium_num = i
		}
	}
	fmt.Printf("第一名：同學%d(%d分)\n", maxium_num+1, maxium_sum)
}
