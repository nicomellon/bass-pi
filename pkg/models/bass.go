package models

type Bass struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Manufacturer_ID  string `json:"manufacturer_id"`
	Launch_year  	 string `json:"launch_year"`
	Image  	 		 string `json:"image"`
}