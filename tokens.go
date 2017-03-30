package ecobee

import(
  "net/http"
  "net/url"
  "fmt"
  "encoding/json"
)

type Tokens struct {
  AccessToken string `json:"access_token"`
  RefreshToken string `json:"refresh_token"`
  TokenType string `json:"token_type"`
  ExpiresIn int `json:"expires_in"`
  Scope string `json:"scope"`
}

func (e *Ecobee) GetTokens() (tokens *Tokens, err error) {
  authCode := url.QueryEscape(e.AuthCode)
  clientId := url.QueryEscape(e.ApiKey)
  url := fmt.Sprintf("https://api.ecobee.com/token?grant_type=ecobeePin&code=%s&client_id=%s", authCode, clientId)
  req, err := http.NewRequest("POST", url, nil)
  if err != nil {
      return tokens, err
  }
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
      return tokens, err
  }
  defer resp.Body.Close()
  err = json.NewDecoder(resp.Body).Decode(&tokens)
  if err != nil {
    return tokens, err
  }
  return tokens, nil
}

func (e *Ecobee) RefreshTokens() (tokens *Tokens, err error) {
  refreshToken := url.QueryEscape(e.RefreshToken)
  clientId := url.QueryEscape(e.ApiKey)
  url := fmt.Sprintf("https://api.ecobee.com/token?grant_type=refresh_token&refresh_token=%s&client_id=%s", refreshToken, clientId)
  req, err := http.NewRequest("POST", url, nil)
  if err != nil {
      return tokens, err
  }
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return tokens, err
  }
  defer resp.Body.Close()
  err = json.NewDecoder(resp.Body).Decode(&tokens)
  if err != nil {
      return tokens, err
  }
  return tokens, nil
}