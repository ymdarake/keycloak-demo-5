package middleware

import (
	"context"
	"fmt"
	"keycloak-demo-5/pkg/authz"
	"net/http"
	"strings"
)

type contextKey struct {
	name string
}

var authTokenCtxKey = &contextKey{"authToken"}

func MustAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeaderValue := r.Header.Get("Authorization")

		if authorizationHeaderValue == "" {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		encodedToken := strings.TrimSpace(strings.TrimPrefix(authorizationHeaderValue, "Bearer "))
		// TODO: introspect/verify token
		ctx := context.WithValue(r.Context(), authTokenCtxKey, encodedToken)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})

}

func GetToken(ctx context.Context) *authz.Token {
	encodedToken := ctx.Value(authTokenCtxKey).(string)

	tok, err := authz.TokenFromString(encodedToken)
	if err != nil {
		fmt.Printf("token decode error: %v\n\n", err)
		return nil
	}
	return tok
}
