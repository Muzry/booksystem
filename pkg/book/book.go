package book

import (
	"booksystem/pkg/database"
	"fmt"
)

const (
	Create = "create"
	Update = "update"
)

func GetBookInfo(bookID string) (*database.Book, error) {

	book := database.Book{}
	db, err := database.GetConnection()
	defer db.Close()
	if err != nil {
		return nil, err
	}

	has, err := db.ID(bookID).Get(&book)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("no such result from book table.")
	}
	return &book, nil
}

func DeleteBook(bookID string) error {
	book := database.Book{}
	db, err := database.GetConnection()
	defer db.Close()

	if err != nil {
		return err
	}

	_, err = db.ID(bookID).Delete(book)
	if err != nil {
		return err
	}
	return nil
}

func GetBooks(limit, start int) (*[]database.Book, error) {
	books := make([]database.Book, 0)
	db, err := database.GetConnection()
	defer db.Close()

	if err != nil {
		return nil, err
	}

	err = db.Limit(limit, start).Find(&books)

	if err != nil {
		return nil, err
	}
	return &books, nil
}

func CreateORUpdateBook(book database.Book, mothodType, bookID string) (*database.Book, error) {
	db, err := database.GetConnection()
	defer db.Close()

	if err != nil {
		return nil, err
	}

	if mothodType == Create {
		book.ID = bookID
		_, err = db.Insert(book)
		if err != nil {
			return nil, err
		}
	} else if mothodType == Update {
		_, err = db.ID(bookID).Update(book)
		if err != nil {
			return nil, err
		}
	}
	return &book, nil
}
