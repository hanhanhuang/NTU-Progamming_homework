package main

import (
	"fmt"
	"strings"
)

type Student struct {
	id   string
	name string
}

var students []Student

func main() {
	isContinueLoop := true
	var selectd int
	for {
		fmt.Println(strings.Repeat("=", 16), "資管系學生手冊", strings.Repeat("=", 16))
		fmt.Printf("(%d)\t%s\n", 1, "顯示學生清單")
		fmt.Printf("(%d)\t%s\n", 2, "新增學生")
		fmt.Printf("(%d)\t%s\n", 3, "刪除學生")
		fmt.Printf("(%d)\t%s\n", 0, "離開")
		fmt.Println(strings.Repeat("=", 46))

		fmt.Print("請輸入功能代碼數字：")
		fmt.Scan(&selectd)

		switch selectd {
		case 1:
			showStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 0:
			isContinueLoop = false
		default:
			fmt.Println("執行錯誤。原因：無此選項")
		}

		if isContinueLoop == false {
			fmt.Println("交易結束")
			break
		} else {
			fmt.Println()
		}
	}
}

func Hasstudent() bool {
	if len(students) == 0 {
		return false
	}
	return true
}

func showStudent() {
	defer handleException()

	if Hasstudent() != true {
		panic("目前尚未有學生資料")
	}
	for i := 0; i < len(students); i++ {
		fmt.Printf("(%d)\t%s\t%s\n", i+1, students[i].id, students[i].name)
	}
}

func addStudent() {
	var newId, newName string

	defer handleException()

	fmt.Print("請輸入學生「學號」：")
	fmt.Scan(&newId)
	fmt.Print("請輸入學生「姓名」：")
	fmt.Scan(&newName)

	for i := 0; i < len(students); i++ {
		if newId == students[i].id {
			panic("此學號已經存在，無法建立")
		}
	}
	newStudent := Student{
		id:   newId,
		name: newName,
	}
	students = append(students, newStudent)
	fmt.Println("建立成功")
}

func deleteStudent() {
	var deleteId string
	var checkExist bool
	var checkNum int

	defer handleException()
	
	if Hasstudent() != true {
		panic("目前尚未有學生資料")
	}
	
	fmt.Println(strings.Repeat("-", 10))
	showStudent()
	fmt.Println(strings.Repeat("-", 10))

	fmt.Print("請輸入學生「學號」：")
	fmt.Scan(&deleteId)
	for i := 0; i < len(students); i++ {
		if deleteId == students[i].id {
			checkExist = true
			checkNum = i
		}
	}
	if checkExist != true {
		panic("此學號尚不存在，無法刪除")
	}
	students = append(students[:checkNum], students[checkNum+1:]...)
	fmt.Println("刪除成功")

}

func handleException() {
	if r := recover(); r != nil {
		fmt.Println("發生錯誤。原因：", r)
	}
}
