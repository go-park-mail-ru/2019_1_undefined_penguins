package auth

import (
	db "2019_1_undefined_penguins/internal/pkg/database"
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"2019_1_undefined_penguins/internal/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
	"time"
)

var SECRET = []byte("myawesomesecret")

type AuthManager struct {
	mu       sync.RWMutex
	token 	 *JWT
	user     *User
}

func NewAuthManager() *AuthManager {
	return &AuthManager{
		mu:       sync.RWMutex{},
		token:    new(JWT),
		user:     new(User),
	}
}

func (am *AuthManager) CreateUser(ctx context.Context, user *User) (*JWT, error) {
	found, _ := db.GetUserByEmail(user.Email)

	if found == nil {
		//w.WriteHeader(http.StatusNotFound)
		return nil, status.Errorf(codes.NotFound, "No such user")
	}

	if !helpers.CheckPasswordHash(user.Password, found.HashPassword) {
		//w.WriteHeader(http.StatusForbidden)
		return nil, status.Errorf(codes.PermissionDenied, "Incorrect password")
	}
	ttl := time.Hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    found.ID,
		"userEmail": user.Email,
		"exp":       time.Now().UTC().Add(ttl).Unix(),
	})

	str, err := token.SignedString(SECRET)
	if err != nil {
		//w.WriteHeader(http.StatusForbidden)
		//w.Write([]byte("=(" + err.Error()))
		return nil, status.Errorf(codes.PermissionDenied, "Incorrect secret")
	}
	am.token.Token = str
	return am.token, nil
}

//TODO
func (am *AuthManager) GetUser(ctx context.Context, token *JWT) (*User, error) {
	t, _ := jwt.Parse(token.Token, func(token *jwt.Token) (interface{}, error) {
		return SECRET, nil
	})

	claims, _ := t.Claims.(jwt.MapClaims)

	temp := claims["userID"]
	mytemp := uint(temp.(float64))

	user, _ := db.GetUserByID(mytemp)
	if user == nil {
		return nil, status.Errorf(codes.Unknown, "Unauthorized")
	}
	return modelToProto(user), nil
}

func (am *AuthManager) DeleteUser(ctx context.Context, token *JWT) (*Nothing, error) {
	return &Nothing{}, nil
}

func modelToProto (user *models.User) *User {
	return &User{
		ID: uint64(user.ID),
		Login: user.Login,
		Email: user.Email,
		Password: user.Password,
		HashPassword: user.HashPassword,
		Score: uint64(user.Score),
		Picture: user.Picture,
		Games: uint64(user.Games),
	}
}

