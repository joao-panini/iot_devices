package models

type Traffic struct {
	Device
	PlateNumber string `json:"plate_number"`
}
