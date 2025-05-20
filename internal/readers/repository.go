package readers

import (
	"database/sql"
	"errors"
	"fmt"

	"libraryproject/database"
)

var ErrReaderNotFound = errors.New("reader is not found")

func GetAllReaders() ([]Reader, error) {
	query := `SELECT id,full_name from readers ORDER BY id`
	rows, err := database.Db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var readers []Reader

	for rows.Next() {
		var reader Reader
		if err = rows.Scan(&reader.Id, &reader.FullName); err != nil {
			return nil, err
		}
		readers = append(readers, reader)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return readers, nil
}

func GetReaderById(readerId int) (Reader, error) {

	var reader Reader
	query := `SELECT id,full_name FROM readers WHERE id=$1 `
	err := database.Db.QueryRow(query, readerId).Scan(&reader.Id, &reader.FullName)
	if err != nil {
		if err == sql.ErrNoRows {
			return Reader{}, ErrReaderNotFound
		}
		return Reader{}, err
	}

	return reader, nil
}

func ChangeReaderById(readerId int, FullName string) error {
	query := `UPDATE readers SET full_name=$1 WHERE id=$2`
	result, err := database.Db.Exec(query, FullName, readerId) //exec для insert, update,delete
	if err != nil {
		return errors.New("failed to execute update query")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("failed to retrieve affected rows count")
	}

	if rowsAffected == 0 {
		return ErrReaderNotFound
	}

	return nil

}

// adding new reader
func AddReader(reader *Reader) error {
	query := `INSERT INTO readers (full_name) VALUES ($1) RETURNING id`
	err := database.Db.QueryRow(query, reader.FullName).Scan(&reader.Id)
	if err != nil {
		fmt.Println(err)
		return errors.New("inserting reader error")
	}
	return nil
}

//delete reader
