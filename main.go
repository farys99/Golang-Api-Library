package main

import (
	"ApiLibrary/app"
	"ApiLibrary/controller"
	"ApiLibrary/helper"
	"ApiLibrary/repository"
	"ApiLibrary/service"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		helper.PanicIfError(err)
	}

	db := app.NewDB()
	log.Println("Database connected")

	validate := validator.New()
	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepository, db, validate)
	bookController := controller.NewBookController(bookService)
	router := app.NewRouter(bookController)

	log.Println("Starting server...")
	server := http.Server{

		Addr: os.Getenv("APP_URL"),
		//Handler: middleware.NewAuthMiddleware(router),
		Handler: router,
	}
	log.Printf("Server running on %s", os.Getenv("APP_URL"))
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
