package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	jsonwebtoken "github.com/golang-jwt/jwt/v5"
	envsettings "github.com/kasariks/project_for_graduating/internal/env_settings"
	"github.com/kasariks/project_for_graduating/tests"
)

const jwtKey = "334dfasdfqewrqwerqwerdkp123afv"

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// смотрим наличие пароля
		pass := envsettings.Env.Password
		if len(pass) > 0 {
			var jwt string // JWT-токен из куки
			// получаем куку
			cookie, err := r.Cookie("token")
			if err == nil {
				jwt = cookie.Value
			}
			var valid bool = true
			// здесь код для валидации и проверки JWT-токена

			if len(jwt) == 0 {
				valid = false
			} else {
				_, err = jsonwebtoken.Parse(jwt, func(token *jsonwebtoken.Token) (interface{}, error) {
					return []byte(jwtKey), nil
				})
				if err != nil {
					valid = false
				}
			}

			if !valid {
				// возвращаем ошибку авторизации 401
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}
		}
		next(w, r)
	})
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	pass := map[string]string{}
	err := json.NewDecoder(r.Body).Decode(&pass)
	if err != nil {
		writeErrorInJson(w, err)
		return
	}

	if pass["password"] == envsettings.Env.Password {
		claims := jwt.MapClaims{
			"exp": time.Now().Add(time.Minute * 15).Unix(),
		}

		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, err := jwtToken.SignedString([]byte(jwtKey))
		if err != nil {
			writeErrorInJson(w, err)
			return
		}

		cookie := &http.Cookie{
			Name:     "token",
			Value:    signedToken,
			Expires:  time.Now().Add(time.Minute * 15),
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
			Path:     "/",
		}

		http.SetCookie(w, cookie)
		err = json.NewEncoder(w).Encode(map[string]string{
			"token": signedToken,
		})
		if err != nil {
			writeErrorInJson(w, err)
			return
		}
		tests.Token = signedToken
	} else {
		writeErrorInJson(w, errors.New("wrong password"))
		return
	}
}

func writeErrorInJson(w http.ResponseWriter, err error) {
	byteErr, _ := json.Marshal(map[string]string{"error": err.Error()})
	w.Write(byteErr)
}
