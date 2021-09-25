package configs

import (
	"bmg/models/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	dsn := "host=fullstack-postgres user=postgres password=Admin1234., dbname=tes_bmg port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	var error error
	DB, error = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		panic("Database failed connection : " + error.Error())
	}
	Migration()
}

func Migration()  {
	DB.AutoMigrate(&user.User{})
}