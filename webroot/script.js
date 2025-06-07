import { exec } from "./RiProG.js";

const output = document.getElementById("output");

async function showDeviceModel() {
  try {
    const res = await exec("getprop ro.product.model");
    if (res.exitCode === 0) {
      output.textContent = res.stdOut.trim() || "(no output)";
    } else {
      output.textContent = `Error: ${res.stdErr.trim() || "Unknown error"}`;
    }
  } catch (e) {
    output.textContent = "Fetch error: " + e.message;
  }
}

showDeviceModel();
