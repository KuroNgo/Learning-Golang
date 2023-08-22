package main

import (
	"fmt"
)

func main() {
	var chieudai, chieurong float64
	fmt.Print("Nhập chiều dài: ")
	fmt.Scan(&chieudai)

	fmt.Print("Nhập chiều rộng: ")
	fmt.Scan(&chieurong)

	fmt.Print("Diện tích hình chữ nhật là: ", hcn(chieudai, chieurong))
}

func hcn(cd, cr float64) float64 {
	dt := cd * cr
	return dt
}
