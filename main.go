package main

import (
	"ApiLibrary/app"
	"ApiLibrary/controller"
	"ApiLibrary/helper"
	"ApiLibrary/middleware"
	"ApiLibrary/repository"
	"ApiLibrary/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {

	log.Println("Server starting ...")
	db := app.NewDB()
	log.Println("Database Connected ...")
	validate := validator.New()
	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepository, db, validate)
	bookController := controller.NewBookController(bookService)
	router := app.NewRouter(bookController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	log.Println("Server started ...")

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
