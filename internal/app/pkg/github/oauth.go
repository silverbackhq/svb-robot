// Copyright 2019 Silverbackhq. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package github

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	// GithubOAuthURL url
	GithubOAuthURL = "https://github.com/login/oauth/authorize"
	// OAuthAccessToken url
	OAuthAccessToken = "https://github.com/login/oauth/access_token"
)

// OAuthApp struct
type OAuthApp struct {
	ClientID     string   `json:"client_id"`
	RedirectURI  string   `json:"redirect_uri"`
	Scope        string   `json:"scope"`
	Scopes       []string `json:"scopes"`
	State        string   `json:"state"`
	AllowSignup  string   `json:"allow_signup"`
	ClientSecret string   `json:"client_secret"`
	AccessToken  string   `json:"access_token"`
	TokenType    string   `json:"token_type"`
}

// OAuthClient struct
type OAuthClient struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

// AccessToken struct
type AccessToken struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	State        string `json:"state"`
}

// GenerateState creates a random string
func (e *OAuthApp) GenerateState() {
	val, err := e.RandomString(20)
	if err == nil {
		e.State = val
	}
}

// GetState returns the state
func (e *OAuthApp) GetState() string {
	return e.State
}

// SetState sets the state
func (e *OAuthApp) SetState(state string) {
	e.State = state
}

// AddScope adds a scope
func (e *OAuthApp) AddScope(scope string) {
	e.Scopes = append(e.Scopes, scope)
	e.Scope = strings.Join(e.Scopes, ",")
}

// AddScopes adds all scopes
func (e *OAuthApp) AddScopes(scopes []string) {
	e.Scopes = scopes
	e.Scope = strings.Join(e.Scopes, ",")
}

// BuildAuthorizeURL get the authorize url
func (e *OAuthApp) BuildAuthorizeURL() string {
	if e.Scope == "" {
		e.Scope = strings.Join(e.Scopes, ",")
	}

	u, err := url.Parse(GithubOAuthURL)

	if err != nil {
		return ""
	}

	q := u.Query()
	q.Set("client_id", e.ClientID)
	q.Set("redirect_uri", e.RedirectURI)
	q.Set("scope", e.Scope)
	q.Set("state", e.State)
	q.Set("allow_signup", e.AllowSignup)
	u.RawQuery = q.Encode()

	return u.String()
}

// RandomString creates a random string
func (e *OAuthApp) RandomString(len int) (string, error) {
	bytes := make([]byte, len)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// FetchAccessToken fetches github access token
func (e *OAuthApp) FetchAccessToken(code string, state string) (bool, error) {

	accessTokenRequest := &AccessToken{
		ClientID:     e.ClientID,
		ClientSecret: e.ClientSecret,
		Code:         code,
		State:        e.State,
	}

	jsonBody, err := accessTokenRequest.ConvertToJSON()

	if err != nil {
		return false, err
	}

	githubOAuthClient := &OAuthClient{}

	if state != e.State {
		return false, fmt.Errorf(
			"Invalid state provided %s, original one is %s",
			state,
			e.State,
		)
	}

	client := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		OAuthAccessToken,
		bytes.NewBufferString(jsonBody),
	)

	if err != nil {
		return false, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return false, err
	}

	err = json.Unmarshal(bodyByte, &githubOAuthClient)

	if err != nil {
		return false, err
	}

	e.AccessToken = githubOAuthClient.AccessToken
	e.TokenType = githubOAuthClient.TokenType

	return true, nil
}

// GetAccessToken gets access token
func (e *OAuthApp) GetAccessToken() string {
	return e.AccessToken
}

// GetTokenType gets token type
func (e *OAuthApp) GetTokenType() string {
	return e.TokenType
}

// ConvertToJSON convert object to json
func (e *AccessToken) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&e)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
