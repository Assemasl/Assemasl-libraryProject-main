package authors

import (
	"database/sql"
	"errors"

	"libraryproject/database"
)

var ErrAuthorNotFound = errors.New("author is not found")

func GetAllAuthors() ([]Author, error) {
	query := `SELECT * from authors ORDER BY id`
	rows, err := database.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []Author

	for rows.Next() {

		var author Author
		if err := rows.Scan(&author.Id, &author.FullName, &author.Specialization); err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return authors, nil
}

func GetAuthorById(authorId int) (Author, error) {
	query := `SELECT id, full_name, specialization from authors WHERE id=$1`
	var author Author

	err := database.Db.QueryRow(query, authorId).Scan(&author.Id, &author.FullName, &author.Specialization)
	if err != nil {
		if err == sql.ErrNoRows {
			return Author{}, ErrAuthorNotFound
		}
		return Author{}, err
	}
	return author, err
}

func AddAuthor(author *Author) error {
	query := `INSERT INTO authors (full_name,specialization) VALUES($1,$2) RETURNING id`

	err := database.Db.QueryRow(query, author.FullName, author.Specialization).Scan(&author.Id)
	if err != nil {
		return errors.New("inserting author error")
	}
	return nil
}

func ChangeAuthorById(authorId int, fullname string, specialization string) error {
	query := `UPDATE authors SET full_name=$1,specialization=$2 WHERE id=$3`
	result, err := database.Db.Exec(query, fullname, specialization, authorId)
	if err != nil {
		return errors.New("failed to execute update query")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("failed to retrieve affected rows count")
	}

	if rowsAffected == 0 {
		return ErrAuthorNotFound
	}

	return nil

}
