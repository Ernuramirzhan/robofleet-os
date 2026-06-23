package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Ernuramirzhan/robofleet-os/backend/internal/alerts"
	"github.com/Ernuramirzhan/robofleet-os/backend/internal/database"
	"github.com/Ernuramirzhan/robofleet-os/backend/internal/models"
	"github.com/Ernuramirzhan/robofleet-os/backend/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	dbPool, err := database.Connect(ctx)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer dbPool.Close()

	robotRepo := repository.NewRobotRepository(dbPool)
	telemetryRepo := repository.NewTelemetryRepository(dbPool)
	alertRepo := repository.NewAlertRepository(dbPool)

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api/v1")
	{
		api.GET("/health", healthHandler)
		api.GET("/db/health", dbHealthHandler(dbPool))
		api.GET("/alerts", alertsHandler(alertRepo))
		api.GET("/robots", robotsHandler(robotRepo))
		api.GET("/telemetry", telemetryListHandler(telemetryRepo))
		api.POST("/telemetry", telemetryHandler(robotRepo, telemetryRepo, alertRepo))
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

func dbHealthHandler(dbPool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := dbPool.Ping(c.Request.Context()); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "error",
				"error":  "database is not reachable",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   "ok",
			"database": "connected",
		})
	}
}

func telemetryHandler(
	robotRepo *repository.RobotRepository,
	telemetryRepo *repository.TelemetryRepository,
	alertRepo *repository.AlertRepository,
) gin.HandlerFunc {
	return func(c *gin.Context) {
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

		robot := models.Robot{
			ID:          input.RobotID,
			Name:        input.RobotID,
			RobotClass:  input.RobotClass,
			Status:      input.TaskStatus,
			CurrentZone: input.CurrentZone,
			TargetZone:  input.TargetZone,
		}

		if err := robotRepo.UpsertRobot(c.Request.Context(), robot); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "failed to save robot",
				"details": err.Error(),
			})
			return
		}

		telemetryID, err := telemetryRepo.SaveTelemetry(c.Request.Context(), record)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "failed to save telemetry",
				"details": err.Error(),
			})
			return
		}

		if err := alertRepo.SaveAlerts(c.Request.Context(), generatedAlerts); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "failed to save alerts",
				"details": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message":      "telemetry received",
			"telemetry_id": telemetryID,
			"data":         record,
			"alerts":       generatedAlerts,
		})
	}
}

func alertsHandler(alertRepo *repository.AlertRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		recentAlerts, err := alertRepo.GetRecentAlerts(c.Request.Context(), 20)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "failed to get alerts",
				"details": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"alerts": recentAlerts,
			"count":  len(recentAlerts),
		})
	}
}

func robotsHandler(robotRepo *repository.RobotRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		robots, err := robotRepo.GetAllRobots(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "failed to get robots",
				"details": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"robots": robots,
			"count":  len(robots),
		})
	}
}

func telemetryListHandler(telemetryRepo *repository.TelemetryRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		records, err := telemetryRepo.GetRecentTelemetry(c.Request.Context(), 20)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "failed to get telemetry records",
				"details": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"telemetry": records,
			"count":     len(records),
		})
	}
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