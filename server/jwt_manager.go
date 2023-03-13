package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

// User contains user's information
type User struct {
	Username       string
	HashedPassword string
	Role           string
}

// JWTManager is a JSON web token manager
type JWTManager struct {
	key           string
	tokenDuration time.Duration
}

// UserClaims is a custom JWT claims that contains some user's information
type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}

func accessibleRoles() map[string][]string {
	const laptopServicePath = "/techschool.pcbook.LaptopService/"

	return map[string][]string{
		laptopServicePath + "CreateLaptop": {"admin"},
		laptopServicePath + "UploadImage":  {"admin"},
		laptopServicePath + "RateLaptop":   {"admin", "user"},
		"/event.grpc.Event/GetEvent":       {"admin", "user"},
	}
}

// NewJWTManager returns a new JWT manager
func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secretKey, tokenDuration}
}

// Generate generates and signs a new token for a user
func (manager *JWTManager) Generate(user *User) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
		Username: user.Username,
		Role:     user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.key))
}

// Verify verifies the access token string and return a user claim if the token is valid
func (manager *JWTManager) Verify(accessToken string) (*UserClaims, error) {

	// test
	/**
	test := func() {
		//var Hash crypto.Hash
		var parts []string

		parts = strings.Split(accessToken, ".")
		token := &jwt.Token{Raw: accessToken}

		token.Signature = parts[2]

		signingString := strings.Join(parts[0:2], ".")

		sig, _ := jwt.DecodeSegment(token.Signature)
		hasher := hmac.New(sha256.New, []byte(manager.key))
		hasher.Write([]byte(signingString))
		if !hmac.Equal(sig, hasher.Sum(nil)) {
			fmt.Println(sig)
			fmt.Println(hasher.Sum(nil))
		}
	}

	test()
	**/

	// test access the claim without verifying the signature to identify the caller

	//
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			kid, ok := token.Header["kid"].(string)
			if !ok {
				return nil, fmt.Errorf("missing or invalid 'kid' claim")
			}
			// we don't check if x5u is there, we consider it as optional
			xfu, _ := token.Header["x5u"].(string)
			fmt.Println(kid)
			fmt.Println(xfu)
			_, ok = token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				_, ok := token.Method.(*jwt.SigningMethodRSA)
				if !ok {
					return nil, fmt.Errorf("unexpected token signing method")
				}
				publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(keyManager.keymap[kid]))
				if err != nil {
					log.Fatal("Error while loading the public key: ", err)
				}
				return publicKey, nil
			}

			return []byte(keyManager.keymap[kid]), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
