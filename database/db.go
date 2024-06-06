package database

import (
	
	"login/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBconect() { //database connecting and table creation

	dsn := os.Getenv("DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("=========FAILED TO CONNECT DATABASE=========")
	}
	DB = db

	DB.AutoMigrate(&model.UserModel{}, &model.AdminModel{})
	

}
