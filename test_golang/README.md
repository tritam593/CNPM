```
cd $name_folder
go mod init $name_folder
```

Install mux package

```
go get -u github.com/gorilla/mux

```

The Difference betwenn localhost and 127.0.0.1

[Link Here](https://www.tutorialspoint.com/difference-between-localhost-and-127-0-0-1#:~:text=The%20most%20significant%20difference%20between,look%20up%20a%20table%20somewhere.)

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

```
ioutil.ReadAll -> io.ReadAll
ioutil.ReadFile -> os.ReadFile
ioutil.ReadDir -> os.ReadDir
// others
ioutil.NopCloser -> io.NopCloser
ioutil.TempDir -> os.MkdirTemp
ioutil.TempFile -> os.CreateTemp
ioutil.WriteFile -> os.WriteFile
```


```
GRANT ALL PRIVILEGES ON *.* TO 'sammy'@'localhost' WITH GRANT OPTION;
```

```
package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	dsn := "tam:1@tcp(127.0.0.1:3306)/temp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}

	// Auto Migrate các bảng
	db.AutoMigrate(&User{}, &Post{})

	// Tạo 10 người dùng với mỗi người dùng có 3 bài viết
	// for i := 1; i <= 10; i++ {
	// 	user := User{
	// 		Name:  fmt.Sprintf("User %d", i),
	// 		Email: fmt.Sprintf("user%d@example.com", i),
	// 	}

	// 	// Tạo người dùng
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

	// Query một người dùng và in thông tin người dùng và các bài viết preload
	var retrievedUser []User
	// preload similar inner join
	if err := db.Preload("Posts").Find(&retrievedUser).Error; err != nil {
		panic("Could not retrieve user")
	}

	// fmt.Println("User:", retrievedUser.Name, "Email:", retrievedUser.Email)
	// fmt.Println("Posts:")
	// for _, post := range retrievedUser.Posts {
	// 	fmt.Println(" -", post.Title)
	// }

	fmt.Println(retrievedUser)
}

```

# create data for two table in one
```
temp := User{
	ID: 100,
	Name: "hahahaha",
	Email: "a@a.aaaa",
	Posts: []Post{
		{Title: "Post 100-1", ID: 1001},
		{Title: "Post 100-2", ID: 1002},
		{Title: "Post 100-3", ID: 1003},
	},
}
db.Create(&temp)
```

GORM sử dụng quy ước và quy tắc mặc định để xác định mối quan hệ giữa các bảng trong trường hợp của bạn, ngay cả khi bạn không sử dụng câu lệnh JOIN rõ ràng. GORM có thể tự động nạp dữ liệu từ các bảng liên quan thông qua mối quan hệ giữa chúng, giúp giảm sự phức tạp của mã và giúp tăng hiệu suất truy vấn.

Trong trường hợp của bạn, khi bạn khai báo một mối quan hệ giữa User và Post trong struct User như sau:

```
type User struct {
	ID    uint
	Name  string
	Email string
	Posts []Post `gorm:"foreignKey:UserID"` // preload sẽ tự động nạp các bài viết của người dùng này
}
```
GORM sẽ hiểu rằng có một mối quan hệ giữa User và Post, và nó sẽ tự động sử dụng UserID trong bảng Post như là khóa ngoại (foreign key) để liên kết với trường ID của bảng User. Khi bạn sử dụng preload, GORM tự động thực hiện các truy vấn cần thiết để nạp dữ liệu từ bảng Post và đảm bảo rằng nó được liên kết đúng với người dùng tương ứng.

Tóm lại, GORM thường sử dụng các quy tắc mặc định và quy ước để hiểu mối quan hệ giữa các bảng, giảm độ phức tạp của mã và cung cấp một trải nghiệm thuận tiện khi làm việc với cơ sở dữ liệu.


```
kubectl wait pod \
--all \
--for=condition=Ready \
--namespace=${ns}
```