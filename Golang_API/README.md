# Các bước thực hiện code trên vs code
## Khởi tạo file golang
mkdir hello
cd hello

## Khởi tạo file mod, hỗ trợ cho việc thiết lập modules
go mod init example/hello

## Thực hiện viết code

## Thực hiện khởi chạy chương trình
go run . ( chạy tất cả file go )

# Import các package trong golang
thực hiện câu lệnh "go get <tên gói>" hoặc go install

# Modules
### Tạo một file module chứa các hàm hoặc phương thức dùng để sử dụng bên lớp main
Cách tạo file tương tự như trên


### Thực hiện chạy lệnh dưới để thực hiện thêm gói tự tạo vào dự án main
go mod edit -replace <tên gói>=<đường dẫn>
ví dụ: go mod edit -replace example.com/greetings=../greetings

### Sau khi thực hiện câu lệnh trên, thực hiện lệnh go mod tidy để đồng bộ dữ liệu nhưng chưa được theo dõi trong modules
go mod tidy

### Biên dịch và cài đặt ứng dụng
#### 1. Từ cmd trong đường dẫn trực tiếp đến file main, chạy lệnh go build để biên dịch code vào trong một file exe
go build

#### 2. Từ cmd trong đường dẫn trực tiếp đến file main, chạy mới exe vừa tạo để xác nhận code làm việc
hello.exe

#### 3. Phát hiện đường dẫn cài đặt Go, ở cái cmd sẽ thực hiện cài đặt gói hiện tại
go list -f '{{.Target}}'

#### 4. Thêm Go Install đường dẫn hệ thống của bạn

set PATH=%PATH%;C:\path\to\your\install\directory
or 
go env -w GOBIN=C:\path\to\your\bin 

#### 5. Mỗi lần bạn update shell path, chạy lện go install để cài đặt gói 
go install

#### 6. Chạy ứng dụng đơn giản của bạn
<tên file>

# Tạo workspace
## Khởi tạo workspace
Trong workspace directory, chạy lệnh:
go work init ./hello

## Chạy chương trình trong workspace directory
go run example.com/hello
Chương trình sẽ tạo ra một file chứa các module 

### Các câu hỏi khác
1. hiển thị thông tin API trên swagger

- Cài đặt thư viện swaggo
go get -u github.com/swaggo/swag/cmd/swag

- Tạo tệp cấu hình swagger
swag init

- tích hợp Swagger vào ứng dụng Gin
- cài đặt AI copilot cho việc nhắc lệnh code 
- Kiểm tra đã cài đặt go trên vs code chưa thông qua các bước sau
( cài đặt go giúp cho việc nhắc code được tốt hơn)
* 1. Thực hiện sử dụng phím Ctrl + Shift + P
* 2. Thực hiện cài đặt bằng câu lệnh go update