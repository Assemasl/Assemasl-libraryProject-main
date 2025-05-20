package auth

import (
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"libraryproject/database"
)

var jwtKey = []byte("supersecretkey")

func RegisterUser(req RegisterRequest) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO author_users (username, password_hash, author_id) VALUES ($1, $2, $3)`
	_, err = database.Db.Exec(query, req.Username, string(hash), req.AuthorID)
	return err
}

func AuthenticateUser(req LoginRequest) (string, error) {
	var user AuthUser
	query := `SELECT id, username, password_hash, author_id FROM author_users WHERE username = $1`
	err := database.Db.QueryRow(query, req.Username).Scan(&user.ID, &user.Username, &user.Password, &user.AuthorID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("user not found")
		}
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.New("invalid password")
	}

	claims := &Claims{
		Username: user.Username,
		AuthorID: user.AuthorID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
