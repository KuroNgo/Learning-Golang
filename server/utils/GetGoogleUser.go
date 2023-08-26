package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"server/model"
)

func GetGoogleUser(access_token string, id_token string) (*model.GoogleUserResult, error) {
	rootUrl := fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?access_token=%s", access_token)

	req, err := http.NewRequest("GET", rootUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", id_token))
	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Could not retrieve user")
	}

	var resBody bytes.Buffer
	_, err = resBody.ReadFrom(res.Body)
	if err != nil {
		return nil, err
	}

	var GoogleUserRes map[string]interface{}
	if err := json.Unmarshal(resBody.Bytes(), &GoogleUserRes); err != nil {
		return nil, err
	}

	userBody := &model.GoogleUserResult{
		Id:             GoogleUserRes["id"].(string),
		Email:          GoogleUserRes["email"].(string),
		Verified_email: GoogleUserRes["verified_email"].(bool),
		Name:           GoogleUserRes["name"].(string),
		Given_name:     GoogleUserRes["given_name"].(string),
		Picture:        GoogleUserRes["picture"].(string),
		Locale:         GoogleUserRes["locale"].(string),
	}
	return userBody, nil
}
