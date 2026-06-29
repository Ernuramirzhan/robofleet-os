import time

from robot_simulator import build_telemetry_batch, send_telemetry


if __name__ == "__main__":
    print("Starting RoboFleet continuous simulator")
    print("Press Ctrl + C to stop")
    print("=" * 60)

    try:
        while True:
            telemetry_batch = build_telemetry_batch()

            print(f"Sending telemetry for {len(telemetry_batch)} robots")
            print("=" * 60)

            for telemetry in telemetry_batch:
                send_telemetry(telemetry)

            print("Waiting 3 seconds before next telemetry batch...")
            print("=" * 60)
            time.sleep(3)

    except KeyboardInterrupt:
        print("\nSimulator stopped by user")