package models

import "time"

type Robot struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	RobotClass  string    `json:"robot_class"`
	Status      string    `json:"status"`
	CurrentZone string    `json:"current_zone"`
	TargetZone  string    `json:"target_zone"`
	CreatedAt   time.Time `json:"created_at"`
}

type TelemetryInput struct {
	RobotID            string    `json:"robot_id" binding:"required"`
	RobotClass         string    `json:"robot_class" binding:"required"`
	Timestamp          time.Time `json:"timestamp" binding:"required"`
	X                  float64   `json:"x" binding:"required"`
	Y                  float64   `json:"y" binding:"required"`
	CurrentZone        string    `json:"current_zone" binding:"required"`
	TargetZone         string    `json:"target_zone" binding:"required"`
	Speed              float64   `json:"speed"`
	BatteryLevel       int       `json:"battery_level"`
	Temperature        float64   `json:"temperature"`
	MotorLoad          float64   `json:"motor_load"`
	TaskStatus         string    `json:"task_status"`
	SignalStrength     int       `json:"signal_strength"`
	ObstacleDetected   bool      `json:"obstacle_detected"`
	DistanceToObstacle float64   `json:"distance_to_obstacle"`
	ErrorCode          *string   `json:"error_code"`
}

type TelemetryRecord struct {
	TelemetryInput

	ConnectionStatus string  `json:"connection_status"`
	RouteDeviation   float64 `json:"route_deviation"`
	IsStuck          bool    `json:"is_stuck"`
}