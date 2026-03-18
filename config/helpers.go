package paymentprocessor

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func generateToken(userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func authenticateUser(c *gin.Context, token string) (int64, error) {
	_, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func generateUUID() (string, error) {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(uuid), nil
}

func hashPassword(password string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func validateEmail(email string) bool {
	if strings.Contains(email, "@") && strings.Contains(email, ".") {
		return true
	}
	return false
}

func sendEmail(to string, subject string, body string) error {
	return http.Error(nil, fmt.Sprintf("Email not sent to %s", to))
}

func sendCustomEmail(to string, subject string, body string, emailBody string) error {
	return http.Error(nil, fmt.Sprintf("Email not sent to %s", to))
}

func sendCustomEmailWithFile(to string, subject string, body string, filePath string) error {
	return http.Error(nil, fmt.Sprintf("File not found: %s", filePath))
}