package imp

import (
	"awesomeProject1/internal/entities/user"
	"awesomeProject1/internal/storage"
	"awesomeProject1/internal/utils"
	"context"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

var jwtSecret = []byte("supersecretkey")

func NewAuthService(repo storage.UserRepository) *AuthServiceImp {
	return &AuthServiceImp{repo: repo}
}

type AuthServiceImp struct {
	repo storage.UserRepository
}

func (a AuthServiceImp) Register(ctx context.Context, username string, password string) error {
	user := user.UserDB{
		Username: username,
		Password: password,
	}
	err := a.repo.Save(ctx, &user)
	if err != nil {
		return err
	}
	return nil
}

func (a AuthServiceImp) Login(ctx context.Context, username string, password string) (string, error) {

	user, err := a.repo.Get(ctx, username)

	if user == nil && err == nil {
		err := a.Register(ctx, username, password)
		if err != nil {
			return "", err
		}
	}

	user, err = a.repo.Get(ctx, username)

	if err != nil {
		return "", err
	}

	if user.Password != password {
		return "", errors.New("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString(jwtSecret)
}

func (a AuthServiceImp) ValidateToken(ctx context.Context, tokenString string) (string, error) {

	if tokenString == "" {
		return "", errors.New("token is missing")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return "", fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return "", errors.New("token has expired")
			}
		}

		username, ok := claims["username"].(string)
		if !ok || username == "" {
			return "", errors.New("username not found in token")
		}

		return username, nil
	}

	return "", errors.New("invalid token")
}

func JWTMiddleware(authService *AuthServiceImp) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/api/auth" {
				next.ServeHTTP(w, r)
				return
			}

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				utils.SendErrorResponse(w, http.StatusUnauthorized, "Missing or invalid Authorization header")
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			username, err := authService.ValidateToken(context.Background(), tokenString)
			if err != nil {
				utils.SendErrorResponse(w, http.StatusUnauthorized, "Invalid token")
				return
			}

			ctx := context.WithValue(r.Context(), "username", username)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
