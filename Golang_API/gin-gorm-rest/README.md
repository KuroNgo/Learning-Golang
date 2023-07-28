## GOLANG
### Thực hiện cài đặt các package của golang
 * go install
### Thực hiện chạy dự án thông qua 2 cách
 * go run . 
 * go run main.go
### Thực hiện đóng gói sản phẩm 
* go build
* Mọi thông tin về golang cũng như hướng dẫn chi tiết bạn có thể lên trang chủ của golang để tìm hiểu
* https://go.dev/doc/
* 
***
### Hoặc thực hiện cài đặt các gói riêng lẻ có trong dự án gin-gorm-rest
* Cài đặt gorm
  go get -u gorm.io/gorm
* Cài đặt các driver theo database (ở đây mình dùng postgreSQL)
  * *  go get gorm.io/driver/postgres
  * * go get gorm.io/gorm
* Cài đặt gin
  * Framework này có sẵn trong library hoặc nếu chưa có thì cài đặt thông qua lệnh
    * go get github.com/gin-gonic/gin

### Project có sử dụng postgreSQL và Docker
#### Về docker
* Thực hiện cài đặt docker và sử dụng docker CLI cài đặt các câu lệnh theo các bước trong dockerfile
    * Mục đích: Sử dụng docker để tạo môi trường ảo thống nhất môi trường làm việc khi làm team tránh bị lỗi
#### Về docker Compose
* Thực hiện cài đặt 
#### Về postgreSQL
* Do mình đã sử dụng docker để tạo môi trường làm việc nên không cần cài đặt postgreSQL mà thông qua image của docker tạo ra
    * Mục đích: tránh bị lỗi phiên bản khi team không thống nhất được version hoặc hệ điều hành máy tính làm việc