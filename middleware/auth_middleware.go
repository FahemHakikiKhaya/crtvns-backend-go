package middleware

import (
	"context"
	"net/http"

	"github.com/FahemHakikiKhaya/crtvns-backend-go/helper"
	"github.com/julienschmidt/httprouter"
)

type ContextKey string

const UserIDKey ContextKey = "userId"

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tokenString, err := helper.ExtractToken(r)
		if err != nil {
			http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}

		_, claims, err := helper.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}

		userIdFloat, ok := claims["userId"].(float64) 
		if !ok {
			http.Error(w, "Unauthorized: Invalid userId format", http.StatusUnauthorized)
			return
		}

		userId := int(userIdFloat) 

		ctx := context.WithValue(r.Context(), UserIDKey, userId)

		next(w, r.WithContext(ctx), ps)
	}
}
