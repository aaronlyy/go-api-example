package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret []byte = []byte("supersecret")

// data about the authenticated user like roles etc
type AccessClaims struct {
	Roles []string `json:"roles,omitempty"`
	jwt.RegisteredClaims
}

// creates a new signed token with given roles and userid
func SignAccessToken(userId string, roles []string, ttl time.Duration) (token string, exp time.Time, err error) {

	// add 15 minutes if ttl <= 0
	if ttl <= 0 {
		ttl = 30 * time.Minute
	}

	// get current time and add ttl to specify expiration date
	now := time.Now().UTC()
	exp = now.Add(ttl)

	// set all needed claims + default claims
	claims := AccessClaims{
		Roles: roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userId,
			Issuer:    "go-api-example-auth",
			Audience:  []string{"go-api-example"},
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	// create token and sign it with given secret
	t := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	// sign token with secret
	signed, err := t.SignedString(secret)

	return signed, exp, err

}

// verify a signed token and get back claims
func ParseAccessToken(tokenStr string) (*AccessClaims, error) {
	tok, err := jwt.ParseWithClaims(tokenStr, &AccessClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	}, jwt.WithIssuer("go-api-example-auth"), jwt.WithAudience("go-api-example"))
	if err != nil {
		return nil, err
	}
	claims, ok := tok.Claims.(*AccessClaims)
	if !ok || !tok.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
