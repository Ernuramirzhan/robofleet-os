# RoboFleet OS

RoboFleet OS is a simulation-based backend system for monitoring a fleet of virtual warehouse robots in real time.

The system receives telemetry data from robots, stores it in PostgreSQL, analyzes robot status and generates alerts when abnormal behavior is detected.

## Project Goal

The goal of RoboFleet OS is to build a backend platform for robot fleet monitoring, telemetry processing, alert generation and future AI-based anomaly detection.

## Current Features

Implemented so far:

* Python robot simulator
* Multi-robot telemetry simulation with different alert scenarios
* Go backend with Gin
* PostgreSQL database
* Docker Compose setup
* Health check endpoint
* Database health check endpoint
* Telemetry ingestion endpoint
* Alerts API endpoint
* Robots API endpoint
* Telemetry records API endpoint
* Robot, telemetry and alert data models
* Rule-based alert evaluation engine
* Saving robots, telemetry records and alerts to PostgreSQL
* Unit tests for alert engine
* Project documentation in `docs/`

## Tech Stack

Current stack:

* Go
* Gin Web Framework
* PostgreSQL
* Docker Compose
* pgx PostgreSQL driver
* Git / GitHub

Planned stack:

* Python robot simulator
* WebSocket real-time updates
* Dashboard
* AI / Ollama anomaly explanation module

## API Endpoints

### Health Check

```http
GET /api/v1/health
```

### Database Health Check

```http
GET /api/v1/db/health
```

### Submit Robot Telemetry

```http
POST /api/v1/telemetry
```

This endpoint receives telemetry data, saves the robot, saves the telemetry record, generates alerts and stores alerts in PostgreSQL.

### Get Alerts

```http
GET /api/v1/alerts
```

Returns recent alerts saved in PostgreSQL.

### Get Robots

```http
GET /api/v1/robots
```

Returns all robots saved in PostgreSQL.

### Get Telemetry Records

```http
GET /api/v1/telemetry
```

Returns recent telemetry records saved in PostgreSQL.

## Alert Types

The current alert engine can detect:

* `CRITICAL_BATTERY`
* `LOW_BATTERY`
* `OVERHEATING`
* `HIGH_TEMPERATURE`
* `MOTOR_OVERLOAD`
* `HIGH_MOTOR_LOAD`
* `WEAK_SIGNAL`
* `ROBOT_OFFLINE`
* `OBSTACLE_DETECTED`
* `OBSTACLE_CRITICAL`
* `ROBOT_STUCK`

## Project Structure

```text
RoboFleet-OS/
├── backend/
│   ├── cmd/
│   │   └── api/
│   │       └── main.go
│   ├── internal/
│   │   ├── alerts/
│   │   ├── database/
│   │   ├── models/
│   │   └── repository/
│   ├── go.mod
│   └── go.sum
├── docs/
├── docker-compose.yml
├── .env.example
├── .gitignore
└── README.md
```

## Run Locally

Start PostgreSQL:

```bash
docker compose up -d
```

Run the backend:

```bash
cd backend
go run ./cmd/api
```

The server starts on:

```text
http://localhost:8080
```

## Run Tests

```bash
cd backend
go test ./...
```

## Documentation

Project documentation is located in the `docs/` folder:

* `project_concept.md`
* `robot_classes.md`
* `telemetry_fields.md`
* `alert_rules.md`
* `warehouse_map.md`

## Individual Contributions

### Yernur Yermekkaliyev

Robotics and automation concept:

* robot fleet concept
* warehouse robot classes
* telemetry field design
* warehouse map and movement scenarios
* alert rules from robotics perspective

### Zhannur Yermekkaliyev

Backend and platform development:

* Go backend structure
* API endpoints
* PostgreSQL integration
* Docker Compose setup
* telemetry ingestion logic
* data models
* alert evaluation engine
* repository layer
* automated tests
* Git and project setup

## Roadmap

Next planned steps:

* improve simulator with continuous telemetry generation
* add Dockerfile for Go backend
* run backend and PostgreSQL together through Docker Compose
* add WebSocket real-time updates
* add dashboard or API-based monitoring view
* connect AI / Ollama for anomaly explanation

## Status

The project is currently in active development.
