package middlewares

import (
	"log"
	"net/http"
	"os"

	"github.com/cocoaster/golang-jwt/configs"
	"github.com/cocoaster/golang-jwt/helper"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func JWTMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var response map[string]string
		//take a cookie token
		c,err := r.Cookie("token")
		if err == http.ErrNoCookie {
			response = map [string]string{"message": "unauthorized"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}
		//take token value
		tokenString := c.Value
		claims := &configs.JwtClaims{}

		//parsing token JWT
		err = godotenv.Load()
		if err != nil{
			log.Fatal("Error loading .Env file")
		}
		jwt_key := []byte(os.Getenv("JWT_KEY"))
		token,err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return jwt_key, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid{
				response = map[string]string{"message": "invalid token signature!"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
			response = map[string]string{"message": "invalid token!"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}
		//if token not valid
		if !token.Valid{
			response = map[string]string{"message": "invalid token!"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}
		next.ServeHTTP(w, r)
	})
}