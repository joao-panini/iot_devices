package api

import (
	"iot_devices/pkg/service"
	"net/http"
)

type Handler struct {
	tempService    service.TemperatureService
	velService     service.VelocityService
	trafficService service.TrafficService
}

func NewHandler(tempService *service.TemperatureService, velService *service.VelocityService, trafficService *service.TrafficService) *Handler {
	return &Handler{
		tempService:    *tempService,
		velService:     *velService,
		trafficService: *trafficService,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/devices":
		switch r.Method {
		case http.MethodGet:
			h.List(w, r)
		case http.MethodPost:
			h.Create(w, r)
		default:
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		}
	default:
		http.NotFound(w, r)
	}
}
