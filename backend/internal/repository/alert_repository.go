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

func (r *AlertRepository) GetRecentAlerts(ctx context.Context, limit int) ([]models.Alert, error) {
	if limit <= 0 {
		limit = 20
	}

	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			robot_id,
			alert_type,
			severity,
			message,
			status,
			created_at
		FROM alerts
		ORDER BY created_at DESC
		LIMIT $1
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Alert

	for rows.Next() {
		var alert models.Alert

		if err := rows.Scan(
			&alert.ID,
			&alert.RobotID,
			&alert.Type,
			&alert.Severity,
			&alert.Message,
			&alert.Status,
			&alert.CreatedAt,
		); err != nil {
			return nil, err
		}

		result = append(result, alert)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *AlertRepository) GetAlertsByRobotID(ctx context.Context, robotID string, limit int) ([]models.Alert, error) {
	if limit <= 0 {
		limit = 20
	}

	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			robot_id,
			alert_type,
			severity,
			message,
			status,
			created_at
		FROM alerts
		WHERE robot_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`, robotID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Alert

	for rows.Next() {
		var alert models.Alert

		if err := rows.Scan(
			&alert.ID,
			&alert.RobotID,
			&alert.Type,
			&alert.Severity,
			&alert.Message,
			&alert.Status,
			&alert.CreatedAt,
		); err != nil {
			return nil, err
		}

		result = append(result, alert)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
