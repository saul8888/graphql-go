package middelware

import (
	"go-graphql/authentication"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"
)

//authentication
var autoConfig = middleware.JWTConfig{
	Claims:        &authentication.Claim{},
	SigningMethod: jwt.SigningMethodHS256.Name,
	SigningKey:    authentication.Keys(),
	//SigningMethod: jwt.SigningMethodRS256.Name,
	//SigningKey:    authentication.PublicKey,
	//TokenLookup: "header:" + echo.HeaderAuthorization,
}

//Bearer {token}
