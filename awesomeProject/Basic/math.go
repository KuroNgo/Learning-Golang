package Basic

import "fmt"

// Mục đích chính của Golang là hỗ trợ phát triển các web apps có tính sẵn sáng cao đồng thời
// Giúp việc mở ộng nhanh và dễ dàng hơn
// Golang có thể ứng dụng trong nhiều lĩnh vực như trong phát triển Web Server,
// phát trieern ứng dụng mobile hay trong các hệ thống microservice hay ERF
// ƯU điểm của Golang so với các ngôn ngữ khác như:
// Golang là một một nữ biên dịch, mã nguồn sẽ được biên dịch sang mã nhị phân (binary
// Đây là phần còn thiếu trong Javascript - Nodejs
// Với golang, chúng ta sẽ không cần cài đặt caác dependences từ server bởi vì Go đã
// loeem kết các moo đun cũng như các dependencies thành một file nhị phân
// Golang sử dụng các goroutin riêng biệt giúp tiết kiệm bộ nhớ và nâng cao hiệu nănng
// cho ứng dụng

// Hello: Tạo một hàm hello cho việc đặt tên một người
// name string: khởi tạo biến name với kiểu dữ liệu là string
// (name string) string kiểu trả về là string
func Hello(name string) string {
	//sprintf để tạo một chuỗi tùy ý
	// %v là để gọi biến vào %v
	message := fmt.Sprintf("Hi, %v Welcome!", name)
	// Trả về một cái tên trong một cái thông báo
	return message
}
