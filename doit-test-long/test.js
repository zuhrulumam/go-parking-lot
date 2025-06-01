import http from "k6/http";
import { check, sleep } from "k6";

// Store vehicle numbers per VU
let parkedVehicles = [];

const vehicleTypes = ["A", "M", "B"];

export let options = {
  stages: [
    { duration: "2m", target: 50 }, // ramp up to 50 users
    { duration: "5m", target: 50 }, // stay at 50 users
    { duration: "2m", target: 0 }, // ramp down
  ],
};

function getRandomVehicleNumber() {
  const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
  const nums = Math.floor(Math.random() * 9000 + 1000); // random 4-digit
  const letters =
    chars.charAt(Math.floor(Math.random() * chars.length)) +
    chars.charAt(Math.floor(Math.random() * chars.length));
  return `${letters}-${nums}`;
}

function parkVehicle() {
  const vehicleType =
    vehicleTypes[Math.floor(Math.random() * vehicleTypes.length)];
  const vehicleNumber = getRandomVehicleNumber();

  const payload = JSON.stringify({
    vehicle_type: vehicleType,
    vehicle_number: vehicleNumber,
  });

  const headers = { "Content-Type": "application/json" };
  // const res = http.post("http://parking.localhost/vehicle/park", payload, {
  //   headers,
  // });
  const res = http.post("http://localhost:8080/vehicle/park", payload, {
    headers,
  });

  const passed = check(res, {
    "park status is 200": (r) => r.status === 200,
  });

  if (!passed) {
    let message = `❌ Park failed for ${vehicleNumber} with status ${res.status}`;

    try {
      const error = res.json().debug_error;
      if (error) message += ` | debug_error: ${error}`;
    } catch (e) {
      message += " | failed to parse debug_error";
    }
    console.error(message);
  } else {
    parkedVehicles.push(vehicleNumber);
  }
}

function unparkVehicle() {
  if (parkedVehicles.length === 0) return;

  const vehicleNumber = parkedVehicles.shift(); // remove first parked
  const payload = JSON.stringify({ vehicle_number: vehicleNumber });
  const headers = { "Content-Type": "application/json" };
  // const res = http.post("http://parking.localhost/vehicle/unpark", payload, {
  //   headers,
  // });
  const res = http.post("http://localhost:8080/vehicle/unpark", payload, {
    headers,
  });

  const passed = check(res, {
    "unpark status is 200": (r) => r.status === 200,
  });

  if (!passed) {
    console.error(
      `❌ Unpark failed for ${vehicleNumber} with status ${res.status}`
    );
  }
}

export default function () {
  // 40% chance to park, 60% to unpark
  if (Math.random() < 0.4) {
    parkVehicle();
  } else {
    unparkVehicle();
  }

  sleep(1); // wait between iterations
}
