package api

import (
	"encoding/json"
	"net/http"
)

type ListRequest struct {
	DeviceType string `json:"device_type"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	Limit      int    `json:"limit"`
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req ListRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	startT, endT, err := req.ParseDates()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.DeviceType == "temperature" {
		resp, err := h.tempService.ListTemperatureData(ctx, startT, endT, req.Limit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Set the response content type to application/json
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	if req.DeviceType == "velocity" {
		resp, err := h.velService.ListVelocityData(ctx, startT, endT, req.Limit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Set the response content type to application/json
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	if req.DeviceType == "traffic" {
		resp, err := h.trafficService.ListTrafficData(ctx, startT, endT, req.Limit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Set the response content type to application/json
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

}
