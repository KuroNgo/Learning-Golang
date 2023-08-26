package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"

	"server/conf"
	"server/model"
)

// GetGoogleOAuthToken là hàm gửi request lên Google OAuth để lấy auth
// Hàm này sẽ trả về một struct chứa auth
// Hàm này sẽ trả về một error nếu có lỗi xảy ra
func GetGoogleOAuthToken(code string) (*model.GoogleOAuthToken, error) {
	const rootURL = "https://oauth2.googleapis.com/token"

	config, _ := conf.LoadConfig(".")
	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("client_id", config.GoogleClientID)
	values.Add("client_secret", config.GoogleClientSecret)
	values.Add("redirect_uri", config.GoogleOAuthRedirectURL)

	query := values.Encode()

	req, err := http.NewRequest("POST", rootURL, bytes.NewBufferString(query))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Google OAuth auth request failed")
	}
	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, res.Body)
	if err != nil {
		return nil, err
	}

	var GoogleOAuthToken map[string]interface{}
	if err := json.Unmarshal(resBody.Bytes(), &GoogleOAuthToken); err != nil {
		return nil, err
	}

	tokenBody := &model.GoogleOAuthToken{
		Access_token: GoogleOAuthToken["access_token"].(string),
		Id_token:     GoogleOAuthToken["id_token"].(string),
	}
	return tokenBody, nil
}
