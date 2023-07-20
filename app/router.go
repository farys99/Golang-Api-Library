package app

import (
	"ApiLibrary/controller"
	"ApiLibrary/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(bookController controller.BookController) *httprouter.Router {
	router := httprouter.New()

	router.GET("api/library", bookController.GetAll)
	router.GET("api/library/:libraryId", bookController.GetById)
	router.POST("api/library", bookController.Create)
	router.PUT("api/library/:libraryId", bookController.Update)
	router.DELETE("api/library/:libraryId", bookController.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router
}
