package models

import "time"

type Device struct {
	ID        string
	DeviceID  int       `json:"device_id"`
	Timestamp time.Time `json:"timestamp"`
}
