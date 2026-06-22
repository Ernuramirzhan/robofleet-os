package repository

import (
	"context"

	"github.com/Ernuramirzhan/robofleet-os/backend/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AlertRepository struct {
	db *pgxpool.Pool
}

func NewAlertRepository(db *pgxpool.Pool) *AlertRepository {
	return &AlertRepository{db: db}
}

func (r *AlertRepository) SaveAlert(ctx context.Context, alert models.Alert) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO alerts (
			id,
			robot_id,
			alert_type,
			severity,
			message,
			status,
			created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (id) DO NOTHING
	`,
		alert.ID,
		alert.RobotID,
		alert.Type,
		alert.Severity,
		alert.Message,
		alert.Status,
		alert.CreatedAt,
	)

	return err
}

func (r *AlertRepository) SaveAlerts(ctx context.Context, alerts []models.Alert) error {
	for _, alert := range alerts {
		if err := r.SaveAlert(ctx, alert); err != nil {
			return err
		}
	}

	return nil
}