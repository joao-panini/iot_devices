# iot_devices
psql -h localhost -U postgres -d postgres

password: postgres

CREATE DATABASE iotdb;

migrate -path pkg/repository/migrations -database postgres://postgres:postgres@localhost:5432/iotdb?sslmode=disable up
