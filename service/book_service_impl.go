package service

import (
	"ApiLibrary/exception"
	"ApiLibrary/helper"
	"ApiLibrary/model/domain"
	"ApiLibrary/model/web"
	"ApiLibrary/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewBookService(BookRepository repository.BookRepository, DB *sql.DB, validate *validator.Validate) BookService {
	return &BookServiceImpl{
		BookRepository: BookRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *BookServiceImpl) GetAll(ctx context.Context) []web.BookResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	books := service.BookRepository.GetAll(ctx, tx)
	return helper.ToBookResponses(books)
}

func (service *BookServiceImpl) Create(ctx context.Context, request web.BookCreateRequest) web.BookResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book := domain.Book{
		Title:     request.Title,
		Years:     request.Years,
		Publisher: request.Publisher,
	}
	book = service.BookRepository.Create(ctx, tx, book)

	return helper.ToBookResponse(book)
}

func (service *BookServiceImpl) GetById(ctx context.Context, bookId int) web.BookResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.GetById(ctx, tx, bookId)

	if err != nil {
		panic(exception.NotFoundError(err.Error()))
	}

	return helper.ToBookResponse(book)
}

func (service *BookServiceImpl) Update(ctx context.Context, request web.BookUpdateRequest) web.BookResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.GetById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NotFoundError(err.Error()))
	}

	book.Title = request.Title
	book.Years = request.Years
	book.Publisher = request.Publisher

	book = service.BookRepository.Update(ctx, tx, book)

	return helper.ToBookResponse(book)
}

func (service *BookServiceImpl) Delete(ctx context.Context, bookId int) {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.GetById(ctx, tx, bookId)
	if err != nil {
		panic(exception.NotFoundError(err.Error()))
	}

	service.BookRepository.Delete(ctx, tx, book)
}
