package ecobee

import(
  "net/http"
  "net/url"
  "fmt"
  "encoding/json"
)

type Thermostats struct {
  Page Page `json:"page"`
  ThermostatList []Thermostat `json:"thermostatList"`
  Status Status `json:"status"`
}

type Page struct {
  Page int `json:"page"`
  TotalPages int `json:"totalPages"`
  PageSize int `json:"pageSize"`
  Total int `json:"total"`
}

type Thermostat struct {
  Identifier string `json:"identifier"`
  Name string `json:"name"`
  ThermostatRev string `json:"thermostatRev"`
  IsRegistered bool `json:"isRegistered"`
  ModelNumber string `json:"modelNumber"`
  Brand string `json:"brand"`
  Features string `json:"features"`
  LastModified string `json:"lastModified"`
  ThermostatTime string `json:"thermostatTime"`
  UTCTime string `json:"utcTime"`
  Alerts []Alert `json:"alerts"`
  // Reminders []ThermostatReminder `json:"reminders"` // documentation not clear on this...
  Settings Settings `json:"settings"`
  Runtime Runtime `json:"runtime"`
  ExtendedRuntime ExtendedRuntime `json:"extendedRuntime"`
  Electricty Electricity `json:"electricity"`
  Devices []Device `json:"devices"`
  Location Location `json:"location"`
  Technician Technician `json:"technician"`
  Utility Utility `json:"utility"`
  Management Management `json:"management"`
  Weather Weather `json:"weather"`
  Events []Event `json:"events"`
  Program Program `json:"program"`
  HouseDetails HouseDetails `json:"houseDetails"`
  // OEMCFG ThermostatOemCfg `json:"oemCfg"` // documentation not clear on this...
  EquipmentStatus string `json:"equipmentStatus"`
  NotificationSettings NotificationSettings `json:"notificationSettings"`
  // Privacy ThermostatPrivacy `json:"privacy"` // documentation not clear on this...
  Version Version `json:"version"`
  SecuritySettings SecuritySettings `json:"securitySettings"`
  RemoteSensors []RemoteSensor `json:"remoteSensors"`
}

type Status struct {
  Code int `json:"code"`
  Message string `json:"message"`
}

type Alert struct {
  AcknowledgeRef string `json:"acknowledgeRef"`
  Date string `json:"date"`
  Time string `json:"time"`
  Severity string `json:"severity"`
  Text string `json:"text"`
  AlertNumber int `json:"alertNumber"`
  AlertType string `json:"alertType"`
  IsOperatorAlert bool `json:"isOperatorAlert"`
  Reminder string `json:"reminder"`
  ShowIdt bool `json:"showIdt"`
  ShowWeb bool `json:"showWeb"`
  SendEmail bool `json:"sendEmail"`
  Acknowledgement string `json:"acknowledgement"`
  RemindMeLater bool `json:"remindMeLater"`
  ThermostatIdentifier string `json:"thermostatIdentifier"`
  NotificationType string`json:"notificationType"`
}

