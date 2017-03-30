package ecobee

import(
  "net/http"
  "net/url"
  "fmt"
  "encoding/json"
)

type Report struct {
  StartDate string `json:"startDate"`
  StartInterval int `json:"startInterval"`
  EndDate string `json:"endDate"`
  EndInterval int `json:"endInterval"`
  Columns string `json:"columns"`
  ReportList []RuntimeReport `json:"reportList"`
  SensorList []RuntimeSensorReport `json:"sensorList"`
}

type RuntimeReport struct {
  ThermostatIdentifier string `json:"thermostatIdentifier"`
  RowCount int `json:"rowCount"`
  RowList []string `json:"rowList"`
}

type RuntimeSensorReport struct {
  ThermostatIdentifier string `json:"thermostatIdentifier"`
  Sensors []RuntimeSensorMetadata `json:"sensors"`
  Columns []string `json:"columns"`
  Data []string `json:"data"`
}

type RuntimeSensorMetadata struct {
  SensorId string `json:"sensorId"`
  SensorName string `json:"sensorName"`
  SensorType string `json:"sensorType"`
  SensorUsage string `json:"sensorUsage"`
}

type ReportRequest struct {
  StartDate string `json:"start_date"`
  EndDate string `json:"end_date"`
  AuxHeat1 bool `json:"aux_heat_1"`
  AuxHeat2 bool `json:"aux_heat_2"`
  AuxHeat3 bool `json:"aux_heat_3"`
  CompCool1 bool `json:"comp_cool_1"`
  CompCool2 bool `json:"comp_cool_2"`
  CompHeat1 bool `json:"comp_heat_1"`
  CompHeat2 bool `json:"comp_heat_2"`
  Dehumidifier bool `json:"dehumidifier"`
  DMOffset bool `json:"dm_offset"`
  Economizer bool `json:"economizer"`
  Fan bool `json:"fan"`
  Humidifier bool `json:"humidifier"`
  HVACMode bool `json:"hvac_mode"`
  OutdoorHumidity bool `json:"outdoor_humidity"`
  OutdoorTemp bool `json:"outdoor_temp"`
  Sky bool `json:"sky"`
  Ventilator bool `json:"ventilator"`
  Wind bool `json:"wind"`
  ZoneAveTemp bool `json:"zone_ave_temp"`
  ZoneCalendarEvent bool `json:"zone_calendar_event"`
  ZoneClimate bool `json:"zone_climate"`
  ZoneCoolTemp bool `json:"zone_cool_temp"`
  ZoneHeatTemp bool `json:"zone_heat_temp"`
  ZoneHumidity bool `json:"zone_humidity"`
  ZoneHumidityHigh bool `json:"zone_humidity_high"`
  ZoneHumidityLow bool `json:"zone_humidity_low"`
  ZoneHVACMode bool `json:"zone_hvac_mode"`
  ZoneOccupancy bool `json:"zone_occupancy"`
}

func (e *Ecobee) GetReport(body string) (report *Report, err error) {
  url := fmt.Sprintf("https://api.ecobee.com/1/runtimeReport?format=json&body=%s", url.QueryEscape(body))
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return report, err
  }
  req.Header.Set("Content-Type", "application/json;charset=UTF-8")
  req.Header.Set("Authorization", "Bearer " + e.AccessToken)
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
      return report, err
  }
  defer resp.Body.Close()
  err = json.NewDecoder(resp.Body).Decode(&report)
  if err != nil {
    return report, err
  }
  return report, nil
}

func (s *ReportRequest) BuildReportRequest(thermostatId string) string {
  columns := ""
  if s.AuxHeat1 { columns = columns + "auxHeat1," }
  if s.AuxHeat2 { columns = columns + "auxHeat2," }
  if s.AuxHeat3 { columns = columns + "auxHeat3," }
  if s.CompCool1 { columns = columns + "compCool1," }
  if s.CompCool2 { columns = columns + "compCool2," }
  if s.CompHeat1 { columns = columns + "compHeat1," }
  if s.CompHeat2 { columns = columns + "compHeat2," }
  if s.Dehumidifier { columns = columns + "dehumidifier," }
  if s.DMOffset { columns = columns + "dmOffset," }
  if s.Economizer { columns = columns + "economizer," }
  if s.Fan { columns = columns + "fan," }
  if s.Humidifier { columns = columns + "humidifier," }
  if s.HVACMode { columns = columns + "hvacMode," }
  if s.OutdoorHumidity { columns = columns + "outdoorHumidity," }
  if s.OutdoorTemp { columns = columns + "outdoorTemp," }
  if s.Sky { columns = columns + "sky," }
  if s.Ventilator { columns = columns + "ventilator," }
  if s.Wind { columns = columns + "wind," }
  if s.ZoneAveTemp { columns = columns + "zonAveTemp," }
  if s.ZoneCalendarEvent { columns = columns + "zoneCalendarEvent," }
  if s.ZoneClimate { columns = columns + "zoneClimate," }
  if s.ZoneCoolTemp { columns = columns + "zoneCoolTemp," }
  if s.ZoneHeatTemp { columns = columns + "zoneHeatTemp," }
  if s.ZoneHumidity { columns = columns + "zoneHumidity," }
  if s.ZoneHumidityHigh { columns = columns + "zoneHumidityHigh," }
  if s.ZoneHumidityLow { columns = columns + "zoneHumidityLow," }
  if s.ZoneHVACMode { columns = columns + "zoneHvacMode," }
  if s.ZoneOccupancy { columns = columns + "zoneOccupancy" }

  return fmt.Sprintf(`{
    "startDate": %s,
    "endDate": %s,
    "columns": %s,
    "includeSensors": "true",
    "selection":{
      "selectionType": "thermostats",
      "selectionMatch": %s
    }
  }`,
  s.StartDate,
  s.EndDate,
  columns,
  thermostatId,)
}