package books

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := GetBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	bookId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid book Id", http.StatusBadRequest)
		return
	}
	book, err := GetBookService(bookId)
	if err != nil {
		if err.Error() == "no books found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	var newBook NewBook
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	addedBook, err := AddBookService(newBook)
	if err != nil {
		http.Error(w, "failed to add book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(addedBook)
}

func CreateBookRequestHandler(w http.ResponseWriter, r *http.Request) {
	var req BookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	err := SaveBookRequest(req)
	if err != nil {
		http.Error(w, "failed to save request", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("request submitted"))
}

func GetPendingRequestsHandler(w http.ResponseWriter, r *http.Request) {
	authorId, _ := strconv.Atoi(r.Header.Get("X-Author-ID")) // временно без middleware
	requests, err := GetRequestsByAuthorId(authorId)
	if err != nil {
		http.Error(w, "failed to load requests", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(requests)
}

func ApproveBookRequestHandler(w http.ResponseWriter, r *http.Request) {
	authorId, _ := strconv.Atoi(r.Header.Get("X-Author-ID"))
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = ApproveRequest(id, authorId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	w.Write([]byte("approved"))
}