type Settings struct {
  HVACMode string `json::"hvacMode"`
  LastServiceDate string `json::"lastServiceDate"`
  ServiceRemindMe bool `json::"serviceRemindMe"`
  MonthsBetweenService int `json::"monthsBetweenService"`
  RemindMeDate string `json::"remindMeDate"`
  Vent string `json::"vent"`
  VentilatorMinOnTime int `json::"ventilatorMinOnTime"`
  ServiceRemindTechnician bool `json:"serviceRemindTechnician"`
  EILocation string `json:"eiLocation"`
  ColdTempAlert int `json:"coldTempAlert"`
  ColdTempAlertEnabled bool `json:"coldTempAlertEnalbed"`
  HotTempAlert int `json:"hotTempAlert"`
  HotTempAlertEnabled bool `json:"hotTempAlertEnalbed"`
  CoolStages int `json:"coolStages"`
  HeatStages int `json:"heatStages"`
  MaxSetBack int `json:"maxSetBack"`
  MaxSetForward int `json:"maxSetForward"`
  QuickSaveSetBack int `json:"quickSaveSetBack"`
  HasHeatPump bool `json:"hasHeatPump"`
  HasForcedAir bool `json:"hasForcedAir"`
  HasBoiler bool `json:"hasBoiler"`
  HasHumidifier bool `json:"hasHumidifier"`
  HasErv bool `json:"hasErv"`
  HasHrv bool `json:"hasHrv"`
  CondensationAvoid bool `json:"condensationAvoid"`
  UseCelcius bool `json:"useCelcius"`
  UseTimeFormat12 bool `json:"UseTimeFormat12"`
  Locale string `json:"locale"`
  HumidifierMode string `json:"humidifierMode"`
  BacklightOnIntensity int `json:"backlightOnIntensity"`
  BacklightSleepIntensity int `json:"backlightSleepIntensity"`
  BacklightOff int `json:"backlightOff"`
  SoundTickVolume int `json:"soundTickVolume"`
  SoundAlertVolume int `json:"soundAlertVolume"`
  CompressorProtectionMinTime int `json:"compressorProtectionMinTime"`
  CompressorProtectionMinTemp int `json:"compressorProtectionMinTemp"`
  Stage1HeatingDifferentialTemp int `json:"stage1HeatingDifferentialTemp"`
  Stage1CoolingDifferentialTemp int `json:"stage1CoolingDifferentialTemp"`
  Stage1HeatingDissipationTime int `json:"stage1HeatingDissipationTime"`
  Stage1CoolingDissipationTime int `json:"stage1CoolingDissipationTime"`
  HeatPumpReversalOnCool bool `json:"heatPumpReversalOnCool"`
  FanControlRequired bool `json:"FanControlRequired"`
  FanMinOnTime int `json:"fanMinOnTime"`
  HeatCoolMinDelta int `json:"heatCoolMinDelta"`
  TempCorrection int `json:"tempCorrection"`
  HoldAction string `json:"HoldAction"`
  HeatPumpGroundWater bool `json:"heatPumpGroundWater"`
  HasElectric bool `json:"hasElectric"`
  HasDehumidifier bool `json:"hasDehumidifier"`
  DehumidifierMode string `json:"dehumidifierMode"`
  DehumidifierLevel int `json:"dehumidifierLevel"`
  DehumidifyWithAC bool `json:"dehumidifyWithAC"`
  DehumidifyOvercoolOffset bool `json:"dehumidifyOvercoolOffset"`
  AutoHeatCoolFeatureEnabled bool `json:"autoHeatCoolFeatureEnabled"`
  WifiOnlineAlert bool `json:"wifiOfflineAlert"`
  HeatMinTemp int `json:"heatMinTemp"`
  HeatMaxTemp int `json:"heatMaxTemp"`
  CoolMinTemp int `json:"coolMinTemp"`
  CoolMaxTemp int `json:"coolMaxTemp"`
  HeatRangeHigh int `json:"heatRangeHigh"`
  HeatRangeLow int `json:"heatRangeLow"`
  CoolRangeHigh int `json:"coolRangeHigh"`
  CoolRangeLow int `json:"coolRangeLow"`
  UserAccessCode string `json:"userAccessCode"`
  UserAccessSetting int `json:"userAccessSetting"`
  AuxRuntimeAlert int `json:"auxRuntimeAlert"`
  AuxOutdoorTempAlert int `json:"auxOutdoorTempAlert"`
  AuxMaxOutdoorTemp int `json:"auxMaxOutdoorTemp"`
  AuxRuntimeAlertNotify bool `json:"auxRuntimeAlertNotfify"`
  AuxRuntimeAlertNotifyTechnician bool `json:"auxRuntimeAlertNotifyTechnician"`
  AuxOutdoorTempAlertNotifyTechnician bool `json:"auxOutdoorTempAlertNotifyTechnician"`
  DisablePreHeating bool `json:"disablePreHeating"`
  DisablePreCooling bool `json:"disablePreCooling"`
  InstallerCodeRequired bool `json:"installerCodeRequired"`
  DrAccept string `json:"drAccept"`
  IsRentalProperty bool `json:"isRentalProperty"`
  UseZoneController bool `json:"useZoneController"`
  RandomStartDelayCool int `json:"randomStartDelayCool"`
  RandomStartDelayHeat int `json:"randomStartDelayHeat"`
  HumidityHighAlert int `json:"humidityHighAlert"`
  HumidityLowAlert int `json:"humidityLowAlert"`
  DisableHeatPumpAlerts bool `json:"disableHeatPumpAlerts"`
  DisableAlertsOnIdt bool `json:"disableAlertsOnIdt"`
  HumidityAlertNotfiy bool `json:"humidityAlertNotfiy"`
  HumidityAlertNotifyTechnician bool `json:"humidityAlertNotifyTechnician"`
  MonthlyElectrictyBillLimit int `json:"monthlyElectrictyBillLimit"`
  EnableElectrictyBillAlert bool `json:"enableElectrictyBillAlert"`
  EnableProjectedElectricityBillAlert bool `json:"enableProjectedElectricityBillAlert"`
  ElectricityBillingDayOfMonth int `json:"electricityBillingDayOfMonth"`
  ElectrictyBillCycleMonths int `json:"electrictyBillCycleMonths"`
  ElectrictyBillStartMonth int `json:"electrictyBillStartMonth"`
  VentilatorMinOnTimeHome int `json:"ventilatorMinOnTimeHome"`
  VentilatorMinOnTimeAway int `json:"ventilatorMinOnTimeAway"`
  BacklightOffDuringSleep bool `json:"backlightOffDuringSleep"`
  AutoAway bool `json:"autoAway"`
  SmartCiculation bool `json:"smartCirculation"`
  FollowMeComfort bool `json:"followMeComfort"`
  VentilatorType string `json:"ventilatorType"`
  IsVentilatorTimerOn bool `json:"isVentilatorTimerOn"`
  VentilatorOffDateTime string `json:"ventilatorOffDateTime"`
  HasUVFilter bool `json:"hasUVFilter"`
  CoolingLockout bool `json:"coolingLockout"`
  VentilatorFreeCooling bool `json:"ventilatorFreeCooling"`
  DehumidifyWhenHeating bool `json:"dehumidifyWhenHeating"`
  VentilatorDehumidify bool `json:"ventilatorDehumidify"`
  GroupRef string `json:"groupRef"`
  GroupName string `json:"groupName"`
  GroupSetting int `json:"groupSetting"`
}

