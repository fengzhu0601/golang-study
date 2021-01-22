package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"time"
	"fmt"
)

func main()  {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:huang@feng@wu!@tcp(127.0.0.1:3306)/test?charset=utf8",
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	//// 创建表
	//db.AutoMigrate(&Product{})
	//
	//// 插入数据
	//db.Create(&Product{Code:"D42", Price:100})
	//db.Create(&Product{Code:"D43", Price:200})
	//
	//// 读取数据
	//var product Product
	////db.First(&product, 1)
	////fmt.Println("xxxxxxxxxx", product)
	//db.First(&product, "code = ?", "D43")
	//fmt.Println("xxxxxxxxxx", product.Price, product.Code)
	//
	//db.Model(&product).Update("Price", 300)
	//fmt.Println("xxxxxxxxxx", product.Price, product.Code)
	//db.Model(&product).Updates(Product{Price:500, Code:"F42"})
	//fmt.Println("xxxxxxxxxx", product.Price, product.Code)
	//
	//db.Table("products").Select("code", "price").Where("code = ?", "D43").Scan(&product)
	//fmt.Println(product)
	//
	//// delete
	//db.Delete(&product, 1)

	var a string
	n := db.Model(&Product{}).Limit(10).Find(&ApiProduct{})
	n.Row().Scan(&a)

	fmt.Println(a)
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type ApiProduct struct {
	Code string
}

