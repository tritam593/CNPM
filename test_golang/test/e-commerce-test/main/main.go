package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"test/e-commerce-test/pkg"
)

func main() {
	fmt.Println("11111")
	hashedPasswordFromDB := "$2a$10$b3iVROe8G6znRo.EXcCkOeTsOMtv51S3LQ3bQoIHhZuCFa2tVtxhu" // Lấy mật khẩu đã được băm từ cơ sở dữ liệu
	userProvidedPassword := "a" // Nhận mật khẩu văn bản từ người dùng trong quá trình đăng nhập

	err := bcrypt.CompareHashAndPassword([]byte(hashedPasswordFromDB), []byte(userProvidedPassword))
	if err == nil {
		// Mật khẩu khớp, xác thực thành công
		fmt.Println("Đăng nhập thành công!")
	} else {
		// Mật khẩu không khớp, xác thực thất bại
		fmt.Println("Đăng nhập thất bại:", err)
	}

	pkg.Run()
}
