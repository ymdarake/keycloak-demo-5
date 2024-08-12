package middleware

import (
	"context"
	"fmt"
	"keycloak-demo-5/pkg/authz"
	"net/http"
)

type contextKey struct {
	name string
}

var authTokenCtxKey = &contextKey{"authToken"}

func MustAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// NOTE: Envoyで検証、抽出している
		jwtPayload := r.Header.Get("X-Verified-Jwt-Payload")
		if jwtPayload == "" {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), authTokenCtxKey, jwtPayload)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})

}

func GetToken(ctx context.Context) *authz.TokenPayload {
	encodedToken := ctx.Value(authTokenCtxKey).(string)
	tok, err := authz.TokenPayloadFromString(encodedToken)
	if err != nil {
		fmt.Printf("token decode error: %v\n\n", err)
		return nil
	}
	return tok
}
