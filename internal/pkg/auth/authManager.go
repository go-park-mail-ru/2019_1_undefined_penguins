package auth

import (
	db "2019_1_undefined_penguins/internal/pkg/auth/database"
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"2019_1_undefined_penguins/internal/pkg/models"
	"fmt"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"sync"
	"time"
)

var SECRET []byte

type AuthManager struct {

	mu       sync.RWMutex
	token 	 *models.JWT
	user     *models.User
	userArray *models.UsersArray
	leaders *models.LeadersInfo

}

func NewAuthManager() *AuthManager {
	return &AuthManager{
		mu:       sync.RWMutex{},
		token:    new(models.JWT),
		user:     new(models.User),
		userArray: new(models.UsersArray),
		leaders: new(models.LeadersInfo),

	}
}

var UsersWantToPlay map[string]*models.User

//TODO check error returns

func (am *AuthManager) LoginUser(ctx context.Context, user *models.User) (*models.JWT, error) {
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
		"userLogin": user.Login,
		"exp":       time.Now().UTC().Add(ttl).Unix(),
	})

	str, err := token.SignedString(SECRET)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Incorrect secret")
	}
	am.token.Token = str
	return am.token, nil
}

func (am *AuthManager) RegisterUser(ctx context.Context, user *models.User) (*models.JWT, error) {
	foundByEmail, _ := db.GetUserByEmail(user.Email)
	foundByLogin, _ := db.GetUserByLogin(user.Login)

	if foundByEmail != nil || foundByLogin != nil {
		return nil, status.Errorf(codes.AlreadyExists, "Such user already exists")
	}

	user.HashPassword = helpers.HashPassword(user.Password)

	err := db.CreateUser(user)
	if err != nil {
		helpers.LogMsg(err)
		return nil, status.Errorf(codes.Internal, "Server error")
	}

	ttl := time.Hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    user.ID,
		"userEmail": user.Email,
		"userLogin": user.Login,
		"exp":       time.Now().UTC().Add(ttl).Unix(),
	})

	str, err := token.SignedString(SECRET)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Incorrect secret")
	}

	am.token.Token = str
	return am.token, nil
}

func (am *AuthManager) GetUser(ctx context.Context, token *models.JWT) (*models.User, error) {
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
	return user, nil
}

func (am *AuthManager) GetUserArray(ctx context.Context, leaders *models.LeadersInfo) (*models.UsersArray, error) {
	am.userArray.Users, _ = db.GetLeaders(int(leaders.ID))
	//fmt.Println(err)
	return am.userArray, nil
}

func (am *AuthManager) GetUserCountInfo(ctx context.Context, nothing *models.Nothing) (*models.LeadersInfo, error) {
	am.leaders, _ = db.UsersCount()
	return am.leaders, nil
}

func (am *AuthManager) ChangeUser(ctx context.Context, user *models.User) (*models.Nothing, error) {
	_, err := db.UpdateUserByID(user, uint(user.ID))
	err = db.UpdateImage(user.Login, user.Picture)
	if err != nil {
		return nil, status.Errorf(codes.AlreadyExists, "Such user already exists")
	}
	return &models.Nothing{}, nil
}

func (am *AuthManager) ChangeUserPicture(ctx context.Context, user *models.User) (*models.Nothing, error) {
	err := db.UpdateImage(user.Login, user.Picture)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unknown error")
	}
	return &models.Nothing{}, nil
}

//TODO DeleteUser() is needed?
func (am *AuthManager) DeleteUser(ctx context.Context, token *models.JWT) (*models.Nothing, error) {
	return &models.Nothing{}, nil
}

func (am *AuthManager) SaveUserGame(ctx context.Context, user *models.User) (*models.Nothing, error) {
	fmt.Println(user.Login, " has ", user.Score)
		err := db.NewRecord(user.Email, int(user.Score))
		if err != nil {
			return nil, status.Errorf(codes.DataLoss, "Error setting record")
		}

	fmt.Println("--------------------------")
	return &models.Nothing{}, nil
}

//here JWT is already valid
func (am *AuthManager) AddUserToGame(ctx context.Context, user *models.User) (*models.Nothing, error) {
	if UsersWantToPlay[user.Login] != nil {
		return nil, status.Errorf(codes.AlreadyExists, "Such user already exists")
	}
	UsersWantToPlay[user.Login] = user
	fmt.Println("Users in game ", UsersWantToPlay)
	return &models.Nothing{}, nil
}

func (am *AuthManager) GetUserForGame(ctx context.Context, user *models.User) (*models.User, error) {
	if UsersWantToPlay[user.Login] != nil {
		return UsersWantToPlay[user.Login], nil
	}
	return nil, nil
}

func (am *AuthManager) DeleteUserFromGame(ctx context.Context, user *models.User) (*models.Nothing, error) {
	if UsersWantToPlay[user.Login] != nil {
		delete(UsersWantToPlay, user.Login)
		return &models.Nothing{}, nil
	}
	return nil, status.Errorf(codes.NotFound, "Such user is not in game")
}

