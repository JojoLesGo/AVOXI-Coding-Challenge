package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	model "../model"
	"github.com/oschwald/geoip2-golang"
)

var (
	//user_id stores the api ID for the country lookup
	userid = "464438"
	//MAXMIND_LICENSE_KEY stores the license for the api
	licensekey = "nDUCKqhMgRteJEdF"
)

// infoController setup using MVC architecture to process info requests
type infoController struct {
	infos model.Info
}

// PostOne ensures that posts with an ID are
func (t *infoController) PostOne(w http.ResponseWriter, r *http.Request, id string) {
	http.Error(w, "Not Allowed", http.StatusMethodNotAllowed)
}

// PostAll creates an Info
func (t *infoController) PostAll(w http.ResponseWriter, r *http.Request) {
	var body map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	IP := body["IP"].(string)
	whitelisted := body["WhiteCountries"]
	list := whitelisted.([]interface{})
	fmt.Println(list)
	db, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP(IP)
	record, err := db.Country(ip)
	if err != nil {
		log.Fatal(err)
	}

	body["Country"] = record.Country.Names["en"]
	fmt.Println(body["Country"])
	for _, country := range list {
		if country.(string) == body["Country"] {
			body["Status"] = true
			break
		} else if body["Status"] == nil {
			body["Status"] = false
		}
	}

	json.NewEncoder(w).Encode(body)
	fmt.Println(body)
}

// Options returns the configured server options
func (t *infoController) Options(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "")
}

// Info constructs the infoController
func Info(infos model.Info) *infoController {
	return &infoController{
		infos: infos,
	}
}