type Runtime struct {
  RuntimeRev string `json:"runtimeRev"`
  Connected bool `json:"connected"`
  FirstConnected string `json:"firstConnected"`
  ConnectDateTime string `json:"connectDateTime"`
  DisconnectDateTime string `json:"disconnectDateTime"`
  LastModified string `json:"lastModified"`
  LastStatusModified string `json:"lastStatusModified"`
  RuntimeDate string `json:"runtimeDate"`
  RuntimeInterval int `json:"runtimeInterval"`
  ActualTemperature int `json:"actualTemperature"`
  ActualHumidity int `json:"actualHumidity"`
  DesiredHeat int `json:"desiredHeat"`
  DesiredCool int `json:"desiredCool"`
  DesiredHumidity int `json:"desiredHumidity"`
  DesiredDehumidity int `json:"DesiredDehumidity"`
  DesiredFanMode string `json:"desiredFanMode"`
  DesiredHeatRange []int `json:"desiredHeatRange"`
  DesiredCoolRange []int `json:"desiredCoolRange"`
}

type ExtendedRuntime struct {
  LastReadingTimestamp string `json:"lastReadingTimestamp"`
  RuntimeDate string `json:"runtimeDate"`
  RuntimeInterval int `json:"runtimeInterval"`
  ActualTemperature []int `json:"actualTemperature"`
  ActualHumidiy []int `json:"actualHumidity"`
  DesiredHeat []int `json:"desiredHeat"`
  DesiredCool []int `json:"desiredCool"`
  DesiredHumidity []int `json:"desiredHumidity"`
  DesiredDehumidity []int `json:"DesiredDehumidity"`
  DRMOffset []int `json:"drmOffset"`
  HVACMode []string `json:"hvacMode"`
  HeatPump1 []int `json:"heatPump1"`
  HeatPump2 []int `json:"heatPump2"`
  AuxHeat1 []int `json:"auxHeat1"`
  AuxHeat2 []int `json:"auxHeat2"`
  AuxHeat3 []int `json:"auxHeat3"`
  Cool1 []int `json:"cool1"`
  Cool2 []int `json:"cool2"`
  Fan []int `json:"fan"`
  Humidifier []int `json:"humidifier"`
  Dehumidifier []int `json:"dehumidifier`
  Economizer []int `json:"economizer"`
  Ventilator []int `json:"ventilator"`
  CurrentElectrictyBool int `json:"currentElectricityBool"`
  ProjectedElectricityBill int `json:"projectedElectricityBill"`
}

