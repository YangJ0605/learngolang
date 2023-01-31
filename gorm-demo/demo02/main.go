package main

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   int64          `gorm:"primary_key"`
	Name sql.NullString `gorm:"default:cc"`
	Age  uint8
}

type Student struct {
	gorm.Model
	Class  string
	UserId int64
	User   User
}

func (u User) TableName() string {
	return "user"
}

func main() {
	dsn := "root:root1234@tcp(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("connect sql err: %v\n", err.Error())
		return
	}

	db.AutoMigrate(&User{})
	db.Table("student").AutoMigrate(&Student{})

	// u1 := User{Name: sql.NullString{
	// 	String: "cc",
	// 	Valid:  true,
	// }, Age: 18}

	// db.Create(&u1)

	// u2 := User{Name: sql.NullString{
	// 	String: "ff",
	// 	Valid:  true,
	// }, Age: 20}

	// db.Create(&u2)

	// 查询
	var u User
	// res := db.First(&u)
	// res := db.Select("Name").First(&u)
	res := db.Debug().Omit("Name").First(&u)

	if res.Error != nil {
		fmt.Println("查询失败")
		return
	}
	fmt.Printf("u: %v\n", u)

	// var users []User
	// var count int64
	// db.Debug().Find(&users).Count(&count)
	// fmt.Printf("users: %v\n", users)
	// fmt.Printf("count: %v\n", count)

	// 更新
	// db.Debug().Model(&User{}).Where("name like ?", "%%").Update("age", gorm.Expr("age+?", 2))

	// db.Delete(&u)
}
