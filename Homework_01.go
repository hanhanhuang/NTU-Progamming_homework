package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Repeat("=", 20), "台科大電影院", strings.Repeat("=", 20))

	fmt.Printf("%-2s\t%-2s\t%-2s\t%-2s\t%-2s\t%-2s\t%-2s\t\n", "編號", "項目", "全票", "學生", "早晚", "會員", "愛心")

	fmt.Printf("%02d\t%-3s\t$%5d\t$%5d\t$%5d\t$%5.0f\t$%5.0f\t\n", 1, "2D", 280, 280-40, 280-80, 280*0.85, 280*0.5)
	fmt.Printf("%02d\t%-3s\t$%5d\t$%5d\t$%5d\t$%5.0f\t$%5.0f\t\n", 2, "3D", 340, 340-40, 340-80, 340*0.85, 340*0.5)

	fmt.Println(strings.Repeat("=", 54))
}
