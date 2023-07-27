package api

import (
	"errors"
	"time"
)

//Create consts to make sure the fields like unit have pre-defined standard values

// this can be a whole lot better
func (r *CreateRequest) Validate() error {
	if r.Velocity == "" && r.Temperature == "" && r.PlateNumber == "" {
		return errors.New("must pass temperature, velocity or plate data")
	}

	return nil
}

func (r *ListRequest) Validate() error {
	if r.DeviceType == "" {
		return errors.New("must pass device type")
	}

	if r.StartDate == "" {
		return errors.New("must pass start date")
	}

	if r.EndDate == "" {
		return errors.New("must pass end date")
	}

	if r.Limit == 0 {
		return errors.New("must pass valid limit")
	}

	return nil
}

func (r *ListRequest) ParseDates() (time.Time, time.Time, error) {
	// Convert start and end to time.Time
	startTime, err := time.Parse("02-01-2006 15:04:05", r.StartDate)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	//
	endTime, err := time.Parse("02-01-2006 15:04:05", r.EndDate)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return startTime, endTime, nil

}
