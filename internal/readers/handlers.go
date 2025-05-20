package readers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// getting all readers
func GetReadersHandler(w http.ResponseWriter, r *http.Request) {
	readers, err := GetReadersService()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(readers)

}

// getting reader by id
func GetReaderHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	readerID, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Invalid reader id", http.StatusBadRequest)
		return
	}

	reader, err := GetReaderService(readerID)
	if err != nil {
		if err.Error() == "reader with this id is not found" { //обработку ошибок переписать(?)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reader)

}

// change reader info
func ChangeReaderHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	readerID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid reader id", http.StatusBadRequest)
		return
	}
	var updateBody UpdateBody

	err = json.NewDecoder(r.Body).Decode(&updateBody)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if updateBody.FullName == "" {
		http.Error(w, "invalid full_name in request body", http.StatusBadRequest)
		return
	}

	err = UpdateReaderService(readerID, updateBody.FullName)
	if err != nil {
		if errors.Is(err, ErrReaderNotFound) { //проверка на соответствие ошибке!
			http.Error(w, "reader is not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to update reader info", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reader updated successfully"))
}

// adding new reader
func AddReaderHandler(w http.ResponseWriter, r *http.Request) {

	var reader Reader
	err := json.NewDecoder(r.Body).Decode(&reader)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if reader.FullName == "" {
		http.Error(w, "Invalid full_name in request body", http.StatusBadRequest)
		return
	}

	err = AddReaderService(&reader)
	if err != nil {
		http.Error(w, "Add reader error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reader)

}
