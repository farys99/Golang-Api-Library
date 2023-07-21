package controller

import (
	"ApiLibrary/helper"
	"ApiLibrary/model/web"
	"ApiLibrary/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type BookControllerImpl struct {
	BookService service.BookService
}

func NewBookController(BookService service.BookService) BookController {
	return &BookControllerImpl{
		BookService: BookService,
	}
}

func (controller *BookControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {

	bookResponse := controller.BookService.GetAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookControllerImpl) Create(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	bookCreateRequest := web.BookCreateRequest{}
	helper.ReadFromRequestBody(request, &bookCreateRequest)

	bookResponse := controller.BookService.Create(request.Context(), bookCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookControllerImpl) GetById(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	bookId := param.ByName("libraryId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	bookResponse := controller.BookService.GetById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookControllerImpl) Update(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	bookUpdateRequest := web.BookUpdateRequest{}
	helper.ReadFromRequestBody(request, &bookUpdateRequest)

	bookId := param.ByName("libraryId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	bookUpdateRequest.Id = id

	bookResponse := controller.BookService.Update(request.Context(), bookUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	bookId := param.ByName("libraryId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	controller.BookService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}
