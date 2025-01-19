package models

type Agency struct {
	AgencyId 		string `json:"agencyId"`
	AgencyUrl 		string `json:"agencyUrl"`
	AgencyTimezone	string `json:"agencyTimezone"`
	AgencyPhone 	string `json:"agencyPhone"`
	AgencyLang 		string `json:"agencyLang"`
}