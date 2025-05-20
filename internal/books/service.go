package books

import (
	"errors"
	"fmt"

	"libraryproject/database"
	"libraryproject/internal/authors"
)

func GetBooks() ([]Book, error) {
	return GetAllBooks()
}

func GetBookService(bookId int) (Book, error) {
	return GetBookById(bookId)
}

func AddBookService(newBook NewBook) (Book, error) {
	var authorId int

	if newBook.AuthorId != nil {
		_, err := authors.GetAuthorService(*newBook.AuthorId)
		if err != nil {
			if errors.Is(err, authors.ErrAuthorNotFound) {
				return Book{}, errors.New("author with is not found")
			}
			return Book{}, errors.New("failed to check author")
		}
		authorId = *newBook.AuthorId
	} else if newBook.AuthorName != nil {
		newAuthor := authors.Author{
			FullName:       *newBook.AuthorName,
			Specialization: "no info",
		}
		err := authors.AddAuthorService(&newAuthor)
		if err != nil {
			return Book{}, errors.New("failed to add author")
		}
		authorId = newAuthor.Id
	} else {
		return Book{}, errors.New("either authorId or authorName must be provided")
	}

	book := Book{
		Title:    newBook.Title,
		Genre:    newBook.Genre,
		IsbnCode: newBook.IsbnCode,
		AuthorId: authorId,
	}

	createdBook, err := AddBook(book)
	if err != nil {
		fmt.Print(err)
		return Book{}, errors.New("failed to add book")
	}

	return createdBook, nil
}

func SaveBookRequest(req BookRequest) error {
	query := `INSERT INTO book_requests (title, genre, isbn_code, author_name, status) VALUES ($1, $2, $3, $4, 'pending')`
	_, err := database.Db.Exec(query, req.Title, req.Genre, req.IsbnCode, req.AuthorName)
	return err
}

func GetRequestsByAuthorId(authorId int) ([]BookRequest, error) {
	query := `SELECT id, title, genre, isbn_code, author_name, status FROM book_requests WHERE author_id = $1 AND status = 'pending'`
	rows, err := database.Db.Query(query, authorId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []BookRequest
	for rows.Next() {
		var r BookRequest
		if err := rows.Scan(&r.Id, &r.Title, &r.Genre, &r.IsbnCode, &r.AuthorName, &r.Status); err != nil {
			return nil, err
		}
		r.AuthorId = &authorId
		result = append(result, r)
	}
	return result, nil
}

func ApproveRequest(id int, authorId int) error {
	tx, err := database.Db.Begin()
	if err != nil {
		return err
	}

	var req BookRequest
	err = tx.QueryRow(`SELECT title, genre, isbn_code, author_name FROM book_requests WHERE id=$1 AND status='pending'`, id).
		Scan(&req.Title, &req.Genre, &req.IsbnCode, &req.AuthorName)
	if err != nil {
		tx.Rollback()
		return errors.New("request not found or already processed")
	}

	_, err = tx.Exec(`INSERT INTO books (title, genre, isbn_code, author_id) VALUES ($1, $2, $3, $4)`, req.Title, req.Genre, req.IsbnCode, authorId)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`UPDATE book_requests SET status='approved', author_id=$1 WHERE id=$2`, authorId, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
