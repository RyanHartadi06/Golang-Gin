package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"
	"pustaka-api/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/go-library?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Successfully connected to database")

	db.AutoMigrate(&book.Book{}, &user.User{})

	bookRepository := book.NewRepositoryBook(db)
	userRepository := user.NewRepositoryUser(db)

	bookService := book.NewService(bookRepository)
	userService := user.NewService(userRepository)

	bookHandler := handler.NewBookHandler(bookService)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", bookHandler.FindAll)
	v1.GET("/:id", bookHandler.FindID)
	v1.POST("/postBook", bookHandler.PostBooksHandler)

	v1.POST("/user", userHandler.RegisterUserHandler)

	router.Run()

}