type Electricity struct {
  devices []ElectricityDevice `json:"devices"`
}

type ElectricityDevice struct {
  Tiers []ElectricityTier `json:"tiers"`
  LastUpdate string `json:"lastUpdate"`
  Cost []string `json:"cost"`
  Consumption []string `json:"consumption"`
}

type ElectricityTier struct {
  Name string `json:"name"`
  Consumption string `json:"consumption"`
  Cost string `json:"cost"`
}

type Device struct {
  DeviceId int `json:"deviceId"`
  Name string `json:"name"`
  Sensors []Sensor `json:"sensors"`
  Outputs []Output `json:"outputs"`
}

type Sensor struct {
  Name string `json:"name"`
  Manufacturer string `json:"manufacturer"`
  Model string `json:"model"`
  Zone int `json:"zone"`
  SensorId int `json:"sensorId"`
  Type string `json:"type"`
  Usage string `json:"usage"`
  NumberOfBits int `json:"numberOfBits"`
  BConstant int `json:"bconstant"`
  ThermistorSize int `json:"thermistorSize"`
  TempCorrection int `json:"tempCorrection"`
  Gain int `json:"gain"`
  MaxVoltage int `json:"maxVoltage"`
  Multiplier int `json:"multiplier"`
  States []State `json:"states"`
}

type State struct {
  MaxValue int `json:"maxValue"`
  MinValue int `json:"minValue"`
  Type string `json:"string"`
  Actions []Action `json:"actions"`
}

type Action struct {
  Type string `json:"type"`
  SendAlert bool `json:"sendAlert"`
  SendUpdate bool `json:"sendUpdate"`
  ActivationDelay int `json:"activationDelay"`
  DeactivationDelay int `json:"deactivationDelay"`
  MinActionDuration int `json:"minActionDuration"`
  HeatAdjustTemp int `json:"heatAdjustTemp"`
  CoolAdjustTemp int `json:"coolAdjustTemp"`
  ActivateRelay string `json:"activateRelay"`
  ActivateRelayOpen bool `json:"activateRelayOpen"`
}

type Output struct {
  Name string `json:"name"`
  Zone int `json:"zone"`
  OutputId int `json:"outputId"`
  Type string `json:"type"`
  SendUpdate bool `json:"sendUpdate"`
  ActiveClosed bool `json:"activeClosed"`
  ActivationTime int `json:"activationTime"`
  DeactivationTime int `json:"deactivationTime"`
}

