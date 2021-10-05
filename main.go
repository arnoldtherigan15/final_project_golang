package main

import (
	"fmt"
	"log"
	"os"

	_todoHandler "github.com/arnoldtherigan15/final_project_golang/todo/handler"
	_todoRepo "github.com/arnoldtherigan15/final_project_golang/todo/repository"
	_todoService "github.com/arnoldtherigan15/final_project_golang/todo/service"

	_userHandler "github.com/arnoldtherigan15/final_project_golang/user/handler"
	_userRepo "github.com/arnoldtherigan15/final_project_golang/user/repository"
	_userService "github.com/arnoldtherigan15/final_project_golang/user/service"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file not found\n", err)
	}
}

func main() {
	DB_HOST := os.Getenv("DB_HOST")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Connection Database Error\n", err.Error())
	}

	todoRepo := _todoRepo.NewRepository(db)
	todoService := _todoService.NewService(todoRepo)

	userRepo := _userRepo.NewRepository(db)
	userService := _userService.NewService(userRepo)

	e := echo.New()
	g := e.Group("/api/v1")

	_todoHandler.NewHandler(g, todoService)
	_userHandler.NewHandler(g, userService)

	PORT := os.Getenv("SERVER_PORT")
	log.Fatal(e.Start(PORT))
}
