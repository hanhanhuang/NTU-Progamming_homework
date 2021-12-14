package main

import "fmt"

type Product interface {
	setData()
	showData()
	getDescription() string
}

type Noodle struct {
	name  string
	price int
}
type Drink struct {
	name  string
	price int
}

func main() {
	var p Product
	var products []Product
	isContinueLoop := true
	selected := ""

	for {
		fmt.Println("============= NTUST =============")
		fmt.Printf("%d)\t%s\n", 1, "顯示清單")
		fmt.Printf("%d)\t%s\n", 2, "新增項目")
		fmt.Printf("%d)\t%s\n", 0, "離開")
		fmt.Println("=================================")

		fmt.Print("請輸入代碼：")
		fmt.Scan(&selected)

		switch selected {
		case "1":
			for i, t := range products {
				fmt.Printf("(%02d)", i+1)
				t.showData()
				fmt.Printf("\t%s\n", t.getDescription())
			}
		case "2":
			choose := noodleOrdrink()
			switch choose {
			case "1":
				p = &Noodle{}
			case "2":
				p = &Drink{}
			default:
				fmt.Println("錯誤，請重新輸入")
			}
			if choose == "1" || choose == "2" {
				p.setData()
				products = append(products, p)
			}
		case "0":
			isContinueLoop = false
		default:
			fmt.Println("錯誤，請重新輸入")
		}
		if isContinueLoop {
			fmt.Println("")
		} else {
			break
		}
	}
}

func noodleOrdrink() string {
	var p string
	fmt.Println("--------------")
	fmt.Printf("%d)\t%s\n", 1, "泡麵")
	fmt.Printf("%d)\t%s\n", 2, "飲料")
	fmt.Println("--------------")

	fmt.Print("請輸入商品代碼：")
	fmt.Scan(&p)
	return p
}

func (p *Noodle) setData() {
	fmt.Print("請輸入商品名稱：")
	fmt.Scan(&p.name)

	fmt.Print("請輸入商品價格：")
	fmt.Scan(&p.price)
}

func (p Noodle) showData() {
	fmt.Printf("\t%s ($%3d)", p.name, p.price)
}

func (p Noodle) getDescription() string {
	return "使用說明：打開包裝 > 加入調料與熱水 > 靜置三分鐘 > 享受美味餐點"
}


func (p *Drink) setData() {
	fmt.Print("請輸入商品名稱：")
	fmt.Scan(&p.name)

	fmt.Print("請輸入商品價格：")
	fmt.Scan(&p.price)
}

func (p Drink) showData() {
	fmt.Printf("\t%s ($%3d)", p.name, p.price)
}

func (p Drink) getDescription() string {
	return "使用說明：用力拉起拉環 > 飲用好喝的飲料"
}
