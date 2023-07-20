package controller

import (
	"ApiLibrary/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type BookControllerImpl struct {
	BookService service.BookService
}

func NewBookController(BookService service.BookService) BookController {
	return &BookControllerImpl{
		BookService: BookService,
	}
}

func (controller *BookControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, param httprouter.Param) {
	//TODO implement me
	panic("implement me")
}

func (controller *BookControllerImpl) Create(writer http.ResponseWriter, request *http.Request, param httprouter.Param) {
	//TODO implement me
	panic("implement me")
}

func (controller *BookControllerImpl) GetById(writer http.ResponseWriter, request *http.Request, param httprouter.Param) {
	//TODO implement me
	panic("implement me")
}

func (controller *BookControllerImpl) Update(writer http.ResponseWriter, request *http.Request, param httprouter.Param) {
	//TODO implement me
	panic("implement me")
}

func (controller *BookControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Param) {
	//TODO implement me
	panic("implement me")
}
