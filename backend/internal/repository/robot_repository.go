package repository

import (
	"context"
	"time"

	"github.com/Ernuramirzhan/robofleet-os/backend/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RobotRepository struct {
	db *pgxpool.Pool
}

func NewRobotRepository(db *pgxpool.Pool) *RobotRepository {
	return &RobotRepository{db: db}
}

func (r *RobotRepository) UpsertRobot(ctx context.Context, robot models.Robot) error {
	if robot.CreatedAt.IsZero() {
		robot.CreatedAt = time.Now().UTC()
	}

	_, err := r.db.Exec(ctx, `
		INSERT INTO robots (
			id,
			name,
			robot_class,
			status,
			current_zone,
			target_zone,
			created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (id) DO UPDATE SET
			robot_class = EXCLUDED.robot_class,
			status = EXCLUDED.status,
			current_zone = EXCLUDED.current_zone,
			target_zone = EXCLUDED.target_zone
	`,
		robot.ID,
		robot.Name,
		robot.RobotClass,
		robot.Status,
		robot.CurrentZone,
		robot.TargetZone,
		robot.CreatedAt,
	)

	return err
}