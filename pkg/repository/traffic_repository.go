package repository

import (
	"context"
	"database/sql"
	"iot_devices/pkg/models"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type TrafficRepository struct {
	db *sql.DB
}

func NewTrafficRepository(db *sql.DB) *TrafficRepository {
	return &TrafficRepository{db: db}
}

func (r *TrafficRepository) AddTrafficData(ctx context.Context, traffic models.Traffic) (*models.Traffic, error) {
	sql := `INSERT INTO traffic_sensor (device_id, plate_number, timestamp)
	 VALUES ($1, $2, $3)
	 RETURNING id,device_id,plate_number,timestamp`

	result := &models.Traffic{}
	err := r.db.QueryRowContext(ctx, sql, traffic.DeviceID, traffic.PlateNumber, traffic.Timestamp).
		Scan(&result.ID, &result.DeviceID, &result.PlateNumber, &result.Timestamp)

	return result, err
}

func (r *TrafficRepository) ListTrafficData(ctx context.Context, start time.Time, end time.Time, limit int) ([]models.Traffic, error) {
	rows, err := r.db.QueryContext(ctx,
		"SELECT id, device_id, plate_number, timestamp FROM traffic_sensor WHERE timestamp BETWEEN $1 AND $2 ORDER BY timestamp DESC LIMIT $3",
		start, end, limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trafficList []models.Traffic
	for rows.Next() {
		var traffic models.Traffic
		if err := rows.Scan(&traffic.ID, &traffic.DeviceID, &traffic.PlateNumber, &traffic.Timestamp); err != nil {
			return nil, err
		}
		trafficList = append(trafficList, traffic)
	}

	return trafficList, nil
}
