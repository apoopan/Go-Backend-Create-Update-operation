package account
/*
import (
	//"context"
	"encoding/json"
	"fmt"
	//"log"
	"net/http"

	
	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"

    "github.com/dgrijalva/jwt-go"
	//"go.mongodb.org/mongo-driver/mongo/options"
    
   "strings"

   // "github.com/dgrijalva/jwt-go"
  //"github.com/gorilla/context"
	//"github.com/mitchellh/mapstructure"
	//"github.com/go-playground/validator/v10"

)

type JwtToken struct {
    Token string `json:"token"`
}

type Exception struct {
    Message string `json:"message"`
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        authorizationHeader := req.Header.Get("authorization")
        if authorizationHeader != "" {
            bearerToken := strings.Split(authorizationHeader, " ")
            if len(bearerToken) == 2 {
                token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
                    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                        return nil, fmt.Errorf("There was an error")
                    }
                    return []byte("secret"), nil
                })
                if error != nil {
                    json.NewEncoder(w).Encode(Exception{Message: error.Error()})
                    return
                }
                if token.Valid {
                    //context.Set(req, "decoded", token.Claims)
                    next(w, req)
                } else {
                    json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
                }
            }
        } else {
            json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
        }
    })
}
*/