package authentication

import (
	"context"
	"go-graphql/database"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/labstack/echo"
)

type Service interface {
	GenerateCustomer(context echo.Context) error
	ValidateToken(context echo.Context) error
}

type service struct {
	repo database.Mongodata
}

func NewService(repo database.Mongodata) Service {
	return &service{repo: repo}
}

func (s *service) GenerateCustomer(c echo.Context) (err error) {
	jsonResult := new(Responsetoken)
	var answer = http.StatusOK
	date := new(dateValidate)
	if err = c.Bind(date); err != nil {
		answer = http.StatusForbidden
		return c.JSON(answer, err)
	}
	datevalidate["email"] = date.Email
	datevalidate["password"] = date.Password
	row, err := s.repo.Search(datevalidate)
	if err != nil {
		answer = http.StatusForbidden
		return c.JSON(answer, err)
	}

	customer := &Customer{}
	for row.Next(context.TODO()) {
		row.Decode(&customer)
	}
	if customer.Name != "" && customer.Email != "" {
		//create a struct of my Claim
		claims := Claim{
			Customer: *customer,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
				Issuer:    "token test", //object of token
			},
		}

		//--------------------encode to base64-----------------//
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokeS, err := token.SignedString(Keys())
		//token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
		//tokeS, err := token.SignedString(PrivateKey)
		if err != nil {
			tokeS = "could not sign private token"
		}
		answer = http.StatusOK
		jsonResult.Token = tokeS
	} else {
		answer = http.StatusForbidden
		jsonResult.Token = "usser or password invalid"
	}
	return c.JSON(answer, jsonResult)
}

func (s *service) ValidateToken(context echo.Context) error {
	jsonResult := new(Responsetoken)
	var answer = http.StatusOK
	token, err := request.ParseFromRequestWithClaims(context.Request(), request.OAuth2Extractor, &Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return PublicKey, nil
		})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				answer = http.StatusUnauthorized
				jsonResult.Token = "your token expired"
			case jwt.ValidationErrorSignatureInvalid:
				answer = http.StatusUnauthorized
				jsonResult.Token = "the signature does not match"
			default:
				answer = http.StatusUnauthorized
				jsonResult.Token = "the signature does not match"
			}
		default:
			answer = http.StatusUnauthorized
			jsonResult.Token = "your token is not valid"
		}
	}
	if token.Valid {
		answer = http.StatusAccepted
		jsonResult.Token = "welcome to the system"
	} else {
		answer = http.StatusUnauthorized
		jsonResult.Token = "your token is not valid"
	}
	return context.JSON(answer, jsonResult)
}
