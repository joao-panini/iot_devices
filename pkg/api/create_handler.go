package api

import (
	"encoding/json"
	"iot_devices/pkg/models"
	"net/http"
	"time"
)

type CreateRequest struct {
	DeviceID    int    `json:"device_id"`
	Temperature string `json:"temperature"`
	PlateNumber string `json:"plate_number"`
	Velocity    string `json:"velocity"`
	Unit        string `json:"unit"`
	Timestamp   string `json:"timestamp"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	timestamp, err := time.Parse("02-01-2006 15:04:05", req.Timestamp)
	if err != nil {
		http.Error(w, "Failed to parse timestamp", http.StatusBadRequest)
		return
	}

	if req.Temperature != "" {
		//temp use case
		tempInput := models.Temperature{
			Device: models.Device{
				DeviceID:  req.DeviceID,
				Timestamp: timestamp,
			},
			Temperature: req.Temperature,
			Unit:        req.Unit,
		}

		res, err := h.tempService.CreateTemperatureData(ctx, tempInput)
		if err != nil {
			http.Error(w, "Failed to save temperature data", http.StatusUnprocessableEntity)
			return
		}
		// Set the response content type to application/json
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
		return
	}

	if req.Velocity != "" {
		//vel use case
		velInput := models.Velocity{
			Device: models.Device{
				DeviceID:  req.DeviceID,
				Timestamp: timestamp,
			},
			Velocity: req.Velocity,
			Unit:     req.Unit,
		}

		res, err := h.velService.CreateVelocityData(ctx, velInput)
		if err != nil {
			http.Error(w, "Failed to save velocity data", http.StatusUnprocessableEntity)
			return
		}
		// Set the response content type to application/json
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
		return
	}

	if req.PlateNumber != "" {
		//vel use case
		trafficInput := models.Traffic{
			Device: models.Device{
				DeviceID:  req.DeviceID,
				Timestamp: timestamp,
			},
			PlateNumber: req.PlateNumber,
		}

		res, err := h.trafficService.CreateTrafficData(ctx, trafficInput)
		if err != nil {
			http.Error(w, "Failed to save traffic data", http.StatusUnprocessableEntity)
			return
		}
		// Set the response content type to application/json
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
		return
	}

}
