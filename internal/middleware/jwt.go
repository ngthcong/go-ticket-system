package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"go-ticket-system/internal/model"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// ErrorResponse ...
type ErrorResponse struct {
	Error string `json:"error"`
}

func TokenAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := TokenValid(r)
		if err != nil {
			ResponseJSON(w, http.StatusUnauthorized, err.Error())
			return
		}

		accessDetail, err := ExtractTokenMetadata(r)
		if err != nil {
			ResponseJSON(w, http.StatusUnauthorized, err.Error())
			return
		}
		r.Header.Set("user_id", strconv.Itoa(accessDetail.UserId))
		r.Header.Set("role", strconv.Itoa(accessDetail.Role))
		next.ServeHTTP(w, r)
	}
}

// ResponseJSON ...
func ResponseJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: err.Error(),
		})
		return
	}
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}
func ExtractTokenMetadata(r *http.Request) (*model.AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, err := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
		if err != nil {
			return nil, err
		}
		role, err := strconv.Atoi(fmt.Sprintf("%v", claims["role"]))
		if err != nil {
			return nil, err
		}
		return &model.AccessDetails{
			UserId: userId,
			Role: role,
		}, nil
	}
	return nil, err
}
