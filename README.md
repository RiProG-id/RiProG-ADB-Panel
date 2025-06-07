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

if (res.success) {
  console.log("Output:", res.output.trim());
} else {
  console.error("Error:", res.error || "Unknown error");
}
```

## Response Format

```
// General Stucture
{
  "success": true,
  "output": "hasil stdout",
  "error": "jika ada error stderr"
}

// Success example
{
  "success": true,
  "output": "Pixel 4\n"
}

// Failure example
{
  "success": false,
  "output": "",
  "error": "sh: something: not found"
}
```

## More Information

**Author:** [RiProG](https://github.com/RiProG-id)

### Visit:

- [Support ME](https://t.me/RiOpSo/2848)
- [Telegram Channel](https://t.me/RiOpSo)
- [Telegram Group](https://t.me/RiOpSoDisc)
