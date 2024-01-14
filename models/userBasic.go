package models

import (
	"errors"
	"fmt"
	"goweb/utils"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Email     string `valid:"email"`
	Password  string
	Salt      string
	IsVip     bool
	Token     string
	Address   string `valid:"matches(^0x[a-fA-F0-9]{40}$)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (table *User) TableName() string {
	return "user"
}
func FindUserByEmail(email string) User {
	user := User{}
	utils.DB.Where("email = ?", email).First(&user)
	return user
}
func FindUserByEmailAndPwd(email string, password string) User {
	user := User{}
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.Md5Encode(str)
	fmt.Println("---------------------------------------", temp)
	utils.DB.Where("email = ? and password = ?", email, password).First(&user)
	utils.DB.Model(&user).Where("email = ?", user.Email).Update("token", temp)
	return user
}

func FindUserByAddress(address string) User {
	user := User{}
	utils.DB.Where("address = ?", &user.Address).First(&user)
	return user
}
func CreateUser(user User) *gorm.DB {
	return utils.DB.Create(&user)
}

func BindAddress(user User, address string) *gorm.DB {
	return utils.DB.Model(&user).Where("email = ?", user.Email).Update("address", address)
}

func DeleteUser(user User) *gorm.DB {
	return utils.DB.Delete(&user)
}

func UpdateUser(user User) *gorm.DB {
	return utils.DB.Model(&user).Updates(User{ID: user.ID, Email: user.Email, Address: user.Address, Password: user.Password})
}

func CheckUserExist(email string) (error error) {
	var user User
	utils.DB.Model(&user).Where("email = ?", email).Find(&user)
	if user.Salt != "" {
		return errors.New("ErrorUserExit")
	}
	return
}
