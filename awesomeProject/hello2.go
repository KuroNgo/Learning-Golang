package main

import (
	"awesomeProject/Basic"
	"fmt"
)

func main() {
	fmt.Println("1. Con trỏ")
	fmt.Println("2. Vòng lặp")
	fmt.Println("3. Đệ quy")
	fmt.Println("4. Khai báo biến")
	fmt.Println("5. External")

	var i int
	fmt.Print("Nhập số: ")
	fmt.Scan(&i)
	switch i {
	case 1:
		fmt.Println("1. Con trỏ")
		fmt.Println("2. Không là con trỏ")
		fmt.Print("Nhập sô: ")
		fmt.Scan(&i)
		switch i {
		case 1:
			fmt.Print("1. Con trỏ: ")
			Point()
			break
		case 2:
			fmt.Print("2. Không là con trỏ: ")
			NoPoint()
			break
		}
		break
	case 2:
		fmt.Println("2. Vòng lặp")
		NoPoint()
		break
	}
}

// Đầu tiên ta khởi tạo biến x
// Sau đó hàm NoPoint sẽ sao chép giá trị của x vào để xử lý và trả về giá trị là x' = 0
// Cuối cùng, do NoPoint không làm việc chính thức trên x nên giá trị của x sẽ không thay đổi
// Và trả về giá trị ban đầu là 5
func NoPoint() {
	x := 5
	Basic.NoPoint(x)
	fmt.Println(x)
}

// Đầu tiên ta khởi tạo biến x
// Sau đó hàm Point sẽ trỏ đến địa chỉ của x để xử lý x
// Cuối cùng, do Point làm việc trên x nên giá trị của x sẽ thay đổi
// Và trả về giá trị là 0
func Point() {
	x := 5
	Basic.Point(&x)
	fmt.Println(x)
}
