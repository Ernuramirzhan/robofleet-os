CREATE TABLE robots (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    robot_class VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    current_zone VARCHAR(100),
    target_zone VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE telemetry_records (
    id BIGSERIAL PRIMARY KEY,
    robot_id VARCHAR(50) NOT NULL REFERENCES robots(id),
    robot_class VARCHAR(50) NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    x DOUBLE PRECISION NOT NULL,
    y DOUBLE PRECISION NOT NULL,
    current_zone VARCHAR(100) NOT NULL,
    target_zone VARCHAR(100) NOT NULL,
    speed DOUBLE PRECISION NOT NULL,
    battery_level INTEGER NOT NULL,
    temperature DOUBLE PRECISION NOT NULL,
    motor_load DOUBLE PRECISION NOT NULL,
    task_status VARCHAR(50) NOT NULL,
    signal_strength INTEGER NOT NULL,
    obstacle_detected BOOLEAN NOT NULL,
    distance_to_obstacle DOUBLE PRECISION,
    error_code VARCHAR(100),
    connection_status VARCHAR(50) NOT NULL,
    route_deviation DOUBLE PRECISION NOT NULL DEFAULT 0,
    is_stuck BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE alerts (
    id VARCHAR(100) PRIMARY KEY,
    robot_id VARCHAR(50) NOT NULL REFERENCES robots(id),
    alert_type VARCHAR(100) NOT NULL,
    severity VARCHAR(50) NOT NULL,
    message TEXT NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_telemetry_robot_id ON telemetry_records(robot_id);
CREATE INDEX idx_telemetry_timestamp ON telemetry_records(timestamp);
CREATE INDEX idx_alerts_robot_id ON alerts(robot_id);
CREATE INDEX idx_alerts_status ON alerts(status);
CREATE INDEX idx_alerts_severity ON alerts(severity);