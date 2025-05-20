package authors

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func GetAuthorsHandler(w http.ResponseWriter, r *http.Request) {
	authors, err := GetAuthors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

func GetAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	authorID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid author Id", http.StatusBadRequest)
		return
	}
	author, err := GetAuthorService(authorID)
	if err != nil {
		if err.Error() == "author is not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}

func AddAuthorHandler(w http.ResponseWriter, r *http.Request) {
	var author Author

	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if author.FullName == "" || author.Specialization == "" {
		http.Error(w, "invalid full_name/specialization in request body", http.StatusBadRequest)
		return
	}

	if err := AddAuthorService(&author); err != nil {
		http.Error(w, "failed to add author", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)
}

func ChangeAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	authorID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid author id", http.StatusBadRequest)
		return
	}

	var updateBody updateAuthorBody
	if err := json.NewDecoder(r.Body).Decode(&updateBody); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if updateBody.FullName == "" || updateBody.Specialization == "" {
		http.Error(w, "invalid full_name/specialization in request body", http.StatusBadRequest)
		return
	}

	if err := UpdateAuthorService(authorID, updateBody.FullName, updateBody.Specialization); err != nil {
		if errors.Is(err, ErrAuthorNotFound) {
			http.Error(w, "author is not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to update author info", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Author updated successfully"))
}
