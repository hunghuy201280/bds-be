package main

import (
	"bds-service/common/l"
	"bds-service/component/tokenprovider/jwt"
	"bds-service/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	ll = l.New()
)

func main() {
	config := configs.LoadConfig()

	//dsn := os.Getenv("DB_URL")
	//jwtSecret := os.Getenv("JWT_SECRET")
	//port := os.Getenv("PORT")
	dsn := config.DBUrl
	jwtSecret := config.JWTSecret
	port := config.Port
	tokenProvider := jwt.NewTokenJWTProvider(jwtSecret)
	refreshTokenProvider := jwt.NewTokenJWTProvider(config.JWTRefreshSecret)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		ll.Fatal("err when opening db connection", l.Error(err))
	}

	if err := RunService(db, port, tokenProvider, refreshTokenProvider); err != nil {
		ll.Fatal("err when starting service", l.Error(err))
	}

}

//root:@tcp(127.0.0.1:3306)/food-delivery?charset=utf8mb4&parseTime=True&loc=Local
