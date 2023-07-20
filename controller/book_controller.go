package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type BookController interface {
	GetAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Create(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	GetById(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}
