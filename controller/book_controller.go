package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type BookController interface {
	GetAll(writer http.ResponseWriter, request *http.Request, param httprouter.Param)
	Create(writer http.ResponseWriter, request *http.Request, param httprouter.Param)
	GetById(writer http.ResponseWriter, request *http.Request, param httprouter.Param)
	Update(writer http.ResponseWriter, request *http.Request, param httprouter.Param)
	Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Param)
}
