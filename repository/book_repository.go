package repository

import (
	"ApiLibrary/model/domain"
	"context"
	"database/sql"
)

type BookRepository interface {
	GetAll(ctx context.Context, tx sql.Tx) []domain.Book
	Create(ctx context.Context, tx sql.Tx, book domain.Book) domain.Book
	GetById(ctx context.Context, tx sql.Tx, bookId int) domain.Book
	Update(ctx context.Context, tx sql.Tx, book domain.Book) domain.Book
	Delete(ctx context.Context, tx sql.Tx, book domain.Book)
}
