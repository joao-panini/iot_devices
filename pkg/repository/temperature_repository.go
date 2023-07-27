package repository

import (
	"context"
	"database/sql"
	"iot_devices/pkg/models"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type TemperatureRepository struct {
	db *sql.DB
}

func NewTemperatureRepository(db *sql.DB) *TemperatureRepository {
	return &TemperatureRepository{db: db}
}

func (r *TemperatureRepository) AddTemperatureData(ctx context.Context, temp models.Temperature) (*models.Temperature, error) {
	sql := `INSERT INTO temperature_sensor (device_id, temperature, unit,timestamp)
	 VALUES ($1, $2, $3, $4)
	 RETURNING id,device_id,temperature,unit,timestamp`

	result := &models.Temperature{}
	err := r.db.QueryRowContext(ctx, sql, temp.DeviceID, temp.Temperature, temp.Unit, temp.Timestamp).
		Scan(&result.ID, &result.DeviceID, &result.Temperature, &result.Unit, &result.Timestamp)

	return result, err
}

func (r *TemperatureRepository) ListTemperatureData(ctx context.Context, start time.Time, end time.Time, limit int) ([]models.Temperature, error) {
	rows, err := r.db.QueryContext(ctx,
		"SELECT id, device_id, temperature,unit, timestamp FROM temperature_sensor WHERE timestamp BETWEEN $1 AND $2 ORDER BY timestamp DESC LIMIT $3",
		start, end, limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tempList []models.Temperature
	for rows.Next() {
		var temp models.Temperature
		if err := rows.Scan(&temp.ID, &temp.DeviceID, &temp.Temperature, &temp.Unit, &temp.Timestamp); err != nil {
			return nil, err
		}
		tempList = append(tempList, temp)
	}

	return tempList, nil
}
