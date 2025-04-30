package main

import (
	httpdelivery "comclas/internal/delivery/http"
	mysqlrepo "comclas/internal/repository/mysql"
	"comclas/internal/service"
	"comclas/internal/usecase"
	"os"

	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	// dsn := fmt.Sprintf("%s@%stcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),"root@tcp(127.0.0.1:3306)/pato_donat")
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error: ", err)
		panic("Failed to connect to the database!")
	}

	r := gin.Default()

	// Setup handler
	repo := mysqlrepo.Init(db)
	svc := service.Init(repo)
	roleUC := usecase.Init(svc)
	httpdelivery.InitRoutes(r, roleUC)

	r.Run(":" + os.Getenv("APP_PORT"))
}
