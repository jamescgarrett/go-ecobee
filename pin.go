package ecobee

import(
  "net/http"
  "net/url"
  "fmt"
  "encoding/json"
)

type Pin struct {
  EcobeePin string `json:"ecobeePin"`
  Code string `json:"code"`
}

func (e *Ecobee) GetPin() (*Pin, error) {
    url := fmt.Sprintf("https://api.ecobee.com/authorize?response_type=ecobeePin&client_id=%s&scope=smartWrite", url.QueryEscape(e.ApiKey))
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    var epin *Pin
    err = json.NewDecoder(resp.Body).Decode(&epin)
    if err != nil {
        return nil, err
    }
    return epin, nil
}