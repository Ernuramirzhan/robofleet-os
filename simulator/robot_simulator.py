import json
from datetime import datetime, timezone
from urllib.error import HTTPError, URLError
from urllib.request import Request, urlopen


API_URL = "http://localhost:8080/api/v1/telemetry"


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
            print("Status:", response.status)
            print("Response:", response_body)

    except HTTPError as error:
        print("HTTP error:", error.code)
        print(error.read().decode("utf-8"))

    except URLError as error:
        print("Connection error:", error.reason)


if __name__ == "__main__":
    telemetry = {
        "robot_id": "DLB-001",
        "robot_class": "DeliveryBot",
        "timestamp": datetime.now(timezone.utc).isoformat(),
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
    }

    send_telemetry(telemetry)