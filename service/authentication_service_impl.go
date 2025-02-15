package service

import (
	"context"
	"database/sql"
	"go-tobfa/exception"
	"go-tobfa/helper"
	"go-tobfa/model"
	"go-tobfa/model/web"
	"go-tobfa/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
	Config         *viper.Viper
}

func NewAuthenticationService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate, config *viper.Viper) AuthenticationService {
	return &AuthenticationServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
		Config:         config,
	}
}

func (service *AuthenticationServiceImpl) Login(ctx context.Context, request web.LoginRequest) web.AuthResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		panic(exception.NewNotFoundError("password tidak cocok"))
	}

	claims := model.JwtClaim{
		Email: request.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretString := []byte(service.Config.GetString("SECRET_KEY"))
	tokenString, err := token.SignedString(secretString)
	helper.PanicIfError(err)

	return helper.ToAuthResponse(user, tokenString)

}