type Location struct {
  TimeZoneOffsetMinutes int `json:"timeZoneOffsetMinutes"`
  TimeZone string `json:"timeZone"`
  IsDaylightSaving bool `json:"isDaylightSaving"`
  StreetAddress string `json:"streetAddress"`
  City string `json:"city"`
  ProvinceState string `json:"provinceState"`
  Country string `json:"country"`
  PostalCode string `json:"postalCode"`
  PhoneNubmer string `json:"phoneNumber"`
  MapCoordinates string `json:"mapCoordinates"`
}

type Technician struct {
  ContractorRef string `json:"contractorRef"`
  Name string `json:"name"`
  Phone string `json:"phone"`
  StreetAddress string `json:"streetAddress"`
  City string `json:"city"`
  ProvinceState string `json:"provinceState"`
  Country string `json:"country"`
  PostalCode string `json:"postalCode"`
  Email string `json:"email"`
  Web string `json:"web"`
}

type Utility struct {
  Name string `json:"name"`
  Phone string `json:"phone"`
  Email string `json:"email"`
  Web string `json:"web"`
}

type Management struct {
  AdministrativeContact string `json:"administrativeContact"`
  BillingContact string `json:"billingContact"`
  Name string `json:"name"`
  Phone string `json:"phone"`
  Email string `json:"email"`
  Web string `json:"web"`
  ShowAlertIdt bool `json:"showAlertIdt"`
  ShowAlertWeb bool `json:"showAlertWeb"`
}

type Weather struct {
  Timestamp string `json:"timestamp"`
  WeatherStation string `json:"weatherStation"`
  Forecasts []WeatherForecast `json:"forecasts"`
}

type WeatherForecast struct {
  WeatherSymbol int `json:"weatherSymbol"`
  DateTime string `json:"dateTime"`
  Condition string `json:"condition"`
  Temperature int `json:"temperature"`
  Pressure int `json:"pressure"`
  RelativeHumidity int `json:"relativeHumidity"`
  Dewpoint int `json:"dewpoint"`
  Visibility int `json:"visibility"`
  WindSpeed int `json:"windSpeed"`
  WindGust int `json:"windGust"`
  WindDirection string `json:"windDirection"`
  WindBearing int `json:"windBearing"`
  Pop int `json:"pop"`
  TempHigh int `json:"tempHigh"`
  TempLow int `json:"tempLow"`
  Sky int `json:"sky"`
}

type Event struct {
  Type string `json:"type"`
  Name string `json:"name"`
  Running bool `json:"running"`
  StartDate string `json:"startDate"`
  StartTime string `json:"startTime"`
  EndDate string `json:"endDate"`
  EndTime string `json:"endTime"`
  IsOccupied bool `json:"isOccupied"`
  IsCoolOff bool `json:"isCoolOff"`
  IsHeatOff bool `json:"isHeatOff"`
  CoolHoldTemp int `json:"coolHoldTemp"`
  HeatHoldTemp int `json:"heatHoldTemp"`
  Fan string `json:"fan"`
  Vent string `json:"vent"`
  VentilatorMinOnTime int `json:"ventilatorMinOnTime"`
  IsOptional bool `json:"isOptional"`
  IsTemperatureRelative bool `json:"isTemperatureRelative"`
  CoolRelativeTemp int `json:"coolRelativeTemp"`
  HasRelativeTemp int `json:"hasRelativeTemp"`
  IsTemperatureAbsolute bool `json:"isTemperatureAbsolute"`
  DutyCyclePercentage int `json:"dutyCyclePercentage"`
  FanMinOnTIme int `json:"fanMinOnTime"`
  OccupiedSensorActive bool `json:"occupiedSensorActive"`
  UnoccupiedSensorActive bool `json:"unoccupiedSensorActive"`
  DrRampupTemp int `json:"drRampUpTemp"`
  DrRampUpTime int `json:"drRampUpTime"`
  LinkRef string `json:"linkRef"`
  HoldClimateRef string `json:"holdClimateRef"`
}

