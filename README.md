# RiProG-ADB-Panel
A minimal ADB command panel served via local web UI on Android.

# RiProG API

RiProG provides a JavaScript module 'RiProG.js' that allows you to run Android shell commands directly from the local browser.

## Usage

```
// Import API
import { exec } from "./RiProG.js";

// Execute a command
const res = await exec("getprop ro.product.model");

if (res.exitCode === 0) {
  console.log("Output:", res.stdOut.trim() || "(no output)");
} else {
  console.error(`Error (exit code ${res.exitCode}):`, res.stdErr.trim() || "Unknown error");
}
```

## Response Format


```
{
  "exitCode": <integer>,   // the shell process exit code, 0 means success
  "stdOut": "<string>",    // the standard output of the command
  "stdErr": "<string>"     // the standard error output of the command, empty if none
}

// Success example
{
  "exitCode": 0,
  "stdOut": "Pixel 4\n",
  "stdErr": ""
}

// Failure example
{
  "exitCode": 172,
  "stdOut": "",
  "stdErr": "sh: something: not found\n"
}
```

## More Information

**Author:** [RiProG](https://github.com/RiProG-id)

### Visit:

- [Support ME](https://t.me/RiOpSo/2848)
- [Telegram Channel](https://t.me/RiOpSo)
- [Telegram Group](https://t.me/RiOpSoDisc)
