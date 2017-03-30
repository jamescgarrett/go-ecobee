# Go Ecobee


Go Ecobee is a package for handling the Ecobee thermostat API. [Ecobee API](https://www.ecobee.com/home/developer/api/introduction/index.shtml).


## Getting Started

1) Follow Ecobee's instructions to get your API KEY. [Instructions Here](https://www.ecobee.com/home/developer/api/examples/ex1.shtml)
2) `go get github.com/jamescgarrett/go-ecobee`
3) Anywhere in your application directory create a mostly blank .yaml file with the following content:
```yaml
api_key: YOUR_API_KEY
pin: ""
access_token: ""
auth_code: ""
refresh_token: ""
expires_in: ""
thermostat_id: ""
interval_revision: ""
```

## Negroni Example
```go
package main

import(
  "net/http"
  "log"
  "io/ioutil"
  "encoding/json"
  
  "github.com/gorilla/mux"
  "github.com/codegangsta/negroni"
  "github.com/unrolled/render"
  "github.com/jamescgarrett/go-ecobee"
)

func main() {  
  n := negroni.New()
    
    // using unrolled/render for this example
    render =: render.New(render.Options{})
    
    // create Ecobee config
    ecob =: ecobee.GetConfig("config.yaml")
    
    // using gorilla mux for this example
    router := mux.NewRouter()
    
    // this route checks whether setup is complete
    router.HandleFunc("/setupstatus", getSetupStatus).Methods("GET")
    
    // route for getting your Ecobee pin
    router.HandleFunc("/pin", getPin).Methods("GET")
    
    // route for getting your access and refresh tokens
    router.HandleFunc("/tokens", getTokens).Methods("GET")
    
    // route for accessing your thermostat data
    router.HandleFunc("/thermostats", getThermostats).Methods("POST")
    
    // route for getting a report
    router.HandleFunc("/report", getReport).Methods("GET")
    
    // handle the /setupstatus route
    func getSetupStatus(w http.ResponseWriter, r *http.Request){
      // looks like setup was complete so pass on data
        if ecob.AccessToken != "" && ecob.RefreshToken != "" {
          render.JSON(w, http.StatusOK, map[string]interface{}{
              "status": true, // I use this to tell my front end to redirect to where i show thermostat data
                "pin": ecob.Pin,
                "accessToken": ecob.AccessToken,
                "refreshToken": ecob.RefreshToken,
            })
          return
        }
    
        // no setup was done...
        render.JSON(w, http.StatusOK, map[string]interface{}{
          "status": false,
        })
    }
    
    // handle the /pin route
    func getPin(w http.ResponseWriter, r *http.Request){
      if ecob.AuthCode != "" && ecob.AuthCode != "" {
          render.JSON(w, http.StatusOK, map[string]interface{}{
            "pin": ecob.Pin,
          })
          return
        }
        pin, err := ecob.GetPin()
        if err != nil {
          render.JSON(w, http.StatusOK, map[string]interface{}{
            "error": err.Error(),
          })
          return
        }
        
        // write your config
      ecob.AuthCode = pin.Code
      ecob.Pin = pin.EcobeePin
      ecob.WriteConfig("config/ecobee.yaml")

        render.JSON(w, http.StatusOK, map[string]interface{}{
          "pin": pin.EcobeePin,
        })
  }
    
    // handle the /tokens route
    func getTokens(w http.ResponseWriter, r *http.Request){
      if ecob.AccessToken != "" {
        render.JSON(w, http.StatusOK, map[string]interface{}{
            "status": true,
        })
        return
      }

      tokens, err := ecob.GetTokens()
      if err != nil {
        rend.JSON(w, http.StatusOK, map[string]interface{}{
            "error": err.Error(),
        })
        return
      }
    
        // write Ecobee config
      ecob.AccessToken = tokens.AccessToken
      ecob.RefreshToken = tokens.RefreshToken
      ecob.ExpiresIn = tokens.ExpiresIn
      ecob.WriteConfig("config/ecobee.yaml")
  
      rend.JSON(w, http.StatusOK, map[string]interface{}{
        "accessToken": tokens.AccessToken,
        "refreshToken": tokens.RefreshToken,
      })
  }
  
    // handle the /thermostats route
  func getThermostats(w http.ResponseWriter, r *http.Request) {

      // check summary before requesting full data
        summary, details, err := ecob.GetSummary()
        if err != nil {
          render.JSON(w, http.StatusOK, map[string]interface{}{
              "error": err.Error(),
          })
          return
        }
        
        // check for expired tokens
      // get and write new tokens to config if expired
      if summary.Status.Code == 14 {
        refresh, err := ecob.RefreshTokens()
        if err != nil {
            render.JSON(w, http.StatusOK, map[string]interface{}{
              "error": err.Error(),
            })
        }
        ecob.AccessToken = refresh.AccessToken
        ecob.RefreshToken = refresh.RefreshToken
        ecob.WriteConfig("config/ecobee.yaml")
         } else {
      
          // check for new revision
        // if revision write new config
          if ecob.IntervalRevision == details.IntervalRevision {
            render.JSON(w, http.StatusOK, map[string]interface{}{
              "error": "The Ecobee API server has not updated yet. Please try again later.",
            })
            return
        }
        ecob.ThermostatId = details.ThermostatId
        ecob.IntervalRevision = details.IntervalRevision
        ecob.WriteConfig("config/ecobee.yaml")
      }

      // so get thermostat data
      // get selections from POST
      decoder := json.NewDecoder(r.Body)
      var req *ecobee.Selections   
      err = decoder.Decode(&req)
      if err != nil {
        render.JSON(w, http.StatusOK, map[string]interface{}{
            "error": err.Error(),
        })
        return
      }
      defer r.Body.Close()

      // build selection string
        // see below for these values
      selections := req.BuildSelections()

      // using selection string get thermostats
      thermostats, err := ecob.GetThermostats(selections)
      if err != nil {
        render.JSON(w, http.StatusOK, map[string]interface{}{
            "error": err.Error(),
        })
      }

      render.JSON(w, http.StatusOK, map[string]interface{}{
        "data": thermostats,
      })
  }
    
    // Run Server
    n.Use(negroni.NewLogger())
    n.UseHandler(routes)
    http.ListenAndServe(":8080", n)
}

```


### Thermostat Request Selections
The route ```/thermostats``` is a POST request. The body of the post should consist of string boolean values for each of the following:
```
includeRuntime
includeExtendedRuntime
includeElectricity
includeSetting
includeLocation
includeProgram
includeEvents
includeDevice
includeTechnician
includeUtility
includeAlerts
includeWeather
includeOemConfig
includeEquipmentStatus
includeNotificationSettings
includePrivacy
includeVersion
includeSecuritySettings
includeSensors
```

Example request: (curl example comin soon)
```html
<form>
  <input type='checkbox' name='includeRuntime' value="true" /> Include Runtime
  
    <input type='checkbox' name='includeExtendedRuntime' value="false"/> include Extended Runtime
  
    <input type='checkbox' name='includeElectricity' value="true" /> Include Electricity

...

```


## To Do
- Add Tests
- Example front end


## Licenses

All source code is licensed under the MIT License.
