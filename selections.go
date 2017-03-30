package ecobee

import(
  "fmt"
)

type Selections struct {
  IncludeRuntime string `json:"includeRuntime"`
  IncludeExtendedRuntime string `json:"includeExtendedRuntime"`
  IncludeElectricity string `json:"includeElectricity"`
  IncludeSettings string `json:"includeSettings"`
  IncludeLocation string `json:"includeLocation"`
  IncludeProgram string `json:"includeProgram"`
  IncludeEvents string `json:"includeEvents"`
  IncludeDevice string `json:"includeDevice"`
  IncludeTechnician string `json:"includeTechnician"`
  IncludeUtility string `json:"includeUtility"`
  IncludeAlerts string `json:"includeAlerts"`
  IncludeWeather string `json:"includeWeather"`
  IncludeOemConfig string `json:"includeOemConfig"`
  IncludeEquipmentStatus string `json:"includeEquipmentStatus"`
  IncludeNotificationSettings string `json:"includeNotificationSettings"`
  IncludePrivacy string `json:"includePrivacy"`
  IncludeVersion string `json:"includeVersion"`
  IncludeSecuritySettings string `json:"includeSecuritySettings"`
  IncludeSensors string `json:"includeSensors"`
}

func (s *Selections) BuildSelections() string {
  return fmt.Sprintf(`{"selection":{"selectionType":"registered","selectionMatch":"","includeRuntime":%s,"includeExtendedRuntime":%s,"includeElectricity":%s,"includeSettings":%s,"includeLocation":%s,"includeProgram":%s,"includeEvents":%s,"includeDevice":%s,"includeTechnician":%s,"includeUtility":%s,"includeAlerts":%s,"includeWeather":%s,"includeOemConfig":%s,"includeEquipmentStatus":%s,"includeNotificationSettings":%s,"includePrivacy":%s,"includeVersion":%s,"includeSecuritySettings":%s,"includeSensors":%s}}`, 
  s.IncludeRuntime,
  s.IncludeExtendedRuntime,
  s.IncludeElectricity,
  s.IncludeSettings,
  s.IncludeLocation,
  s.IncludeProgram,
  s.IncludeEvents,
  s.IncludeDevice,
  s.IncludeTechnician,
  s.IncludeUtility,
  s.IncludeAlerts,
  s.IncludeWeather,
  s.IncludeOemConfig,
  s.IncludeEquipmentStatus,
  s.IncludeNotificationSettings,
  s.IncludePrivacy,
  s.IncludeVersion,
  s.IncludeSecuritySettings,
  s.IncludeSensors,)
}