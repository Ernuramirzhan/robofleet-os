package alerts_test

import (
	"testing"

	"github.com/Ernuramirzhan/robofleet-os/backend/internal/alerts"
	"github.com/Ernuramirzhan/robofleet-os/backend/internal/models"
)

func TestEvaluateTelemetryGeneratesCriticalAlerts(t *testing.T) {
	record := models.TelemetryRecord{
		TelemetryInput: models.TelemetryInput{
			RobotID:            "HLB-001",
			BatteryLevel:       8,
			Temperature:        75.5,
			MotorLoad:          97.0,
			ObstacleDetected:   true,
			DistanceToObstacle: 0.7,
		},
		ConnectionStatus: "weak_signal",
		IsStuck:          true,
	}

	generatedAlerts := alerts.EvaluateTelemetry(record)

	expectedAlertTypes := []string{
		"CRITICAL_BATTERY",
		"OVERHEATING",
		"MOTOR_OVERLOAD",
		"WEAK_SIGNAL",
		"OBSTACLE_CRITICAL",
		"ROBOT_STUCK",
	}

	for _, alertType := range expectedAlertTypes {
		if !hasAlertType(generatedAlerts, alertType) {
			t.Fatalf("expected alert type %s, but it was not generated", alertType)
		}
	}

	if len(generatedAlerts) != len(expectedAlertTypes) {
		t.Fatalf("expected %d alerts, got %d", len(expectedAlertTypes), len(generatedAlerts))
	}
}

func TestEvaluateTelemetryReturnsNoAlertsForNormalRobot(t *testing.T) {
	record := models.TelemetryRecord{
		TelemetryInput: models.TelemetryInput{
			RobotID:            "DEL-001",
			BatteryLevel:       76,
			Temperature:        42.5,
			MotorLoad:          55.0,
			ObstacleDetected:   false,
			DistanceToObstacle: 2.8,
		},
		ConnectionStatus: "online",
		IsStuck:          false,
	}

	generatedAlerts := alerts.EvaluateTelemetry(record)

	if len(generatedAlerts) != 0 {
		t.Fatalf("expected 0 alerts, got %d", len(generatedAlerts))
	}
}

func hasAlertType(alertList []models.Alert, alertType string) bool {
	for _, alert := range alertList {
		if alert.Type == alertType {
			return true
		}
	}

	return false
}