package models

type Agency struct {
	Id uint8 `json:"id"`
	AgencyId string `json:"agencyId"`
	AgencyUrl string `json:"agencyUrl"`
	AgencyTimezone string `json:"agencyTimezone"`
	AgencyPhone string `json:"agencyPhone"`
	AgencyLang string `json:"agencyLang"`
}