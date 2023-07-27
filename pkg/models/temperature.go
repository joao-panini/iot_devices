package models

type Temperature struct {
	Device
	Temperature string `json:"temperature"`
	Unit        string `json:"unit"`
}
