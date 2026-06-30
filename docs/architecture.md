\# RoboFleet OS Architecture



This document describes the current architecture of RoboFleet OS.



\## High-Level Architecture



RoboFleet OS consists of several main parts:



```text

Python Simulator

&#x20;       ↓

Go Backend API

&#x20;       ↓

PostgreSQL Database

&#x20;       ↓

Alert Engine

&#x20;       ↓

Ollama AI Explanation Module

```



\## Main Components



\### 1. Python Simulator



The simulator generates telemetry data for virtual warehouse robots.



Simulator files:



```text

simulator/robot\_simulator.py

simulator/continuous\_simulator.py

simulator/randomized\_simulator.py

```



The simulator sends telemetry data to the backend using HTTP POST requests.



Main endpoint used by the simulator:



```http

POST /api/v1/telemetry

```



The simulator is used instead of real robots during development and demonstration.



\---



\### 2. Go Backend API



The backend is written in Go using the Gin framework.



Main backend entry point:



```text

backend/cmd/api/main.go

```



The backend is responsible for:



\* receiving telemetry data

\* validating request data

\* calculating derived robot status fields

\* saving robots to PostgreSQL

\* saving telemetry records to PostgreSQL

\* generating alerts

\* saving alerts to PostgreSQL

\* providing API endpoints for monitoring

\* calling Ollama for AI explanations



\---



\### 3. PostgreSQL Database



PostgreSQL stores the system data.



Main database schema file:



```text

backend/internal/database/schema.sql

```



Current database tables:



```text

robots

telemetry\_records

alerts

```



The database allows the system to keep historical robot telemetry and alert data.



\---



\### 4. Repository Layer



The repository layer is responsible for database operations.



Repository files:



```text

backend/internal/repository/robot\_repository.go

backend/internal/repository/telemetry\_repository.go

backend/internal/repository/alert\_repository.go

```



The repository layer separates SQL logic from API handler logic.



This makes the backend easier to maintain and extend.



\---



\### 5. Alert Engine



The alert engine analyzes telemetry data and generates alerts.



Main file:



```text

backend/internal/alerts/engine.go

```



Examples of detected alert types:



```text

LOW\_BATTERY

CRITICAL\_BATTERY

HIGH\_TEMPERATURE

OVERHEATING

HIGH\_MOTOR\_LOAD

MOTOR\_OVERLOAD

WEAK\_SIGNAL

OBSTACLE\_DETECTED

OBSTACLE\_CRITICAL

ROBOT\_STUCK

```



The alert engine is rule-based.



For example:



```text

battery\_level < 10  → CRITICAL\_BATTERY

temperature > 70    → OVERHEATING

motor\_load > 95     → MOTOR\_OVERLOAD

signal\_strength < 30 → WEAK\_SIGNAL

```



\---



\### 6. Ollama AI Explanation Module



The AI module uses Ollama to generate explanations for robot problems.



Main file:



```text

backend/internal/ai/client.go

```



AI endpoint:



```http

POST /api/v1/ai/explain

```



Example request:



```json

{

&#x20; "robot\_id": "HLB-002"

}

```



The backend collects:



\* latest telemetry for the robot

\* recent alerts for the robot



Then it builds a prompt and sends it to Ollama.



The AI explanation includes:



\* why the robot is risky

\* likely causes

\* recommended operator actions

\* priority level



\---



\## Data Flow



\### Telemetry Flow



```text

1\. Python simulator creates telemetry JSON.

2\. Simulator sends POST request to /api/v1/telemetry.

3\. Go backend receives the request.

4\. Backend calculates connection\_status and is\_stuck.

5\. Backend saves robot information to PostgreSQL.

6\. Backend saves telemetry record to PostgreSQL.

7\. Alert engine checks telemetry values.

8\. Backend saves generated alerts to PostgreSQL.

9\. Backend returns telemetry result and alerts.

```



\### AI Explanation Flow



```text

1\. User sends POST request to /api/v1/ai/explain.

2\. Backend receives robot\_id.

3\. Backend loads latest telemetry for that robot.

4\. Backend loads recent alerts for that robot.

5\. Backend builds an AI prompt.

6\. Backend sends the prompt to Ollama.

7\. Ollama generates explanation.

8\. Backend returns the explanation through API.

```



\---



\## Docker Architecture



The project uses Docker Compose.



Main Docker Compose file:



```text

docker-compose.yml

```



Current services:



```text

postgres

backend

```



The backend container connects to PostgreSQL using the Docker service name:



```text

postgres:5432

```



The backend connects to Ollama running on the host machine using:



```text

http://host.docker.internal:11434

```



\---



\## Main API Endpoints



```http

GET  /api/v1/health

GET  /api/v1/db/health

POST /api/v1/telemetry

GET  /api/v1/robots

GET  /api/v1/robots/:id

GET  /api/v1/robots/:id/telemetry

GET  /api/v1/robots/:id/alerts

GET  /api/v1/telemetry

GET  /api/v1/alerts

POST /api/v1/ai/explain

```



\---



\## Current Architecture Summary



RoboFleet OS is built as a modular backend system.



The current architecture includes:



\* simulation layer

\* API layer

\* repository layer

\* database layer

\* alert evaluation layer

\* AI explanation layer

\* Docker-based runtime environment



This structure makes the project easy to demonstrate, explain and extend.



