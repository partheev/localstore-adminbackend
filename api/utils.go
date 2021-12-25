package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt"
)

type res struct {
	Message string
}

func resMessage(s string) res {
	response := res{
		Message: s,
	}
	return response
}
func (app *application) writeJson(w http.ResponseWriter, data interface{}, status int) {
	jsonData, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		w.Write([]byte("error"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
}
func (app *application) writeError(w http.ResponseWriter, err error) {
	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	app.writeJson(w, theError, http.StatusBadRequest)

}

type Claims struct {
	UserId int `json:"userId"`
	jwt.StandardClaims
}

func (app *application) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		cookieToken, err := req.Cookie("session")
		tokenString := cookieToken.Value
		if err != nil {
			app.writeError(w, err)
		}
		claims := Claims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(a *jwt.Token) (interface{}, error) {
			return []byte(app.Config.JwtKey), nil
		})
		if err != nil {
			app.writeError(w, err)
		}
		if claim, ok := token.Claims.(*Claims); ok && token.Valid {
			adminUser := app.DB.DBModel.GetAdminUser(claim.UserId)
			ctx := context.WithValue(req.Context(), "adminuser", adminUser)
			next.ServeHTTP(w, req.WithContext(ctx))
		} else {
			app.writeError(w, err)
		}
	})
}
