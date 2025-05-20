package server

import (
	"net/http"

	"libraryproject/auth"
	"libraryproject/internal/authors"
	"libraryproject/internal/books"
	"libraryproject/internal/readers"
	"libraryproject/utils"
)

func Run() {
	mux := http.NewServeMux()

	// --- Auth Routes ---
	mux.HandleFunc("/register", auth.RegisterHandler)
	mux.HandleFunc("/login", auth.LoginHandler)

	// --- Books Routes ---
	mux.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			id := r.URL.Query().Get("id")
			if id != "" {
				books.GetBookHandler(w, r)
			} else {
				books.GetBooksHandler(w, r)
			}
		} else if r.Method == http.MethodPost {
			books.AddBookHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/books/request", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			books.CreateBookRequestHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/books/pending", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			books.GetPendingRequestsHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/books/approve", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPatch {
			books.ApproveBookRequestHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// --- Authors Routes ---
	mux.HandleFunc("/authors", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			id := r.URL.Query().Get("id")
			if id != "" {
				authors.GetAuthorHandler(w, r)
			} else {
				authors.GetAuthorsHandler(w, r)
			}
		} else if r.Method == http.MethodPost {
			authors.AddAuthorHandler(w, r)
		} else if r.Method == http.MethodPatch {
			authors.ChangeAuthorHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// --- Readers Routes ---
	mux.HandleFunc("/readers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			id := r.URL.Query().Get("id")
			if id != "" {
				readers.GetReaderHandler(w, r)
			} else {
				readers.GetReadersHandler(w, r)
			}
		} else if r.Method == http.MethodPatch {
			readers.ChangeReaderHandler(w, r)
		} else if r.Method == http.MethodPost {
			readers.AddReaderHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	utils.InfoLog.Println("server is running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		utils.ErrorLog.Fatalf("Error starting server: %v", err)
	}
}
