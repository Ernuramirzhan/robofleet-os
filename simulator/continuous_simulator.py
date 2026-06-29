import time

from robot_simulator import build_telemetry_batch, send_telemetry


BATCH_COUNT = 5
DELAY_SECONDS = 3


if __name__ == "__main__":
    print("Starting RoboFleet continuous simulator")
    print(f"Batch count: {BATCH_COUNT}")
    print(f"Delay between batches: {DELAY_SECONDS} seconds")
    print("=" * 60)

    for batch_number in range(1, BATCH_COUNT + 1):
        telemetry_batch = build_telemetry_batch()

        print(f"Batch {batch_number}/{BATCH_COUNT}")
        print(f"Sending telemetry for {len(telemetry_batch)} robots")
        print("=" * 60)

        for telemetry in telemetry_batch:
            send_telemetry(telemetry)

        if batch_number < BATCH_COUNT:
            print(f"Waiting {DELAY_SECONDS} seconds before next telemetry batch...")
            print("=" * 60)
            time.sleep(DELAY_SECONDS)

    print("Continuous simulator finished")