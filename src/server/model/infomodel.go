package model

// Info model struct
type Info struct {
	IP             string   `json:"IP"`
	Country        string   `json:"Country"`
	Status         bool     `json:"Status"` //True=Whitelisted, False=Blacklisted
	WhiteCountries []string `json:"WhiteCountries"`
}
