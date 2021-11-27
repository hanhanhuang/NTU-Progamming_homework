package main

import (
	"fmt"
	"math"
)

func main() {
	var choose int

	continueLoop := true

	for {
		showMenu()

		fmt.Print("請輸入選項：")
		fmt.Scan(&choose)
		fmt.Println()

		switch choose {
		case 1:
			calculateCircle()
		case 2:
			calculateTriangle()
		case 3:
			calculateRectangle()
		case 0:
			continueLoop = false
		default:
			fmt.Println("[Error] 無此選項")
		}
		if continueLoop == false {
			break
		} else {
			fmt.Println()
		}
	}
}

func showMenu() {
	fmt.Println("============= 圖形計算機 =============")
	fmt.Println("1       圓形")
	fmt.Println("2       三角形")
	fmt.Println("3       矩形")
	fmt.Println("0       離開")
	fmt.Println("=====================================")
}

type Circle int

//圓形
func calculateCircle() {
	var cir Circle

	fmt.Print("請輸入「整數」半徑：")
	fmt.Scan(&cir)

	fmt.Printf("面積：%.2f\n", cir.getCircleArea())
}
func (a Circle) getCircleArea() float64 {
	pi := math.Pi
	circleReturn := float64(a) * float64(a) * pi
	return circleReturn
}

type Triangle [2]int

//三角形
func calculateTriangle() {
	var tri Triangle

	fmt.Print("請輸入「整數」底：")
	fmt.Scan(&tri[0])
	fmt.Print("請輸入「整數」高：")
	fmt.Scan(&tri[1])

	fmt.Printf("面積：%.2f\n", tri.getTriangleArea())
}
func (a Triangle) getTriangleArea() float64 {
	triangleReturn := float64(a[0]) * float64(a[1]) / 2
	return triangleReturn
}

type Rectangle [2]int

//矩形
func calculateRectangle() {
	var rec Rectangle

	fmt.Print("請輸入「整數」底：")
	fmt.Scan(&rec[0])
	fmt.Print("請輸入「整數」高：")
	fmt.Scan(&rec[1])

	fmt.Printf("面積：%d\n", rec.getRectangleArea())
}
func (a Rectangle) getRectangleArea() int {
	rectangleReturn := a[0] * a[1]
	return int(rectangleReturn)
}

