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

//圓形
func calculateCircle() {
	var Circle int

	fmt.Print("請輸入「整數」半徑：")
	fmt.Scan(&Circle)

	fmt.Printf("面積：%.2f\n", getCircleArea(Circle))
}
func getCircleArea(a int) float64 {
	pi := math.Pi
	circleReturn := float64(a) * float64(a) * pi
	return circleReturn
}

//三角形
func calculateTriangle() {
	var Triangle [2]int

	fmt.Print("請輸入「整數」底：")
	fmt.Scan(&Triangle[0])
	fmt.Print("請輸入「整數」高：")
	fmt.Scan(&Triangle[1])

	fmt.Printf("面積：%.2f\n", getTriangleArea(Triangle[0], Triangle[1]))
}
func getTriangleArea(a, b int) float64 {
	triangleReturn := float64(a) * float64(b) / 2
	return triangleReturn
}

//矩形
func calculateRectangle() {
	var Rectangle [2]int

	fmt.Print("請輸入「整數」底：")
	fmt.Scan(&Rectangle[0])
	fmt.Print("請輸入「整數」高：")
	fmt.Scan(&Rectangle[1])

	fmt.Printf("面積：%d\n", getRectangleArea(Rectangle[0], Rectangle[1]))
}
func getRectangleArea(a, b int) int {
	rectangleReturn := a * b
	return rectangleReturn
}
