package repository

import (
	"context"

	"github.com/Ernuramirzhan/robofleet-os/backend/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TelemetryRepository struct {
	db *pgxpool.Pool
}

func NewTelemetryRepository(db *pgxpool.Pool) *TelemetryRepository {
	return &TelemetryRepository{db: db}
}

func (r *TelemetryRepository) SaveTelemetry(ctx context.Context, record models.TelemetryRecord) (int64, error) {
	var telemetryID int64

	err := r.db.QueryRow(ctx, `
		INSERT INTO telemetry_records (
			robot_id,
			robot_class,
			timestamp,
			x,
			y,
			current_zone,
			target_zone,
			speed,
			battery_level,
			temperature,
			motor_load,
			task_status,
			signal_strength,
			obstacle_detected,
			distance_to_obstacle,
			error_code,
			connection_status,
			route_deviation,
			is_stuck
		)
		VALUES (
			$1, $2, $3, $4, $5,
			$6, $7, $8, $9, $10,
			$11, $12, $13, $14, $15,
			$16, $17, $18, $19
		)
		RETURNING id
	`,
		record.RobotID,
		record.RobotClass,
		record.Timestamp,
		record.X,
		record.Y,
		record.CurrentZone,
		record.TargetZone,
		record.Speed,
		record.BatteryLevel,
		record.Temperature,
		record.MotorLoad,
		record.TaskStatus,
		record.SignalStrength,
		record.ObstacleDetected,
		record.DistanceToObstacle,
		record.ErrorCode,
		record.ConnectionStatus,
		record.RouteDeviation,
		record.IsStuck,
	).Scan(&telemetryID)

	if err != nil {
		return 0, err
	}

	return telemetryID, nil
}

func (r *TelemetryRepository) GetRecentTelemetry(ctx context.Context, limit int) ([]models.TelemetryRecord, error) {
	if limit <= 0 {
		limit = 20
	}

	rows, err := r.db.Query(ctx, `
		SELECT
			robot_id,
			robot_class,
			timestamp,
			x,
			y,
			current_zone,
			target_zone,
			speed,
			battery_level,
			temperature,
			motor_load,
			task_status,
			signal_strength,
			obstacle_detected,
			distance_to_obstacle,
			error_code,
			connection_status,
			route_deviation,
			is_stuck
		FROM telemetry_records
		ORDER BY timestamp DESC
		LIMIT $1
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.TelemetryRecord

	for rows.Next() {
		var record models.TelemetryRecord

		if err := rows.Scan(
			&record.RobotID,
			&record.RobotClass,
			&record.Timestamp,
			&record.X,
			&record.Y,
			&record.CurrentZone,
			&record.TargetZone,
			&record.Speed,
			&record.BatteryLevel,
			&record.Temperature,
			&record.MotorLoad,
			&record.TaskStatus,
			&record.SignalStrength,
			&record.ObstacleDetected,
			&record.DistanceToObstacle,
			&record.ErrorCode,
			&record.ConnectionStatus,
			&record.RouteDeviation,
			&record.IsStuck,
		); err != nil {
			return nil, err
		}

		result = append(result, record)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *TelemetryRepository) GetTelemetryByRobotID(ctx context.Context, robotID string, limit int) ([]models.TelemetryRecord, error) {
	if limit <= 0 {
		limit = 20
	}

	rows, err := r.db.Query(ctx, `
		SELECT
			robot_id,
			robot_class,
			timestamp,
			x,
			y,
			current_zone,
			target_zone,
			speed,
			battery_level,
			temperature,
			motor_load,
			task_status,
			signal_strength,
			obstacle_detected,
			distance_to_obstacle,
			error_code,
			connection_status,
			route_deviation,
			is_stuck
		FROM telemetry_records
		WHERE robot_id = $1
		ORDER BY timestamp DESC
		LIMIT $2
	`, robotID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.TelemetryRecord

	for rows.Next() {
		var record models.TelemetryRecord

		if err := rows.Scan(
			&record.RobotID,
			&record.RobotClass,
			&record.Timestamp,
			&record.X,
			&record.Y,
			&record.CurrentZone,
			&record.TargetZone,
			&record.Speed,
			&record.BatteryLevel,
			&record.Temperature,
			&record.MotorLoad,
			&record.TaskStatus,
			&record.SignalStrength,
			&record.ObstacleDetected,
			&record.DistanceToObstacle,
			&record.ErrorCode,
			&record.ConnectionStatus,
			&record.RouteDeviation,
			&record.IsStuck,
		); err != nil {
			return nil, err
		}

		result = append(result, record)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
