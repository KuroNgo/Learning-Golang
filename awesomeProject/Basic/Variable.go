package Basic

func TinhTongAvaB(a int, b int) int {
	// TODO: Sử dụng khai báo biến nhanh
	c := a + b
	return c
}

func TinhTongABC(a int, b float64, c int) float64 {
	// TODO: Sử dụng khai báo biến nhanh trong nhiều dòng
	d, e := (a+c)-int(b), float64(a-c)-b
	//f := float64(d) + e
	return float64(d) + e
}
