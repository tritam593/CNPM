package models

import (

	"strings"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Addresses     string
	CartID        *Cart
	FirstName     string `gorm:"size:100;not null"`
	LastName      string `gorm:"size:100;not null"`
	Email         string `gorm:"size:100;not null;uniqueIndex"`
	Password      string `gorm:"size:255;not null"`
	RememberToken string `gorm:"size:255;not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

func (u *User) FindByEmail(db *gorm.DB, email string) (*User, error) {
	var err error
	var user User
	// fmt.Println(email + "aaaaaaa")
	err = db.Debug().Model(User{}).Where("LOWER(email) = ?", strings.ToLower(email)).
		First(&user).
		Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) FindByID(db *gorm.DB, userID string) (*User, error) {
	var err error
	var user User

	err = db.Debug().Model(User{}).Where("id = ?", userID).
		First(&user).
		Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUser creates a new user in the database using the provided User parameter.
//
// Parameters:
//   - db: A pointer to the Gorm database instance.
//   - param: A pointer to the User struct containing user details.
//
// Returns:
//   - *User: A pointer to the newly created User instance.
//   - error: An error, if any, encountered during the database operation.
//
// Example:
//   user := &User{
//       ID:        1,
//       FirstName: "John",
//       LastName:  "Doe",
//       Email:     "john.doe@example.com",
//       Password:  "hashed_password",
//   }
//   createdUser, err := CreateUser(databaseInstance, user)
//   if err != nil {
//       log.Fatal(err)
//   }
//   fmt.Printf("User created with ID: %d\n", createdUser.ID)
func (u *User) CreateUser(db *gorm.DB, param *User) (*User, error) {
	user := &User{
		ID:        param.ID,
		FirstName: param.FirstName,
		Addresses: param.Addresses,
		LastName:  param.LastName,
		Email:     param.Email,
		Password:  param.Password,
	}

	err := db.Debug().Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}