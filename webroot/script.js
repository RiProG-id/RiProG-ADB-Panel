import { exec } from "./RiProG.js";

const output = document.getElementById("output");

async function showDeviceModel() {
  try {
    const res = await exec("getprop ro.product.model");
    if (res.success) {
      output.textContent = res.output.trim() || "(no output)";
    } else {
      output.textContent = `Error: ${res.error || "Unknown error"}`;
    }
  } catch (e) {
    output.textContent = "Fetch error: " + e.message;
  }
}

showDeviceModel();
