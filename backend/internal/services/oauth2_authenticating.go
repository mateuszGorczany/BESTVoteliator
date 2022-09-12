package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	repository "github.com/mateuszGorczany/BESTVoteliator/internal/dao"
	"github.com/mateuszGorczany/BESTVoteliator/internal/datastruct"
	"github.com/spf13/viper"
	// "github.com/mateuszGorczany/BESTVoteliator/utils/common"
)

type OAuth2AuthenticationService interface {
	Login() error
}

type oAuth2AuthenticationService struct {
	repository repository.DAO
}

func NewOAuth2AuthenticationService(repo repository.DAO) OAuth2AuthenticationService {
	return &oAuth2AuthenticationService{repository: repo}
}

func (o *oAuth2AuthenticationService) Login() error {
	return nil
}

func getGooglePublicKey(keyID string) (string, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/certs")
	if err != nil {
		return "", err
	}
	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	myResp := map[string]string{}
	err = json.Unmarshal(dat, &myResp)
	if err != nil {
		return "", err
	}
	key, ok := myResp[keyID]
	if !ok {
		return "", errors.New("key not found")
	}
	return key, nil
}

func ValidateGoogleJWT(tokenString string) (datastruct.GoogleClaims, error) {
	claimsStruct := datastruct.GoogleClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) {
			pem, err := getGooglePublicKey(fmt.Sprintf("%s", token.Header["kid"]))
			if err != nil {
				return nil, err
			}
			key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
			if err != nil {
				return nil, err
			}
			return key, nil
		},
	)
	if err != nil {
		return datastruct.GoogleClaims{}, err
	}

	claims, ok := token.Claims.(*datastruct.GoogleClaims)
	if !ok {
		return datastruct.GoogleClaims{}, errors.New("Invalid Google JWT")
	}

	if claims.Issuer != "accounts.google.com" && claims.Issuer != "https://accounts.google.com" {
		return datastruct.GoogleClaims{}, errors.New("iss is invalid")
	}

	if claims.Audience != viper.GetString("Google_OAuth_Client_ID") {
		return datastruct.GoogleClaims{}, errors.New("aud is invalid")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return datastruct.GoogleClaims{}, errors.New("JWT is expired")
	}

	return *claims, nil
}
