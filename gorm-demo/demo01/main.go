package main

import (
	"database/sql"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// UserInfo 用户信息
// type UserInfo struct {
// 	ID     uint `gorm:"primary_key"`
// 	Name   string
// 	Gender string
// 	Hobby  string
// }

type UserInfo struct {
	gorm.Model
	Name         string `gorm:"column:user_name"`
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`
	MemberNumber *string `gorm:"unique;not null"`
	Num          int     `gorm:"AUTO_INCREMENT"`
	Address      string  `gorm:"index:addr"`
	IgnoreMe     int     `gorm:"-"`
}

type Product struct {
	gorm.Model
	Name *string `gorm:"default:cc"`
}

// 重命名表名
func (u UserInfo) TableName() string {
	return "user"
}

func main() {
	dsn := "root:root1234@tcp(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&UserInfo{})

	// 也可以这样指定表名 优先级更高
	db.Table("user2").AutoMigrate(&UserInfo{})

	db.Table("product").AutoMigrate(&Product{})

	// 创建一条记录
	// db.Create(&Product{Code: "110", Price: 100})

	// var p Product
	// result := db.First(&p, 1)
	// fmt.Printf("p1: %v\n", p)
	// result := db.First(&p, "code = ?", "110")
	// if result.Error != nil {
	// 	fmt.Println("查询失败")
	// } else {
	// 	fmt.Printf("p2: %v\n", p)
	// }
	// db.Model(&p).Update("price", 800)

	// db.Model(&p).Updates(Product{Price: 900, Code: "101"})

	// db.Model(&p).Updates(map[string]interface{}{"price": 800, "code": "1010"})

	// db.Delete(&p, 1)

	// u1 := UserInfo{1, "cc", "男", "run"}
	// u2 := UserInfo{2, "ff", "女", "run2"}
	// db.Create(&u1)
	// db.Create(&u2)

	// u := new(UserInfo)
	// db.First(&u)
	// fmt.Printf("u: %#v\n", u)

	// var uu UserInfo
	// db.Find(&uu, "Gender=?", "女")
	// fmt.Printf("uu: %#v\n", uu)

	// db.Model(&uu).Update("hobby", "kkk")

	// db.Delete(&uu)

	// db.Create(&UserInfo{MemberNumber: new(string)})
	// 通过指针解决
	name := "ccff"
	db.Table("product").Create(&Product{Name: &name}) // 如果tag中设置了默认值，如果设置零值是不会存入零值的，只会填入默认值
}
