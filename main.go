package main

import (
	"ApiLibrary/app"
	"ApiLibrary/controller"
	"ApiLibrary/helper"
	"ApiLibrary/repository"
	"ApiLibrary/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepository, db, validate)
	bookController := controller.NewBookController(bookService)
	router := app.NewRouter(bookController)

	server := http.Server{

		Addr: "localhost:3000",
		//Handler: middleware.NewAuthMiddleware(router),
		Handler: router,
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
