import copy
import random

from robot_simulator import build_telemetry_batch, send_telemetry


def clamp(value, min_value, max_value):
    return max(min_value, min(value, max_value))


def randomize_telemetry(telemetry):
    item = copy.deepcopy(telemetry)

    item["x"] = round(item["x"] + random.uniform(-1.5, 1.5), 2)
    item["y"] = round(item["y"] + random.uniform(-1.5, 1.5), 2)

    item["speed"] = round(clamp(item["speed"] + random.uniform(-0.2, 0.2), 0, 2.5), 2)
    item["battery_level"] = int(clamp(item["battery_level"] - random.randint(0, 3), 0, 100))
    item["temperature"] = round(clamp(item["temperature"] + random.uniform(-2.0, 2.0), 20, 90), 1)
    item["motor_load"] = round(clamp(item["motor_load"] + random.uniform(-5.0, 5.0), 0, 100), 1)
    item["signal_strength"] = int(clamp(item["signal_strength"] + random.randint(-8, 8), 0, 100))

    if item["robot_id"] == "HLB-002":
        item["speed"] = round(random.uniform(0.0, 0.04), 2)
        item["battery_level"] = random.randint(5, 9)
        item["temperature"] = round(random.uniform(72.0, 78.0), 1)
        item["motor_load"] = round(random.uniform(96.0, 100.0), 1)
        item["signal_strength"] = random.randint(10, 24)
        item["obstacle_detected"] = True
        item["distance_to_obstacle"] = round(random.uniform(0.3, 0.9), 2)

    if item["robot_id"] == "PKB-001":
        item["battery_level"] = random.randint(14, 19)

    if item["robot_id"] == "PTB-001":
        item["signal_strength"] = random.randint(15, 29)

    if item["robot_id"] == "MNB-001":
        item["temperature"] = round(random.uniform(60.5, 65.0), 1)
        item["motor_load"] = round(random.uniform(85.0, 92.0), 1)
        item["obstacle_detected"] = True
        item["distance_to_obstacle"] = round(random.uniform(1.1, 1.9), 2)

    return item


if __name__ == "__main__":
    base_batch = build_telemetry_batch()
    randomized_batch = [randomize_telemetry(telemetry) for telemetry in base_batch]

    print(f"Sending randomized telemetry for {len(randomized_batch)} robots")
    print("=" * 60)

    for telemetry in randomized_batch:
        send_telemetry(telemetry)