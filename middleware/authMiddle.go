package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/ardhisaif/golang_backend/helpers"
)

func AuthMiddle(role ...string) Middle {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-type", "application/json") // set header to json
			
			var header string
			var valid bool

			header = r.Header.Get("Authorization")
			if header == "" {
				helpers.New(http.StatusUnauthorized, "header not provide, please login").Send(w)
				return
			}

			if !strings.Contains(header, "Bearer") {
				helpers.New(http.StatusUnauthorized, "invalid header type").Send(w)
				return
			}

			tokens := strings.Replace(header, "Bearer ", "", -1)

			checkToken, err := helpers.CheckToken(tokens)
			if err != nil {
				helpers.New(http.StatusUnauthorized, err.Error()).Send(w)
				return
			}

			for _, role := range role {
				if role == checkToken.Role {
					valid = true
				}
			}

			if !valid {
				helpers.New(http.StatusUnauthorized, "you don`t have access permission").Send(w)
				return
			}

			log.Println("Auth middleware pass")

			ctx := context.WithValue(r.Context(), "user", checkToken.UserId)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
