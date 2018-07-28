package skaioskit

import (
    "errors"
    "net/http"

    jwt "github.com/dgrijalva/jwt-go"
)

type JWTData struct {
    jwt.StandardClaims

    UserId uint
    Email string
}

func GenerateJWTStr(data JWTData, key []byte) (string, error) {
    token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &data)
    return token.SignedString(key)
}

func ParseJWT(jwtStr string, key []byte) (*JWTData, error) {
    token, err := jwt.ParseWithClaims(jwtStr, &JWTData{}, func(token *jwt.Token) (interface{}, error) {
        return key, nil
    })

    if claims, ok := token.Claims.(*JWTData); ok && token.Valid {
        return claims, nil
    } else {
        return claims, err
    }
}

func GetJwtForRequest(r *http.Request, key []byte) (*JWTData, error) {
    //Authorization
    authHeader := r.Header.Get("Authorization")
    if len(authHeader) < len("Bearer ") {
        return nil, errors.New("missing auth header.")
    }

    token := authHeader[len("Bearer "):]
    return ParseJWT(token, key)
}

func JWTEnforceMiddleware(key []byte) func(next http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            _, err := GetJwtForRequest(r, key)

            if err != nil {
                http.Error(w, http.StatusText(401), 401)
                return
            }

            next.ServeHTTP(w, r)
        })
    }
}
