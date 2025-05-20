package books

import (
	"database/sql"

	"libraryproject/database"
)

func GetAllBooks() ([]Book, error) {
	query := `SELECT * FROM books`
	rows, err := database.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.Id, &book.Title, &book.Genre, &book.IsbnCode, &book.AuthorId); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

func GetBookById(bookId int) (Book, error) {
	query := `SELECT id,title,genre,isbn_code,author_id FROM books WHERE id=$1`
	var book Book
	err := database.Db.QueryRow(query, bookId).Scan(&book.Id, &book.Title, &book.Genre, &book.IsbnCode, &book.AuthorId)
	if err != nil {
		if err == sql.ErrNoRows {
			return Book{}, errNoBooksFound
		}
		return Book{}, err
	}
	return book, nil
}

func AddBook(book Book) (Book, error) {
	query := `INSERT INTO books (title,genre,isbn_code,author_id) VALUES ($1,$2,$3,$4) RETURNING id`
	err := database.Db.QueryRow(query, book.Title, book.Genre, book.IsbnCode, book.AuthorId).Scan(&book.Id)
	if err != nil {
		return Book{}, err
	}
	return book, nil
}
