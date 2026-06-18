package alerts

import (
	"fmt"
	"time"

	"github.com/Ernuramirzhan/robofleet-os/backend/internal/models"
)

func EvaluateTelemetry(record models.TelemetryRecord) []models.Alert {
	var result []models.Alert

	if record.BatteryLevel < 10 {
		result = append(result, newAlert(
			record.RobotID,
			"CRITICAL_BATTERY",
			"critical",
			fmt.Sprintf("Robot %s has critically low battery: %d%%", record.RobotID, record.BatteryLevel),
		))
	} else if record.BatteryLevel < 20 {
		result = append(result, newAlert(
			record.RobotID,
			"LOW_BATTERY",
			"warning",
			fmt.Sprintf("Robot %s has low battery: %d%%", record.RobotID, record.BatteryLevel),
		))
	}

	if record.Temperature > 70 {
		result = append(result, newAlert(
			record.RobotID,
			"OVERHEATING",
			"critical",
			fmt.Sprintf("Robot %s is overheating: %.1f°C", record.RobotID, record.Temperature),
		))
	} else if record.Temperature > 60 {
		result = append(result, newAlert(
			record.RobotID,
			"HIGH_TEMPERATURE",
			"warning",
			fmt.Sprintf("Robot %s has high temperature: %.1f°C", record.RobotID, record.Temperature),
		))
	}

	if record.MotorLoad > 95 {
		result = append(result, newAlert(
			record.RobotID,
			"MOTOR_OVERLOAD",
			"critical",
			fmt.Sprintf("Robot %s has motor overload: %.1f%%", record.RobotID, record.MotorLoad),
		))
	} else if record.MotorLoad > 85 {
		result = append(result, newAlert(
			record.RobotID,
			"HIGH_MOTOR_LOAD",
			"warning",
			fmt.Sprintf("Robot %s has high motor load: %.1f%%", record.RobotID, record.MotorLoad),
		))
	}

	if record.ConnectionStatus == "offline" {
		result = append(result, newAlert(
			record.RobotID,
			"ROBOT_OFFLINE",
			"critical",
			fmt.Sprintf("Robot %s is offline", record.RobotID),
		))
	} else if record.ConnectionStatus == "weak_signal" {
		result = append(result, newAlert(
			record.RobotID,
			"WEAK_SIGNAL",
			"warning",
			fmt.Sprintf("Robot %s has weak signal", record.RobotID),
		))
	}

	if record.ObstacleDetected && record.DistanceToObstacle < 1 {
		result = append(result, newAlert(
			record.RobotID,
			"OBSTACLE_CRITICAL",
			"warning",
			fmt.Sprintf("Robot %s has a critical obstacle distance: %.2fm", record.RobotID, record.DistanceToObstacle),
		))
	} else if record.ObstacleDetected && record.DistanceToObstacle < 2 {
		result = append(result, newAlert(
			record.RobotID,
			"OBSTACLE_DETECTED",
			"info",
			fmt.Sprintf("Robot %s detected an obstacle: %.2fm", record.RobotID, record.DistanceToObstacle),
		))
	}

	if record.IsStuck {
		result = append(result, newAlert(
			record.RobotID,
			"ROBOT_STUCK",
			"critical",
			fmt.Sprintf("Robot %s appears to be stuck", record.RobotID),
		))
	}

	return result
}

func newAlert(robotID, alertType, severity, message string) models.Alert {
	return models.Alert{
		ID:        fmt.Sprintf("%s-%d", alertType, time.Now().UnixNano()),
		RobotID:   robotID,
		Type:      alertType,
		Severity:  severity,
		Message:   message,
		Status:    "active",
		CreatedAt: time.Now().UTC(),
	}
}