package repository

import (
	"ApiLibrary/helper"
	"ApiLibrary/model/domain"
	"context"
	"database/sql"
	"errors"
)

type BookRepositoryImpl struct {
}

func NewBookRepository() BookRepository {
	/*belum di proses*/
	return nil
}

func (repository *BookRepositoryImpl) GetAll(ctx context.Context, tx sql.Tx) []domain.Book {
	SQL := "select id, title, years, publisher from book"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var books []domain.Book
	for rows.Next() {
		book := domain.Book{}
		err := rows.Scan(&book.Id, &book.Title, &book.Years, &book.Publisher)
		helper.PanicIfError(err)
		books = append(books, book)
	}
	return books
}

func (repository *BookRepositoryImpl) Create(ctx context.Context, tx sql.Tx, book domain.Book) domain.Book {

	SQL := "insert into book(title, years, publisher) values(?)"
	result, err := tx.ExecContext(ctx, SQL, book.Title, book.Years, book.Publisher)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	book.Id = int(id)
	return book
}

func (repository *BookRepositoryImpl) GetById(ctx context.Context, tx sql.Tx, bookId int) (domain.Book, error) {
	SQL := "select id, title, years, publisher from book where id =?"
	rows, err := tx.QueryContext(ctx, SQL, bookId)
	helper.PanicIfError(err)
	defer rows.Close()

	books := domain.Book{}
	if rows.Next() {
		err := rows.Scan(&books.Id, &books.Title, &books.Years, &books.Publisher)
		helper.PanicIfError(err)
		return books, nil
	} else {
		return books, errors.New("book is not found")
	}
}

func (repository *BookRepositoryImpl) Update(ctx context.Context, tx sql.Tx, book domain.Book) domain.Book {
	//TODO implement me
	panic("implement me")
}

func (repository *BookRepositoryImpl) Delete(ctx context.Context, tx sql.Tx, book domain.Book) {
	//TODO implement me
	panic("implement me")
}
