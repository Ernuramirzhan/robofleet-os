\# RoboFleet OS



RoboFleet OS is a simulation-based backend system for monitoring a fleet of virtual warehouse robots in real time.



The project models warehouse robots that send telemetry data such as position, speed, battery level, temperature, motor load, signal strength, obstacle detection and task status. The backend receives this telemetry, processes it and generates alerts when abnormal robot behavior is detected.



\## Project Goal



The goal of RoboFleet OS is to build a backend platform for robot fleet monitoring, telemetry processing and alert generation.



The system is designed as a portfolio-level engineering project combining:



\* robotics simulation concept

\* IoT-style telemetry ingestion

\* Go backend development

\* PostgreSQL-ready architecture

\* alert evaluation logic

\* automated tests

\* future AI anomaly detection



\## Current Features



Implemented so far:



\* Go backend with Gin

\* Health check endpoint

\* Telemetry ingestion endpoint

\* Robot and telemetry data models

\* Alert data model

\* Rule-based alert evaluation engine

\* Automatic alert generation from telemetry

\* Unit tests for alert engine

\* Project documentation for robot classes, telemetry fields, alert rules and warehouse map



\## Tech Stack



Current stack:



\* Go

\* Gin Web Framework

\* Git / GitHub

\* PowerShell for local development



Planned stack:



\* PostgreSQL

\* Docker

\* WebSocket real-time updates

\* Python robot simulator

\* AI / Ollama anomaly explanation module



\## API Endpoints



\### Health Check



```http

GET /api/v1/health

```



Example response:



```json

{

&#x20; "service": "robofleet-api",

&#x20; "status": "ok",

&#x20; "version": "0.1.0"

}

```



\### Submit Robot Telemetry



```http

POST /api/v1/telemetry

```



Example request:



```json

{

&#x20; "robot\_id": "HLB-001",

&#x20; "robot\_class": "HeavyLoadBot",

&#x20; "timestamp": "2026-06-19T10:45:00Z",

&#x20; "x": 40.0,

&#x20; "y": 15.0,

&#x20; "current\_zone": "Storage Zone",

&#x20; "target\_zone": "Packing Zone",

&#x20; "speed": 0.01,

&#x20; "battery\_level": 8,

&#x20; "temperature": 75.5,

&#x20; "motor\_load": 97.0,

&#x20; "task\_status": "moving",

&#x20; "signal\_strength": 20,

&#x20; "obstacle\_detected": true,

&#x20; "distance\_to\_obstacle": 0.7,

&#x20; "error\_code": null

}

```



Example response includes processed telemetry and generated alerts:



```json

{

&#x20; "message": "telemetry received",

&#x20; "alerts": \[

&#x20;   {

&#x20;     "type": "CRITICAL\_BATTERY",

&#x20;     "severity": "critical",

&#x20;     "status": "active"

&#x20;   },

&#x20;   {

&#x20;     "type": "OVERHEATING",

&#x20;     "severity": "critical",

&#x20;     "status": "active"

&#x20;   },

&#x20;   {

&#x20;     "type": "MOTOR\_OVERLOAD",

&#x20;     "severity": "critical",

&#x20;     "status": "active"

&#x20;   }

&#x20; ]

}

```



\## Alert Types



The current alert engine can detect:



\* `CRITICAL\_BATTERY`

\* `LOW\_BATTERY`

\* `OVERHEATING`

\* `HIGH\_TEMPERATURE`

\* `MOTOR\_OVERLOAD`

\* `HIGH\_MOTOR\_LOAD`

\* `WEAK\_SIGNAL`

\* `ROBOT\_OFFLINE`

\* `OBSTACLE\_DETECTED`

\* `OBSTACLE\_CRITICAL`

\* `ROBOT\_STUCK`



\## Project Structure



```text

RoboFleet-OS/

├── backend/

│   ├── cmd/

│   │   └── api/

│   │       └── main.go

│   ├── internal/

│   │   ├── alerts/

│   │   │   ├── engine.go

│   │   │   └── engine\_test.go

│   │   └── models/

│   │       ├── alert.go

│   │       └── telemetry.go

│   ├── go.mod

│   └── go.sum

├── docs/

│   ├── alert\_rules.md

│   ├── project\_concept.md

│   ├── robot\_classes.md

│   ├── telemetry\_fields.md

│   └── warehouse\_map.md

├── .gitignore

└── README.md

```



\## Run Locally



Go to the backend folder:



```bash

cd backend

```



Run the API server:



```bash

go run ./cmd/api

```



The server will start on:



```text

http://localhost:8080

```



\## Run Tests



```bash

cd backend

go test ./...

```



Expected result:



```text

ok github.com/Ernuramirzhan/robofleet-os/backend/internal/alerts

```



\## Documentation



Project documentation is located in the `docs/` folder:



\* `project\_concept.md` — project idea and goal

\* `robot\_classes.md` — warehouse robot classes and behavior

\* `telemetry\_fields.md` — telemetry data structure

\* `alert\_rules.md` — alert rules and severity levels

\* `warehouse\_map.md` — warehouse zones, routes and obstacles



\## Individual Contributions



\### Yernur Yermekkaliyev



Robotics and automation concept:



\* robot fleet concept

\* warehouse robot classes

\* telemetry field design

\* warehouse map and movement scenarios

\* alert rules from robotics perspective



\### Zhannur Yermekkaliyev



Backend and platform development:



\* Go backend structure

\* API endpoints

\* telemetry ingestion logic

\* data models

\* alert evaluation engine

\* automated tests

\* Git and project setup



\## Roadmap



Next planned steps:



\* add PostgreSQL database

\* save telemetry records

\* save generated alerts

\* add robot registry endpoint

\* build Python robot simulator

\* add WebSocket real-time updates

\* add Docker Compose

\* connect AI / Ollama for anomaly explanation

\* add dashboard or API-based monitoring view



\## Status



The project is currently in active development.



