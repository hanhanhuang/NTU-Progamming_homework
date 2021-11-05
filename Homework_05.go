package main

import "fmt"

func main() {
	Movie := make(map[string]int)
	var choose int
	continueLoop := true

	var movie_name string
	var movie_price int

	for {
		showMenu()
		fmt.Print("請輸入代號：")
		fmt.Scan(&choose)

		switch choose {
		case 1:
			fmt.Print("\n請輸入要新增的電影名稱：")
			fmt.Scan(&movie_name)

			_, ok := Movie[movie_name]
			if ok {
				fmt.Printf("執行錯誤。原因：「 %s 」名稱重複\n", movie_name)
			} else {
				fmt.Print("請輸入要新增的電影價格：")
				fmt.Scan(&movie_price)
				Movie[movie_name] = movie_price
				//fmt.Println(movie_name, "的價格 ", Movie[movie_name], "元")
			}

		case 2:
			fmt.Println("..... Movies .....")

			if len(Movie) > 0 {
				index := 1
				for movie_name, movie_price := range Movie {
					fmt.Println("(", index, ")\t", movie_name, "...$ ", movie_price)
					index += 1
				}
			} else {
				fmt.Println("目前沒有任何電影")
			}
			fmt.Println(".................")
		case 3:
			if len(Movie) > 0 {
				fmt.Println("..... Movies .....")
				index := 1
				for movie_name, movie_price := range Movie {
					fmt.Println("(", index, ")\t", movie_name, "...$ ", movie_price)
					index += 1
				}
				fmt.Println(".................")

				var number int
				fmt.Print("請輸入電影編號：")
				fmt.Scan(&number)

				if number > index-1 || number < 1 {
					fmt.Println("執行錯誤。原因：無此選項")
				} else {
					i := 1
					for movie_name, movie_price := range Movie {
						if number == i {
							fmt.Printf("「%s」購買成功。請收款%d元", movie_name, movie_price)
							break
						}
						i += 1
					}
				}
			} else {
				fmt.Println("執行錯誤。原因：尚未新增電影")
			}

		case 4:
			continueLoop = false
		default:
			fmt.Println("執行錯誤。原因：無此選項")
		}
		if continueLoop == false {
			fmt.Println("感謝使用，祝您生意興隆")
			fmt.Println("------------------------------------------------\n")
			break
		} else {
			fmt.Println("------------------------------------------------\n")
		}
	}
}

func showMenu() {
	fmt.Println("--------------- NTUST影廳 POS系統 ---------------")
	fmt.Printf("(%d)\t%s\n", 1, "管理 - 新增電影")
	fmt.Printf("(%d)\t%s\n", 2, "管理 - 檢視菜單")
	fmt.Printf("(%d)\t%s\n", 3, "銷售")
	fmt.Printf("(%d)\t%s\n", 4, "退出")
}