type Program struct {
  schedule map[string]string `json:"schedule"`
  Climates []Climate `json:"climates"`
  CurrentClimateRef string `json:"currentClimateRef"`
}

type Climate struct {
  Name string `json:"name"`
  ClimateRef string `json:"climateRef"`
  IsOccupied bool `json:"isOccupied"`
  IsOptimized bool `json:"isOptimized"`
  CoolFan string `json:"coolFan"`
  HeatFan string `json:"heatFan"`
  Vent string `json:"vent"`
  VentilatorMinOnTime int `json:"ventilatorMinOnTime"`
  Owner string `json:"owner"`
  Type string `json:"type"`
  Colour int `json:"colour"`
  CoolTemp int `json:"coolTemp"`
  HeatTemp int `json:"heatTemp"`
  Sensors []RemoteSensor `json:"sensors"`
}

type RemoteSensor struct {
  Id string `json:"id"`
  Name string `json:"name"`
  Type string `json:"type"`
  Code string `json:"code"`
  InUse bool `json:"inUse"`
  Capability []RemoteSensorCapability `json:"capability"`
}

type RemoteSensorCapability struct {
  Id string `json:"id"`
  Type string `json:"type"`
  Value string `json:"value"`
}

type HouseDetails struct {
  Style string `json:"style"`
  Size int `json:"size"`
  NumberOfFloors int `json:"numberOfFloors"`
  NumberOfRooms int `json:"numberOfRooms"`
  NumberOfOccupants int `json:"numberOfOccupants"`
  Age int `json:"age"`
  WindowEfficiency int `json:"windowEfficiency"`
}

type NotificationSettings struct {
  EmailAddresses []string `json:"emailAddresses"`
  EmailNotificationsEnabled bool `json:"emailNotificationsEnabled"`
  Equipment []EquipmentSetting `json:"equipment"`
  General []GeneralSetting `json:"general"`
  Limit []LimitSetting `json:"limit"`
}

type EquipmentSetting struct {
  FilterLastChanged string `json:"filterLastChanged"`
  FilterLife int `json:"filterLife"`
  FilterLifeUnits string `json:"filterLifeUnits"`
  RemindMeDate string `json:"remindMeDate"`
  Enabled bool `json:"enabled"`
  Type string `json:"type"`
  RemindTechnician bool `json:"remindTechnician"`
}

type GeneralSetting struct {
  Enabled bool `json:"enabled"`
  Type string `json:"type"`
  RemindTechnician bool `json:"remindTechnician"`
}

type LimitSetting struct {
  Limit int `json:"limit"`
  Enabled bool `json:"enabled"`
  Type string `json:"type"`
  RemindTechnician bool `json:"remindTechnician"`
}

type Version struct {
  ThermostatFirmwareVersion string `json:"thermostatFirmwareVersion"`
}

type SecuritySettings struct {
  UserAccessCode string `json:"userAccessCode"`
  AllUserAccess bool `json:"allUserAccess"`
  ProgramAccess bool `json:"programAccess"`
  DetailsAccess bool `json:"detailsAccess"`
  QuickSaveAccess bool `json:"quickSaveAccess"`
  VacationAccess bool `json:"vacationAccess"`
}

func (e *Ecobee) GetThermostats(selection string) (thermostats *Thermostats, err error) {
  url := fmt.Sprintf("https://api.ecobee.com/1/thermostat?json=%s", url.QueryEscape(selection))
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return thermostats, err
  }
  req.Header.Set("Content-Type", "application/json;charset=UTF-8")
  req.Header.Set("Authorization", "Bearer " + e.AccessToken)
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
      return thermostats, err
  }
  defer resp.Body.Close()
  err = json.NewDecoder(resp.Body).Decode(&thermostats)
  if err != nil {
    return thermostats, err
  }
  return thermostats, nil
}