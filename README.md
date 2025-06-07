# RiProG-ADB-Panel
A minimal ADB command panel served via local web UI on Android.

# RiProG API

RiProG provides a JavaScript module 'RiProG.js' that allows you to run Android shell commands directly from the local browser.

## Usage

```
// Import API
// Import API
import { exec } from "./RiProG.js";

// Execute a command
const res = await exec("getprop ro.product.model");

if (res.errno === 0) {
  console.log("Output:", res.stdout.trim() || "(no output)");
} else {
  console.error(`Error (exit code ${res.errno}):`, res.stderr.trim() || "Unknown error");
}
```

## Response Format


```
{
  "errno": <integer>,       // the shell process exit code, 0 means success
  "stdout": "<string>",     // the standard output of the command
  "stderr": "<string>"      // the standard error output of the command, empty if none
}

// Success example
{
  "errno": 0,
  "stdout": "Pixel 4\n",
  "stderr": ""
}

// Failure example
{
  "errno": 172,
  "stdout": "",
  "stderr": "sh: something: not found\n"
}
```

## More Information

**Author:** [RiProG](https://github.com/RiProG-id)

### Visit:

- [Support ME](https://t.me/RiOpSo/2848)
- [Telegram Channel](https://t.me/RiOpSo)
- [Telegram Group](https://t.me/RiOpSoDisc)
