package models

type Velocity struct {
	Device
	Velocity string `json:"velocity"`
	Unit     string `json:"unit"`
}
