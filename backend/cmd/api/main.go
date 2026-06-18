package main

import (
	"log"
	"net/http"

	"github.com/Ernuramirzhan/robofleet-os/backend/internal/alerts"
	"github.com/Ernuramirzhan/robofleet-os/backend/internal/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api/v1")
	{
		api.GET("/health", healthHandler)
		api.POST("/telemetry", telemetryHandler)
	}

	log.Println("RoboFleet API is running on http://localhost:8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "robofleet-api",
		"version": "0.1.0",
	})
}

func telemetryHandler(c *gin.Context) {
	var input models.TelemetryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid telemetry payload",
			"details": err.Error(),
		})
		return
	}

	record := models.TelemetryRecord{
		TelemetryInput:   input,
		ConnectionStatus: calculateConnectionStatus(input.SignalStrength),
		RouteDeviation:   0.0,
		IsStuck:          input.Speed < 0.05 && input.TaskStatus == "moving",
	}

	generatedAlerts := alerts.EvaluateTelemetry(record)

	c.JSON(http.StatusCreated, gin.H{
		"message": "telemetry received",
		"data":    record,
		"alerts":  generatedAlerts,
	})
}

func calculateConnectionStatus(signalStrength int) string {
	if signalStrength <= 0 {
		return "offline"
	}

	if signalStrength < 30 {
		return "weak_signal"
	}

	return "online"
}