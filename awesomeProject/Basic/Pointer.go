package Basic

import "fmt"

func pointer2() {
	var x *int //tạo biến con trỏ
	var y int
	y = 0
	x = &y          // x trỏ đến địa chỉ của y
	fmt.Println(x)  // trả ket quả về địa chỉ của x theo y
	fmt.Println(&y) // trả kết qu về địa chỉ y
	fmt.Println(*x) // trả kết quả về x sau khi lấy địa chỉ y
	fmt.Println(y)  // trả về kết quả của y
	*x = 1
	fmt.Println(*x) // trả về kết quả của x
	fmt.Println(y)  // trả về kết quả của y
}

// Tạo một hàm chứa tham số là con trỏ của biến x
// Con trỏ này sẽ thực hiện việc trỏ đến vị trí của biến x
// Thay đổi giá trị của biến x = 0
func Point(x *int) {
	*x = 0
}

// Tạo một hàm không chứa con trỏ cho biến x
// Thực hiện việc thay đổi giá trị của x = 0
func NoPoint(x int) {
	x = 0
}
