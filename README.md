# RoboFleet OS

RoboFleet OS is a simulation-based backend system for monitoring a fleet of virtual warehouse robots.

The system receives telemetry data from robots, stores it in PostgreSQL, analyzes robot status and generates alerts when abnormal behavior is detected.

## Project Goal

The goal of RoboFleet OS is to build a backend platform for robot fleet monitoring, telemetry processing, alert generation and future AI-based anomaly explanation.

## Current Features

Implemented so far:

* Go backend with Gin
* PostgreSQL database
* Docker Compose setup
* Dockerfile for Go backend
* Backend and PostgreSQL run together through Docker Compose
* Python robot simulator
* Multi-robot telemetry simulation with different alert scenarios
* Limited continuous simulator with 5 telemetry batches
* Randomized telemetry simulator
* Health check endpoint
* Database health check endpoint
* Telemetry ingestion endpoint
* Alerts API endpoint
* Robots API endpoint
* Telemetry records API endpoint
* Robot details API endpoint
* Robot-specific telemetry API endpoint
* Robot-specific alerts API endpoint
* Robot, telemetry and alert data models
* Rule-based alert evaluation engine
* Saving robots, telemetry records and alerts to PostgreSQL
* Unit tests for alert engine
* Project documentation in `docs/`
* Demo scenario documentation

## Tech Stack

Current stack:

* Go
* Gin Web Framework
* PostgreSQL
* Docker
* Docker Compose
* pgx PostgreSQL driver
* Python
* Git / GitHub

Planned stack:

* WebSocket real-time updates
* Dashboard
* AI / Ollama anomaly explanation module

## API Endpoints

### Health Check

```http
GET /api/v1/health
```

Example response:

```json
{
  "service": "robofleet-api",
  "status": "ok",
  "version": "0.1.0"
}
```

### Database Health Check

```http
GET /api/v1/db/health
```

Example response:

```json
{
  "database": "connected",
  "status": "ok"
}
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

### Get Robot by ID

```http
GET /api/v1/robots/:id
```

Example:

```http
GET /api/v1/robots/HLB-002
```

Returns details for a specific robot.

### Get Robot Telemetry

```http
GET /api/v1/robots/:id/telemetry
```

Example:

```http
GET /api/v1/robots/HLB-002/telemetry
```

Returns recent telemetry records for a specific robot.

### Get Robot Alerts

```http
GET /api/v1/robots/:id/alerts
```

Example:

```http
GET /api/v1/robots/HLB-002/alerts
```

Returns recent alerts for a specific robot.

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
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
├── docs/
│   ├── alert_rules.md
│   ├── demo_scenario.md
│   ├── project_concept.md
│   ├── robot_classes.md
│   ├── telemetry_fields.md
│   └── warehouse_map.md
├── simulator/
│   ├── robot_simulator.py
│   ├── continuous_simulator.py
│   └── randomized_simulator.py
├── docker-compose.yml
├── .env.example
├── .gitignore
└── README.md
```

## Run Locally

Start PostgreSQL and the backend API with Docker Compose:

```bash
docker compose up -d --build
```

The API server will be available at:

```text
http://localhost:8080
```

Check API health:

```http
GET /api/v1/health
```

Check database connection:

```http
GET /api/v1/db/health
```

To stop all containers:

```bash
docker compose down
```

For local Go development without Docker backend:

```bash
cd backend
go run ./cmd/api
```

Note: if the Docker backend is already running on port 8080, stop it before using `go run`:

```bash
docker compose stop backend
```

## Run Simulators

Send one telemetry batch for multiple robots:

```bash
python simulator/robot_simulator.py
```

Send 5 telemetry batches with a delay between batches:

```bash
python simulator/continuous_simulator.py
```

Send one randomized telemetry batch:

```bash
python simulator/randomized_simulator.py
```

The randomized simulator slightly changes robot position, speed, battery level, temperature, motor load and signal strength on every run.

The backend must be running before starting the simulator.

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
* `demo_scenario.md`

## Demo Scenario

The demo scenario is described in:

```text
docs/demo_scenario.md
```

It includes the basic project demonstration flow:

* start Docker Compose
* check API health
* check database health
* run robot simulator
* check robots
* check telemetry records
* check alerts
* stop containers

## Individual Contributions

### Yernur Yermekkaliyev

Robotics, automation concept and simulation:

* robot fleet concept
* warehouse robot classes
* telemetry field design
* warehouse map and movement scenarios
* alert rules from robotics perspective
* Python robot simulator
* continuous telemetry simulation
* randomized telemetry simulation

### Zhannur Yermekkaliyev

Backend and platform development:

* Go backend structure
* API endpoints
* PostgreSQL integration
* Docker Compose setup
* Dockerfile for backend
* telemetry ingestion logic
* data models
* alert evaluation engine
* repository layer
* automated tests
* Git and project setup
* README and project documentation

## Roadmap

Next planned steps:

* improve README and demo documentation
* add WebSocket real-time updates
* add dashboard or API-based monitoring view
* connect AI / Ollama for anomaly explanation
* prepare final demo video

## Status

The project is currently in active development.
