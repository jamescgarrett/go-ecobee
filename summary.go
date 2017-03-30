package ecobee

import(
  "net/http"
  "net/url"
  "fmt"
  "encoding/json"
  "strings"
)

type Summary struct {
  RevisionList []string `json:"revisionList"`
  ThermostatCount int `json:"thermostatCount"`
  StatusList []string `json:"statusList"`
  Status Status `json:"status"`
}

type SummaryDetails struct {
  ThermostatId string
  IntervalRevision string
}

func (e *Ecobee) GetSummary() (summary *Summary, details SummaryDetails, err error) {
  selection := `{
    "selection":{
      "selectionType": "registered",
      "selectionMatch":"",
      "includeEquipmentStatus": true
    }
  }`
  url := fmt.Sprintf("https://api.ecobee.com/1/thermostatSummary?json=%s", url.QueryEscape(selection))
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return summary, details, err
  }
  req.Header.Set("Content-Type", "application/json;charset=UTF-8")
  req.Header.Set("Authorization", "Bearer " + e.AccessToken)
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
      return summary, details, err
  }
  defer resp.Body.Close()
  err = json.NewDecoder(resp.Body).Decode(&summary)
  if err != nil {
    return summary, details, err
  }
  for _, value := range summary.RevisionList {
    details.ThermostatId = getStringFromCSV(value, 0)
    details.IntervalRevision = getStringFromCSV(value, 6)
  }
  return summary, details, nil
}

/* 
  "Thermostat Identifier" : strings[0]
  "Thermostat Name" : strings[1]
  "Connected" : strings[2]
  "Thermostat Revision" : strings[3]
  "Alerts Revision" : strings[4]
  "Runtime Revision" : strings[5]
  "Interval Revision" : strings[6]
*/
func getStringFromCSV(str string, idx int) string {
  strings := strings.Split(str, ":")
  return strings[idx]
}
