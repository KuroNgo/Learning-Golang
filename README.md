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
