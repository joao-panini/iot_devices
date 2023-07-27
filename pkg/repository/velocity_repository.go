package repository

import (
	"context"
	"database/sql"
	"iot_devices/pkg/models"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type VelocityRepository struct {
	db *sql.DB
}

func NewVelocityRepository(db *sql.DB) *VelocityRepository {
	return &VelocityRepository{db: db}
}

func (r *VelocityRepository) AddVelocityData(ctx context.Context, vel models.Velocity) (*models.Velocity, error) {
	sql := `INSERT INTO velocity_sensor (device_id, velocity, unit,timestamp)
	 VALUES ($1, $2, $3, $4)
	 RETURNING id,device_id,Velocity,unit,timestamp`

	result := &models.Velocity{}
	err := r.db.QueryRowContext(ctx, sql, vel.DeviceID, vel.Velocity, vel.Unit, vel.Timestamp).Scan(&result.ID,
		&result.DeviceID, &result.Velocity, &result.Unit, &result.Timestamp)

	return result, err
}

func (r *VelocityRepository) ListVelocityData(ctx context.Context, start time.Time, end time.Time, limit int) ([]models.Velocity, error) {
	rows, err := r.db.QueryContext(ctx,
		"SELECT id, device_id, velocity,unit, timestamp FROM velocity_sensor WHERE timestamp BETWEEN '$1' AND '$2' ORDER BY timestamp DESC LIMIT $3",
		start, end, limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var velList []models.Velocity
	for rows.Next() {
		var vel models.Velocity
		if err := rows.Scan(&vel.ID, &vel.DeviceID, &vel.Velocity, &vel.Unit, &vel.Timestamp); err != nil {
			return nil, err
		}
		velList = append(velList, vel)
	}

	return velList, nil
}
