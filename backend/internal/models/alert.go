package models

import "time"

type Alert struct {
	ID        string    `json:"id"`
	RobotID   string    `json:"robot_id"`
	Type      string    `json:"type"`
	Severity  string    `json:"severity"`
	Message   string    `json:"message"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}