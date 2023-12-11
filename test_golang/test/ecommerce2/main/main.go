package main

import (
	// "fmt"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
	"test/ecommerce2/pkg"
	// "encoding/json"
)

// User là một mô hình đơn giản để minh họa preload
type User struct {
	ID    uint
	Name  string
	Email string
	Posts []Post `gorm:"foreignKey:UserID"` // preload sẽ tự động nạp các bài viết của người dùng này
}

// Post là một bài viết của người dùng
type Post struct {
	ID     uint
	Title  string
	UserID uint
}

func main() {
	// Mở kết nối với MySQL
	// dsn := "tam:1@tcp(127.0.0.1:3306)/temp?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic("Could not connect to the database")
	// }

	// Auto Migrate các bảng
	// db.AutoMigrate(&User{}, &Post{})

	// // Tạo 10 người dùng với mỗi người dùng có 3 bài viết
	// for i := 1; i <= 10; i++ {
	// 	user := User{
	// 		Name:  fmt.Sprintf("User %d", i),
	// 		Email: fmt.Sprintf("user%d@example.com", i),
	// 	}

	// // 	// Tạo người dùng
	// 	db.Create(&user)

	// 	// Tạo 3 bài viết cho người dùng
	// 	for j := 1; j <= 3; j++ {
	// 		post := Post{
	// 			Title:  fmt.Sprintf("Post %d-%d", i, j),
	// 			UserID: user.ID,
	// 		}
	// 		db.Create(&post)
	// 	}
	// }
	// temp := User{
	// 	ID: 100,
	// 	Name: "hahahaha",
	// 	Email: "a@a.aaaa",
	// 	Posts: []Post{
	// 		{Title: "Post 100-1", ID: 1001},
	// 		{Title: "Post 100-2", ID: 1002},
	// 		{Title: "Post 100-3", ID: 1003},
	// 	},
	// }
	// db.Create(&temp)
	// Query một người dùng và in thông tin người dùng và các bài viết preload
	// var retrievedUser User
	// if err := db.Preload("Posts").Find(&retrievedUser).Error; err != nil {
	// 	panic("Could not retrieve user")
	// }

	// fmt.Println("User:", retrievedUser.Name, "Email:", retrievedUser.Email)
	// fmt.Println("Posts:")
	// for _, post := range retrievedUser.Posts {
	// 	fmt.Println(" -", post.Title)
	// }
	// jsonData, err := json.MarshalIndent(retrievedUser, "", "    ")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// In ra chuỗi JSON
	// fmt.Println(string(jsonData))

	// for _, u := range retrievedUser {
	// 	fmt.Println(u)
	// }

	pkg.Run()
}
