\# RoboFleet OS Demo Scenario



This document describes a simple demo flow for showing how RoboFleet OS works.



\## Demo Goal



The goal of the demo is to show the full data flow:



```text

Python simulator

→ Go backend API

→ PostgreSQL database

→ alert generation

→ API monitoring endpoints

```



\## 1. Start the Project



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



\## 2. Check API Health



Open in browser:



```text

http://localhost:8080/api/v1/health

```



Expected response:



```json

{

&#x20; "service": "robofleet-api",

&#x20; "status": "ok",

&#x20; "version": "0.1.0"

}

```



\## 3. Check Database Connection



Open in browser:



```text

http://localhost:8080/api/v1/db/health

```



Expected response:



```json

{

&#x20; "database": "connected",

&#x20; "status": "ok"

}

```



\## 4. Run Robot Simulator



Run one telemetry batch:



```bash

python simulator/robot\_simulator.py

```



Or run limited continuous simulation:



```bash

python simulator/continuous\_simulator.py

```



The simulator sends telemetry for multiple warehouse robots:



```text

DLB-001 — DeliveryBot

PKB-001 — PickerBot

PTB-001 — PatrolBot

HLB-002 — HeavyLoadBot

MNB-001 — MaintenanceBot

```



\## 5. Check Robots



Open in browser:



```text

http://localhost:8080/api/v1/robots

```



This endpoint shows all robots saved in PostgreSQL.



\## 6. Check Telemetry Records



Open in browser:



```text

http://localhost:8080/api/v1/telemetry

```



This endpoint shows recent telemetry records received from the simulator.



\## 7. Check Alerts



Open in browser:



```text

http://localhost:8080/api/v1/alerts

```



This endpoint shows generated alerts.



Example alert types:



```text

LOW\_BATTERY

WEAK\_SIGNAL

CRITICAL\_BATTERY

OVERHEATING

MOTOR\_OVERLOAD

OBSTACLE\_CRITICAL

ROBOT\_STUCK

```



\## 8. Stop the Project



Stop all containers:



```bash

docker compose down

```



\## Demo Summary



RoboFleet OS demonstrates a backend system for robot fleet monitoring.



The project shows:



\* telemetry ingestion

\* robot registry

\* PostgreSQL persistence

\* rule-based alert generation

\* Docker Compose deployment

\* Python-based robot simulation

\* API-based monitoring endpoints



