# RoboFleet OS Demo Scenario

This document describes a simple demo flow for showing how RoboFleet OS works.

## Demo Goal

The goal of the demo is to show the full data flow:

```text
Python simulator
→ Go backend API
→ PostgreSQL database
→ alert generation
→ API monitoring endpoints
```

## 1. Start the Project

Run PostgreSQL and the backend API with Docker Compose:

```bash
docker compose up -d --build
```

Check running containers:

```bash
docker ps
```

Expected containers:

```text
robofleet-postgres
robofleet-api
```

## 2. Check API Health

Open in browser:

```text
http://localhost:8080/api/v1/health
```

Expected response:

```json
{
  "service": "robofleet-api",
  "status": "ok",
  "version": "0.1.0"
}
```

## 3. Check Database Connection

Open in browser:

```text
http://localhost:8080/api/v1/db/health
```

Expected response:

```json
{
  "database": "connected",
  "status": "ok"
}
```

## 4. Run Robot Simulator

Run one telemetry batch:

```bash
python simulator/robot_simulator.py
```

Or run limited continuous simulation:

```bash
python simulator/continuous_simulator.py
```

Or run one randomized telemetry batch:

```bash
python simulator/randomized_simulator.py
```

The randomized simulator slightly changes robot position, speed, battery level, temperature, motor load and signal strength on every run.

The simulator sends telemetry for multiple warehouse robots:

```text
DLB-001 — DeliveryBot
PKB-001 — PickerBot
PTB-001 — PatrolBot
HLB-002 — HeavyLoadBot
MNB-001 — MaintenanceBot
```

## 5. Check Robots

Open in browser:

```text
http://localhost:8080/api/v1/robots
```

This endpoint shows all robots saved in PostgreSQL.

## 6. Check Telemetry Records

Open in browser:

```text
http://localhost:8080/api/v1/telemetry
```

This endpoint shows recent telemetry records received from the simulator.

## 7. Check Alerts

Open in browser:

```text
http://localhost:8080/api/v1/alerts
```

This endpoint shows generated alerts.

Example alert types:

```text
LOW_BATTERY
WEAK_SIGNAL
CRITICAL_BATTERY
OVERHEATING
MOTOR_OVERLOAD
OBSTACLE_CRITICAL
ROBOT_STUCK
```

## 8. Check Specific Robot Data

Open in browser:

```text
http://localhost:8080/api/v1/robots/HLB-002
```

This endpoint shows details for one specific robot.

Open robot telemetry:

```text
http://localhost:8080/api/v1/robots/HLB-002/telemetry
```

This endpoint shows recent telemetry records only for robot `HLB-002`.

Open robot alerts:

```text
http://localhost:8080/api/v1/robots/HLB-002/alerts
```

This endpoint shows recent alerts only for robot `HLB-002`.

This is useful for checking a problematic robot separately from the whole fleet.

## 9. Stop the Project

Stop all containers:

```bash
docker compose down
```

## Demo Summary

RoboFleet OS demonstrates a backend system for robot fleet monitoring.

The project shows:

* telemetry ingestion
* robot registry
* PostgreSQL persistence
* rule-based alert generation
* Docker Compose deployment
* Python-based robot simulation
* API-based monitoring endpoints
