package database

import (
	"github.com/JustSteveKing/example-go-api/pkg/kernel"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(app *kernel.Application) *gorm.DB {
	connectionString := "root:root@tcp(localhost:3306)/spicy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		app.Logger.Fatal(err.Error())
		panic(err)
	}

	return db
}
