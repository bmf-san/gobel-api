package domain

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

// A JWT represents the singular of jwt authentication.
type JWT struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessUUID   string `json:"access_uuid"`
	RefreshUUID  string `json:"refresh_uuid"`
	AtExpires    int64  `json:"at_expires"`
	RtExpires    int64  `json:"rt_expires"`
}

// CreateAccessToken creates an access token.
func (j *JWT) CreateAccessToken(id int) (string, error) {
	jat := jwt.New(jwt.SigningMethodHS256)
	atClaims := jat.Claims.(jwt.MapClaims)
	atClaims["authorized"] = true
	atClaims["access_uuid"] = j.AccessUUID
	atClaims["id"] = id
	atClaims["exp"] = j.AtExpires
	var err error
	if j.AccessToken, err = jat.SignedString([]byte(os.Getenv("JWT_ACCESS_TOKEN_SECRET"))); err != nil {
		return "", err
	}

	return j.AccessToken, nil
}

// CreateRefreshToken creates a refresh token.
func (j *JWT) CreateRefreshToken(id int) (string, error) {
	jrt := jwt.New(jwt.SigningMethodHS256)
	rtClaims := jrt.Claims.(jwt.MapClaims)
	rtClaims["refresh_uuid"] = j.RefreshUUID
	rtClaims["id"] = id
	rtClaims["exp"] = j.RtExpires
	var err error
	if j.RefreshToken, err = jrt.SignedString([]byte(os.Getenv("JWT_REFRESH_TOKEN_SECRET"))); err != nil {
		return "", err
	}

	return j.RefreshToken, nil
}

// GetVerifiedAccessToken gets a verified token from a bearer token.
func (j *JWT) GetVerifiedAccessToken(bearerToken string) (*jwt.Token, error) {
	token := j.ExtractToken(bearerToken)
	verifiedToken, err := j.VerifyToken(token, os.Getenv("JWT_ACCESS_TOKEN_SECRET"))
	if err != nil {
		return nil, err
	}
	return verifiedToken, nil
}

// GetVerifiedRefreshToken gets a verified token from a bearer token.
func (j *JWT) GetVerifiedRefreshToken(bearerToken string) (*jwt.Token, error) {
	token := j.ExtractToken(bearerToken)

	verifiedToken, err := j.VerifyToken(token, os.Getenv("JWT_REFRESH_TOKEN_SECRET"))
	if err != nil {
		return nil, err
	}
	return verifiedToken, nil
}

// GetAccessUUID gets an access uuid from request Authorization header.
func (j *JWT) GetAccessUUID(verifiedToken *jwt.Token) (string, error) {
	accessUUID, err := j.ExtractAccessUUID(verifiedToken)
	if err != nil {
		return "", err
	}

	return accessUUID, nil
}

// GetRefreshUUID gets an refresh uuid from request Authorization header.
func (j *JWT) GetRefreshUUID(verifiedToken *jwt.Token) (string, error) {
	refreshUUID, err := j.ExtractRefreshUUID(verifiedToken)
	if err != nil {
		return "", err
	}

	return refreshUUID, nil
}

// ExtractToken extract a token from authorization header in request.
// For example, if bearToken is "Bearer abcd.efgh.ijkl", return "abcd..efgh.ijkl"
func (j *JWT) ExtractToken(bearToken string) string {
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// VerifyToken verifies jwt.
// See: https://godoc.org/github.com/dgrijalva/jwt-go#example-Parse--Hmac
func (j *JWT) VerifyToken(token string, secret string) (*jwt.Token, error) {
	verifiedToken, err := jwt.Parse(token, func(jt *jwt.Token) (interface{}, error) {
		if _, ok := jt.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jt.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !verifiedToken.Valid {
		return nil, errors.New("invalid token")
	}

	return verifiedToken, nil
}

// ExtractAccessUUID extract data from a verified token.
func (j *JWT) ExtractAccessUUID(t *jwt.Token) (string, error) {
	claims, ok := t.Claims.(jwt.MapClaims)
	if ok && t.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return "", errors.New("type assertion failed")
		}
		return accessUUID, nil
	}

	return "", errors.New("failed extract token data")
}

// ExtractRefreshUUID extract data from a verified token.
func (j *JWT) ExtractRefreshUUID(t *jwt.Token) (string, error) {
	claims, ok := t.Claims.(jwt.MapClaims)
	if ok && t.Valid {
		refreshUUID, ok := claims["refresh_uuid"].(string)
		if !ok {
			return "", errors.New("type assertion failed")
		}
		return refreshUUID, nil
	}

	return "", errors.New("failed extract token data")
}
