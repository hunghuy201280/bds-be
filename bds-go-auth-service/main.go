package main

import (
	"bds-go-auth-service/common/l"
	"bds-go-auth-service/component/tokenprovider/jwt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var (
	ll = l.New()
)

func main() {
	dsn := os.Getenv("DB_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	port := os.Getenv("PORT")
	tokenProvider := jwt.NewTokenJWTProvider(jwtSecret)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		ll.Fatal("err when opening db connection", l.Error(err))
	}

	if err := RunService(db, port, tokenProvider); err != nil {
		ll.Fatal("err when starting service", l.Error(err))
	}

}

//root:@tcp(127.0.0.1:3306)/food-delivery?charset=utf8mb4&parseTime=True&loc=Local
