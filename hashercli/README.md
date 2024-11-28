# HasherCLI

This tool is part of the `luabindingtutorial` repository and demonstrates how to create a basic CLI binder.

## Usage

Hasher allows you to hash strings using different algorithms. Below is an example of how to use the tool:

```sh
local hasherbinder = require("hasherbinder")
print(hasherbinder.hash("sha256_raw", "https://example.com/some/link?q1=1&q2=2&g=3"))
```