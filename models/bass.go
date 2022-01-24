package models

import "database/sql"

type Bass struct {
	ID             uint16         `json:"id"`
	ManufacturerID uint16         `json:"manufacturer_id"`
	Name           string         `json:"name"`
	Strings        uint8          `json:"strings"`
	LaunchYear     int            `json:"launch_year"`
	Image          sql.NullString `json:"image"`
}
