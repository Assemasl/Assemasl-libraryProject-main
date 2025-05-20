package books

import "errors"

type Book struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	IsbnCode int    `json:"isbnCode"`
	AuthorId int    `json:"authorId"`
}

type NewBook struct {
	Title      string  `json:"title"`
	Genre      string  `json:"genre"`
	IsbnCode   int     `json:"isbnCode"`
	AuthorId   *int    `json:"authorId,omitempty"`
	AuthorName *string `json:"authorName,omitempty"`
}

type BookRequest struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Genre      string `json:"genre"`
	IsbnCode   int    `json:"isbnCode"`
	AuthorName string `json:"authorName"`
	Status     string `json:"status"`
	AuthorId   *int   `json:"authorId,omitempty"`
}

var errNoBooksFound = errors.New("no books found")
