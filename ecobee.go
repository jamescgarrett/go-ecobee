package ecobee

import(
  "log"
  "fmt"
  "io/ioutil"

  "gopkg.in/yaml.v2"
)

type Ecobee struct {
  ApiKey string `yaml:"api_key"`
  Pin string `yaml:"pin"`
  AccessToken string `yaml:"access_token"`
  AuthCode string `yaml:"auth_code"`
  RefreshToken string `yaml:"refresh_token"`
  ExpiresIn int `yaml:"expires_in"`
  ThermostatId string `yaml:"thermostat_id"`
  IntervalRevision string `yaml:"interval_revision"`
}

func GetConfig(file string) *Ecobee {
  config, err := ioutil.ReadFile(file)
  if err != nil {
    log.Print(err)
  }
  var e *Ecobee
  err = yaml.Unmarshal(config, &e)
  if err != nil {
    log.Fatal(err)
  }
  return e
}

func (e *Ecobee) WriteConfig(file string) {
  w, _ := yaml.Marshal(e)
  err := ioutil.WriteFile(file, w, 0644)
  if err != nil {
      fmt.Println("Could not write config file.")
  }
}
