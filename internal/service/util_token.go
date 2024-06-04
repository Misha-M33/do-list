package service

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// type Keys struct {
// 	AccessKey  string `json:"access_key"`
// 	RefreshKey string `json:"refresh_key"`
// }

// var KeyTokens = Keys{
// 	AccessKey:  "access",
// 	RefreshKey: "refresh",
// }

func CreateAccessToken(id string, c *fiber.Ctx) (accessToken string, err error) {
	godotenv.Load()
	secret := os.Getenv("ACCESS_KEY")

	claims := jwt.MapClaims{
		"userID": id,
		"exp":    time.Now().Add(time.Minute * 15).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return t, c.Status(500).SendString(err.Error())
	}
	return t, err
}

func CreateRefreshToken(id string, c *fiber.Ctx) (refreshToken string, err error) {
	godotenv.Load()
	secret := os.Getenv("REFRESH_KEY")

	claims := jwt.MapClaims{
		"userID": id,
		"exp":    time.Now().Add(time.Hour * 24 * 15).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return t, c.Status(500).SendString(err.Error())
	}
	return t, err
}

func ExtractIDFromToken(requestToken string, secret string) (string, error) {

	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("error get user claims from token")
	}
	return claims["userID"].(string), nil
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetUserIdFromHeader(c *fiber.Ctx) (string, error) {

	type auth struct {
		Authorization []string
	}

	var t auth
	err := c.ReqHeaderParser(&t)
	if err != nil {
		log.Println(err)
		return "", c.Status(500).SendString(err.Error())
	}

	val := strings.Split(t.Authorization[0], " ")

	tok, _ := jwt.Parse(val[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})

	claims := tok.Claims.(jwt.MapClaims)

	userID := fmt.Sprintf("%v", claims["userID"])

	return userID, nil
}
