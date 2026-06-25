import json
from datetime import datetime, timezone
from urllib.error import HTTPError, URLError
from urllib.request import Request, urlopen


API_URL = "http://localhost:8080/api/v1/telemetry"


def current_timestamp():
    return datetime.now(timezone.utc).isoformat()


def send_telemetry(payload):
    data = json.dumps(payload).encode("utf-8")

    request = Request(
        API_URL,
        data=data,
        headers={"Content-Type": "application/json"},
        method="POST",
    )

    try:
        with urlopen(request, timeout=10) as response:
            response_body = response.read().decode("utf-8")
            print(f"{payload['robot_id']} -> Status: {response.status}")
            print(response_body)
            print("-" * 60)

    except HTTPError as error:
        print(f"{payload['robot_id']} -> HTTP error: {error.code}")
        print(error.read().decode("utf-8"))
        print("-" * 60)

    except URLError as error:
        print(f"{payload['robot_id']} -> Connection error: {error.reason}")
        print("-" * 60)


def build_telemetry_batch():
    return [
        {
            "robot_id": "DLB-001",
            "robot_class": "DeliveryBot",
            "timestamp": current_timestamp(),
            "x": 12.5,
            "y": 8.3,
            "current_zone": "Loading Zone",
            "target_zone": "Packing Zone",
            "speed": 1.2,
            "battery_level": 76,
            "temperature": 38.5,
            "motor_load": 45.0,
            "task_status": "moving",
            "signal_strength": 88,
            "obstacle_detected": False,
            "distance_to_obstacle": 5.0,
            "error_code": None,
        },
        {
            "robot_id": "PKB-001",
            "robot_class": "PickerBot",
            "timestamp": current_timestamp(),
            "x": 22.0,
            "y": 17.4,
            "current_zone": "Shelf Zone",
            "target_zone": "Packing Zone",
            "speed": 0.8,
            "battery_level": 18,
            "temperature": 42.0,
            "motor_load": 52.0,
            "task_status": "picking",
            "signal_strength": 75,
            "obstacle_detected": False,
            "distance_to_obstacle": 4.2,
            "error_code": None,
        },
        {
            "robot_id": "PTB-001",
            "robot_class": "PatrolBot",
            "timestamp": current_timestamp(),
            "x": 35.0,
            "y": 6.5,
            "current_zone": "Security Zone",
            "target_zone": "Storage Zone",
            "speed": 1.5,
            "battery_level": 64,
            "temperature": 36.8,
            "motor_load": 38.0,
            "task_status": "patrolling",
            "signal_strength": 25,
            "obstacle_detected": False,
            "distance_to_obstacle": 6.0,
            "error_code": None,
        },
        {
            "robot_id": "HLB-002",
            "robot_class": "HeavyLoadBot",
            "timestamp": current_timestamp(),
            "x": 44.5,
            "y": 19.0,
            "current_zone": "Storage Zone",
            "target_zone": "Loading Zone",
            "speed": 0.02,
            "battery_level": 9,
            "temperature": 74.0,
            "motor_load": 98.0,
            "task_status": "moving",
            "signal_strength": 18,
            "obstacle_detected": True,
            "distance_to_obstacle": 0.6,
            "error_code": "MOTOR_STRESS",
        },
        {
            "robot_id": "MNB-001",
            "robot_class": "MaintenanceBot",
            "timestamp": current_timestamp(),
            "x": 5.0,
            "y": 25.5,
            "current_zone": "Maintenance Zone",
            "target_zone": "Shelf Zone",
            "speed": 0.6,
            "battery_level": 55,
            "temperature": 61.5,
            "motor_load": 86.0,
            "task_status": "repairing",
            "signal_strength": 92,
            "obstacle_detected": True,
            "distance_to_obstacle": 1.5,
            "error_code": None,
        },
    ]


if __name__ == "__main__":
    telemetry_batch = build_telemetry_batch()

    print(f"Sending telemetry for {len(telemetry_batch)} robots")
    print("=" * 60)

    for telemetry in telemetry_batch:
        send_telemetry(telemetry)